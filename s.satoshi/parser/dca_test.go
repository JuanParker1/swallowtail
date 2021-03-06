package parser

import (
	"context"
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	tradeengineproto "swallowtail/s.trade-engine/proto"
)

func TestDCAParser(t *testing.T) {
	tests := []struct {
		name                  string
		content               string
		username              string
		currentValue          float64
		expectedTradeStrategy *tradeengineproto.TradeStrategy
		withErr               bool
	}{
		{
			name: "internal_case_1",
			content: `Hey guys I'm LONG BTC here.

			ENTRY: 51000-50000
			STOP: 49000

			TP1: 52000
			TP2: 54000
			TP3: 58000

			This should give us an 4.5RR and 15.7% increase
			`,
			username:     "alexjperkins",
			currentValue: 50000,
			expectedTradeStrategy: &tradeengineproto.TradeStrategy{
				HumanizedActorName: "ALEXJPERKINS",
				ActorType:          tradeengineproto.ACTOR_TYPE_EXTERNAL,
				ExecutionStrategy:  tradeengineproto.EXECUTION_STRATEGY_DCA_FIRST_MARKET_REST_LIMIT,
				Asset:              "BTC",
				Pair:               tradeengineproto.TRADE_PAIR_USDT,
				TradeSide:          tradeengineproto.TRADE_SIDE_LONG,
				InstrumentType:     tradeengineproto.INSTRUMENT_TYPE_FUTURE_PERPETUAL,
				Entries:            []float32{50000, 51000},
				StopLoss:           49000,
				CurrentPrice:       50000,
				TakeProfits: []float32{
					52000, 54000, 58000,
				},
				TradeableVenues: []tradeengineproto.VENUE{
					tradeengineproto.VENUE_BINANCE,
				},
			},
		},
		{
			name:         "lrc_example_1",
			content:      `Lrc long 0.404 DCA till 0.395 SL 0.38 TP 0.5 , (50%) and moon bag everyone manage risks`,
			username:     "tahervag",
			currentValue: 0.420,
			expectedTradeStrategy: &tradeengineproto.TradeStrategy{
				HumanizedActorName: "TAHERVAG",
				ActorType:          tradeengineproto.ACTOR_TYPE_EXTERNAL,
				ExecutionStrategy:  tradeengineproto.EXECUTION_STRATEGY_DCA_ALL_LIMIT,
				Asset:              "LRC",
				Pair:               tradeengineproto.TRADE_PAIR_USDT,
				TradeSide:          tradeengineproto.TRADE_SIDE_LONG,
				InstrumentType:     tradeengineproto.INSTRUMENT_TYPE_FUTURE_PERPETUAL,
				Entries:            []float32{0.395, 0.404},
				StopLoss:           0.38,
				CurrentPrice:       0.420,
				TakeProfits: []float32{
					0.5,
				},
				TradeableVenues: []tradeengineproto.VENUE{
					tradeengineproto.VENUE_BINANCE,
				},
			},
		},
		{
			name:         "eli_srm_example_1",
			content:      `Long srm area 8.08 8 stop 7.80 everyone`,
			username:     "eli",
			currentValue: 8.10,
			expectedTradeStrategy: &tradeengineproto.TradeStrategy{
				HumanizedActorName: "ELI",
				ActorType:          tradeengineproto.ACTOR_TYPE_EXTERNAL,
				ExecutionStrategy:  tradeengineproto.EXECUTION_STRATEGY_DCA_FIRST_MARKET_REST_LIMIT,
				Asset:              "SRM",
				Pair:               tradeengineproto.TRADE_PAIR_USDT,
				TradeSide:          tradeengineproto.TRADE_SIDE_LONG,
				InstrumentType:     tradeengineproto.INSTRUMENT_TYPE_FUTURE_PERPETUAL,
				Entries:            []float32{8.0, 8.08},
				StopLoss:           7.80,
				CurrentPrice:       8.1,
				TakeProfits:        []float32{},
				TradeableVenues: []tradeengineproto.VENUE{
					tradeengineproto.VENUE_BINANCE,
				},
			},
		},
		{
			name: "johnny_short_link",
			content: `
			LINK LIMIT SHORT $27 - $27.25 

			SL $27.66 everyone
			`,
			username:     "cryptogodjohn",
			currentValue: 26.0,
			expectedTradeStrategy: &tradeengineproto.TradeStrategy{
				HumanizedActorName: "CRYPTOGODJOHN",
				ActorType:          tradeengineproto.ACTOR_TYPE_EXTERNAL,
				ExecutionStrategy:  tradeengineproto.EXECUTION_STRATEGY_DCA_ALL_LIMIT,
				Asset:              "LINK",
				Pair:               tradeengineproto.TRADE_PAIR_USDT,
				TradeSide:          tradeengineproto.TRADE_SIDE_SHORT,
				InstrumentType:     tradeengineproto.INSTRUMENT_TYPE_FUTURE_PERPETUAL,
				Entries:            []float32{27.25, 27.0},
				StopLoss:           27.66,
				CurrentPrice:       26.0,
				TakeProfits:        []float32{},
				TradeableVenues: []tradeengineproto.VENUE{
					tradeengineproto.VENUE_BINANCE,
				},
			},
		},
		{
			name:         "johnny_short_link_missing_entry",
			currentValue: 26.0,
			content: `
			LINK LIMIT SHORT $27

			SL $27.66 everyone
			`,
			withErr: true,
		},
		{
			name:         "johnny_short_link_one_invalid_entry",
			currentValue: 26.0,
			content: `
			LINK LIMIT SHORT 3 $27

			SL $27.66 everyone
			`,
			withErr: true,
		},
		{
			name: "swings_single_entry",
			content: `
			entry now 3502
			Stop 3407
			target 3704.7 ETH
			`,
			withErr: true,
		},
		{
			name: "swings_full_example",
			content: `
bluntz/hfsp [crypto-scalp-trade-ideas-89]: trade idea scalp long btc/usd   scalp longiong btc again here on this little 15min timeframe dip we are still reduced to scalps at these levels as we need larger dips to position for any kind of longer timeframe core longs plus we are already long btc from 40k   entry 66100  stop 65148  target 68669   2.58rr   everyone  [attachments]  https://cdn.discordapp.com/attachments/671977297829429255/900478428711624775/unknown.png
			`,
			withErr: true,
		},
	}

	originalBinanceAssetPairs := binanceInstruments
	binanceInstruments = map[string]bool{
		"btc":  true,
		"lrc":  true,
		"srm":  true,
		"link": true,
		"eth":  true,
	}

	originalFetcher := fetchLatestPrice
	t.Cleanup(func() {
		binanceInstruments = originalBinanceAssetPairs
		fetchLatestPrice = originalFetcher
	})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fetchLatestPrice = func(_ context.Context, _ string) (float64, error) {
				return tt.currentValue, nil
			}

			trade, err := (&DCAParser{}).Parse(context.Background(), tt.content, &discordgo.MessageCreate{
				Message: &discordgo.Message{
					Author: &discordgo.User{
						Username: tt.username,
					},
				},
			}, tradeengineproto.ACTOR_TYPE_EXTERNAL)

			switch {
			case !tt.withErr:
				require.NoError(t, err)
				assert.Equal(t, tt.expectedTradeStrategy, trade)
			default:
				require.Error(t, err)
			}
		})
	}
}
