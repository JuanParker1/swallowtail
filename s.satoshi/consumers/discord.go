package consumers

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/monzo/slog"
	"github.com/monzo/terrors"

	"swallowtail/libraries/util"
	discordclient "swallowtail/s.discord/client"
	discordproto "swallowtail/s.discord/proto"
	"swallowtail/s.satoshi/formatter"
	"swallowtail/s.satoshi/parser"
	tradeengineproto "swallowtail/s.trade-engine/proto"
)

const (
	discordConsumerID = "discord-consumer"
	discordChannelUrl = "https://discord.com/api/v9/channels/%s/messages?limit=1"
)

var (
	discordConsumerToken string
	discordConsumerName  = "satoshi_consumer"
)

func init() {
	discordConsumerToken = util.SetEnv("SATOSHI_DISCORD_CONSUMER_1_API_TOKEN")
	register(discordConsumerID, DiscordConsumer{
		Active: true,
	})
}

type DiscordConsumer struct {
	Active bool
}

func (dc DiscordConsumer) Receiver(ctx context.Context, c chan *ConsumerMessage, d chan struct{}, _ bool) {
	discordClient := discordclient.New(discordConsumerName, discordConsumerToken, false)

	// Add handlers.
	discordClient.AddHandler(handleModMessages(ctx, c, dc.Active))
	discordClient.AddHandler(handleSwingMessages(ctx, c, dc.Active))
	discordClient.AddHandler(handleInternalCallsMessages(ctx, c, dc.Active))

	defer slog.Warn(ctx, "Discord consumer stop signal received.")
	defer discordClient.Close()

	select {
	case <-d:
	case <-ctx.Done():
	}
}

func (dc DiscordConsumer) IsActive() bool {
	return dc.Active
}

