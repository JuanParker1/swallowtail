package commands

import (
	"context"
	"fmt"

	"github.com/bwmarrin/discordgo"

	"swallowtail/libraries/gerrors"
	satoshiproto "swallowtail/s.satoshi/proto"
)

const (
	pingCommandID    = "ping"
	pingCommandUsage = `!ping`
)

func init() {
	register(pingCommandID, &Command{
		ID:          pingCommandID,
		Usage:       pingCommandUsage,
		Description: "Checks check if satoshi bot is live. Also gives satoshis version number.",
		Handler:     pingCommand,
	})
}

func pingCommand(ctx context.Context, tokens []string, s *discordgo.Session, m *discordgo.MessageCreate) error {
	if _, err := s.ChannelMessageSend(m.ChannelID, fmt.Sprintf(":wave: Hello <@%s>, what can `sat %s` do for you?", m.Author.ID, satoshiproto.SatoshiVersion)); err != nil {
		return gerrors.Augment(err, "failed_to_ping_discord", nil)
	}

	return nil
}