func handleModMessages(
	ctx context.Context, c chan *ConsumerMessage, isActive bool,
) func(s *discordgo.Session, m *discordgo.MessageCreate) {
	return func(s *discordgo.Session, mc *discordgo.MessageCreate) {
		m := mc.Message

		if m.ChannelID != discordproto.DiscordMoonModMessagesChannel {
			return
		}

		// Parse our content.
		parsedContent, err := getLatestChannelMessages(ctx, s, m.ChannelID)
		if err != nil {
			slog.Error(ctx, "Failed to get latest mod message: %v", err)
			return
		}

		msgs := []*ConsumerMessage{}
		for i, pc := range parsedContent {
			// First lets check if the content is part of any 1-10k challenge.
			msgs = append(msgs, &ConsumerMessage{
				ConsumerID:       discordConsumerID,
				DiscordChannelID: discordproto.DiscordSatoshiModMessagesChannel,
				Message:          formatContent(ctx, pc.Author.Username, pc.Timestamp, pc.Content),
				Created:          time.Now(),
				IsActive:         isActive,
				Attachments:      m.Attachments,
				Metadata: map[string]string{
					"message": fmt.Sprintf("%v", i),
					"total":   fmt.Sprintf("%v", len(parsedContent)),
				},
			})

			switch {
			case containsAstekz1To10kChallenge(m.Author.Username, strings.ToLower(pc.Content)):
				slog.Debug(ctx, "Astekz 1-10k challenge message received: %s", pc.Content)
				msg := &ConsumerMessage{
					ConsumerID:       discordConsumerID,
					DiscordChannelID: discordproto.DiscordSatoshiChallengesChannel,
					Message:          warning("1-10k challenge update from Astekz", pc.Content),
					Created:          time.Now(),
					IsActive:         isActive,
				}
				msgs = append(msgs, msg)
			case containsRego1To10kChallenge(m.Author.Username, strings.ToLower(pc.Content)):
				slog.Debug(ctx, "Rego 1-10k challenge message received: %s", pc.Content)
				msg := &ConsumerMessage{
					ConsumerID:       discordConsumerID,
					DiscordChannelID: discordproto.DiscordSatoshiChallengesChannel,
					Message:          warning("1-10k challenge update from Rego", pc.Content),
					Created:          time.Now(),
					IsActive:         isActive,
				}
				msgs = append(msgs, msg)
			}

			// Attempt to parse a trade strategy.
			tradeStrategy, err := parser.Parse(ctx, discordproto.DiscordMoonModMessagesChannel, pc.Content, mc, tradeengineproto.ACTOR_TYPE_EXTERNAL)
			if err != nil {
				// No trade strategy can be parsed; so lets continue.
				slog.Trace(ctx, "Failed to parse trade strategy: %+v, content: %s", err, pc.Content)
				continue
			}

			if tradeStrategy == nil {
				continue
			}

			now := time.Now().UTC()
			idempotencyKey := fmt.Sprintf("tradestrategy-%s-%s-%v-%v-%v", tradeStrategy.ActorId, tradeStrategy.Asset, entriesAsString(tradeStrategy.Entries), tradeStrategy.StopLoss, now.Truncate(time.Hour))

			// Sign our trade with our idempotency key.
			tradeStrategy.IdempotencyKey = idempotencyKey

			rsp, err := createTradeStrategy(ctx, tradeStrategy)
			if err != nil {
				slog.Error(ctx, "Failed to create trade: %v, Error: %v", tradeStrategy, err)
				continue
			}

			tradeStrategy.TradeStrategyId = rsp.TradeStrategyId
			tradeStrategy.Created = rsp.Created
			tradeContent := formatter.FormatTradeStrategy("WWG", tradeStrategy, pc.Content)

			msg := &ConsumerMessage{
				ConsumerID:       discordConsumerID,
				DiscordChannelID: discordproto.DiscordSatoshiModTradesChannel,
				Message:          tradeContent,
				Created:          time.Now(),
				IsActive:         isActive,
				Metadata: map[string]string{
					"message":           fmt.Sprintf("%v", i),
					"total":             fmt.Sprintf("%v", len(parsedContent)),
					"trade_strategy_id": tradeStrategy.TradeStrategyId,
				},
				Poller: func(ctx context.Context, messageID string) error {
					// Inject the trade ID.
					return startTradeParticipantsPoller(ctx, messageID, tradeStrategy.TradeStrategyId)
				},
			}

			msgs = append(msgs, msg)
		}

		// Lets publish our messages.
		for _, msg := range msgs {
			select {
			case c <- msg:
			default:
				slog.Warn(ctx, "Failed to publish satoshi mods msg; blocked channel")
			}
		}
	}
}

func handleSwingMessages(
	ctx context.Context, c chan *ConsumerMessage, isActive bool,
) func(s *discordgo.Session, m *discordgo.MessageCreate) {
	return func(s *discordgo.Session, mc *discordgo.MessageCreate) {
		m := mc.Message

		if m.ChannelID != discordproto.DiscordMoonSwingGroupChannel {
			return
		}

		parsedContent, err := getLatestChannelMessages(ctx, s, m.ChannelID)
		if err != nil {
			slog.Error(ctx, "Failed to get latest mod message: %v", err)
			return
		}

		msgs := []*ConsumerMessage{}
		for i, pc := range parsedContent {
			// Add message regardless.
			msgs = append(msgs, &ConsumerMessage{
				ConsumerID:       discordConsumerID,
				DiscordChannelID: discordproto.DiscordSatoshiSwingsChannel,
				Message:          formatContent(ctx, pc.Author.Username, pc.Timestamp, pc.Content),
				Created:          time.Now(),
				IsActive:         isActive,
				Attachments:      m.Attachments,
				Metadata: map[string]string{
					"message": fmt.Sprintf("%v", i),
					"total":   fmt.Sprintf("%v", len(parsedContent)),
				},
			})

			// Attempt to parse a trade.
			tradeStrategy, err := parser.Parse(ctx, discordproto.DiscordMoonSwingGroupChannel, pc.Content, mc, tradeengineproto.ACTOR_TYPE_EXTERNAL)
			if err != nil {
				// No trade strategy can be parsed; so lets continue.
				slog.Trace(ctx, "Failed to parse trade strategy: %+v, content: %s", err, pc.Content)
				continue
			}

			now := time.Now().UTC()
			idempotencyKey := fmt.Sprintf("tradestrategy-%s-%s-%v-%v-%v", tradeStrategy.ActorId, tradeStrategy.Asset, entriesAsString(tradeStrategy.Entries), tradeStrategy.StopLoss, now.Truncate(time.Minute))

			// Sign our trade strategy with an idempotency key.
			tradeStrategy.IdempotencyKey = idempotencyKey

			rsp, err := createTradeStrategy(ctx, tradeStrategy)
			if err != nil {
				// Best effort for now.
				slog.Error(ctx, "Failed to create tradeStrategy: %v, Error: %v", tradeStrategy, err)
				continue
			}

			tradeStrategy.TradeStrategyId = rsp.TradeStrategyId
			tradeStrategy.Created = rsp.Created
			tradeContent := formatter.FormatTradeStrategy("SWINGS & SCALPS", tradeStrategy, pc.Content)

			msg := &ConsumerMessage{
				ConsumerID:       discordConsumerID,
				DiscordChannelID: discordproto.DiscordSatoshiModTradesChannel,
				Created:          time.Now(),
				Message:          tradeContent,
				Attachments:      m.Attachments,
				IsActive:         isActive,
				Metadata: map[string]string{
					"message":           fmt.Sprintf("%v", i),
					"total":             fmt.Sprintf("%v", len(parsedContent)),
					"trade_strategy_id": tradeStrategy.TradeStrategyId,
				},
				Poller: func(ctx context.Context, messageID string) error {
					// Inject the trade ID.
					return startTradeParticipantsPoller(ctx, messageID, tradeStrategy.TradeStrategyId)
				},
			}

			select {
			case c <- msg:
			default:
				slog.Warn(ctx, "Failed to publish satoshi swings msg; blocked channel")
			}
		}

		// Lets publish our messages.
		for _, msg := range msgs {
			select {
			case c <- msg:
			default:
				slog.Warn(ctx, "Failed to publish satoshi mods msg; blocked channel")
			}

		}
	}
}

func handleInternalCallsMessages(
	ctx context.Context, c chan *ConsumerMessage, isActive bool,
) func(s *discordgo.Session, m *discordgo.MessageCreate) {
	return func(s *discordgo.Session, mc *discordgo.MessageCreate) {
		m := mc.Message

		if m.ChannelID != discordproto.DiscordSatoshiInternalCallsChannel {
			return
		}

		parsedContent, err := getLatestChannelMessages(ctx, s, m.ChannelID)
		if err != nil {
			slog.Error(ctx, "Failed to get latest mod message: %v", err)
			return
		}

		msgs := []*ConsumerMessage{}
		for i, pc := range parsedContent {
			// Attempt to parse a trade.
			tradeStrategy, err := parser.Parse(ctx, discordproto.DiscordMoonSwingGroupChannel, pc.Content, mc, tradeengineproto.ACTOR_TYPE_INTERNAL)
			if err != nil {
				// No trade can be parsed; so lets continue.
				slog.Trace(ctx, "Failed to parse trade: %+v, content: %s", err, pc.Content)
				return
			}

			now := time.Now().UTC()
			idempotencyKey := fmt.Sprintf("tradestrategy-%s-%s-%v-%v-%v", tradeStrategy.ActorId, tradeStrategy.Asset, entriesAsString(tradeStrategy.Entries), tradeStrategy.StopLoss, now.Truncate(time.Minute))

			// Sign our trade strategy with the timestamp.
			tradeStrategy.IdempotencyKey = idempotencyKey

			rsp, err := createTradeStrategy(ctx, tradeStrategy)
			if err != nil {
				// Best effort for now.
				slog.Error(ctx, "Failed to create trade: %v, Error: %v", tradeStrategy, err)
				continue
			}

			tradeStrategy.TradeStrategyId = rsp.TradeStrategyId
			tradeStrategy.Created = rsp.Created
			tradeContent := formatter.FormatTradeStrategy("SCG INTERNAL CALL", tradeStrategy, pc.Content)

			msg := &ConsumerMessage{
				ConsumerID:       discordConsumerID,
				DiscordChannelID: discordproto.DiscordSatoshiModTradesChannel,
				Created:          time.Now(),
				Message:          tradeContent,
				Attachments:      m.Attachments,
				IsActive:         isActive,
				Metadata: map[string]string{
					"trade_strategy_id": tradeStrategy.TradeStrategyId,
					"message":           fmt.Sprintf("%v", i),
					"total":             fmt.Sprintf("%v", len(parsedContent)),
				},
				Poller: func(ctx context.Context, messageID string) error {
					// Inject the trade ID.
					return startTradeParticipantsPoller(ctx, messageID, tradeStrategy.TradeStrategyId)
				},
			}

			msgs = append(msgs, msg)
		}

		for _, msg := range msgs {
			select {
			case c <- msg:
			default:
				slog.Warn(ctx, "Failed to publish satoshi swings msg; blocked channel")
			}
		}
	}
}

type channelMessage struct {
	ID        string                `json:"id"`
	ChannelID string                `json:"channel_id"`
	Author    *channelMessageAuthor `json:"author"`
	Content   string                `json:"content"`
	Timestamp string                `json:"timestamp"`
}

type channelMessageAuthor struct {
	Username string `json:"username"`
}

func getLatestChannelMessages(ctx context.Context, s *discordgo.Session, channelID string) ([]*channelMessage, error) {
	// TODO: A hack for testing purposes.
	if s == nil {
		return nil, nil
	}

	url := fmt.Sprintf(discordChannelUrl, channelID)
	slogParams := map[string]string{
		"channel_id": channelID,
		"url":        url,
	}

	// Create request
	req, err := http.NewRequestWithContext(
		ctx, "GET", url, nil,
	)
	if err != nil {
		return nil, terrors.Augment(err, "Failed to create request", slogParams)
	}

	req.Header.Set("authorization", discordConsumerToken)

	// Execute request.
	rsp, err := s.Client.Do(req)
	if err != nil {
		return nil, terrors.Augment(err, "Failed to make request", slogParams)
	}

	defer rsp.Body.Close()

	// Parse Body.
	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, terrors.Augment(err, "Failed to read response.", slogParams)
	}

	var msgList []*channelMessage
	json.Unmarshal([]byte(body), &msgList)

	return msgList, nil
}

func containsAstekz1To10kChallenge(modUsername string, content string) bool {
	var (
		contains1k     bool
		contains10k    bool
		containsAstekz bool
	)

	if strings.Contains(strings.ToLower(modUsername), "astekz") {
		containsAstekz = true
	}

	tokens := strings.Fields(content + modUsername)
	for _, token := range tokens {
		if contains1k && contains10k && containsAstekz {
			return true
		}

		token := strings.ToLower(token)
		if strings.Contains(token, "1k") {
			contains1k = true
		}

		if strings.Contains(token, "10k") {
			contains10k = true
		}
	}

	return contains1k && contains10k && containsAstekz
}

func containsRego1To10kChallenge(modUsername, content string) bool {
	var (
		contains1k   bool
		contains10k  bool
		containsRego bool
	)

	if strings.Contains(strings.ToLower(modUsername), "rego") {
		containsRego = true
	}

	tokens := strings.Fields(content + modUsername)
	for _, token := range tokens {
		if contains1k && contains10k && containsRego {
			return true
		}

		token := strings.ToLower(token)
		if strings.Contains(token, "1k") {
			contains1k = true
		}

		if strings.Contains(token, "10k") {
			contains10k = true
		}
	}

	return containsRego && contains1k && contains10k
}

// Formats a message for a standardized warning.
func warning(greeting, content string) string {
	return fmt.Sprintf(":rotating_light: %s :rotating_light:\n```%s```", greeting, content)
}

func entriesAsString(ff []float32) string {
	var ss []string
	for _, f := range ff {
		ss = append(ss, fmt.Sprintf("%.5f", f))
	}
	return strings.Join(ss, ",")
}
