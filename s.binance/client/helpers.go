package client

import (
	"fmt"
	"strings"
)

// Binance :)
// TODO: this has to be fixed
func buildQueryStringFromFuturesPerpetualTrade(req *ExecutePerpetualFuturesOrderRequest) string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("symbol=%s", req.Symbol))
	sb.WriteString(fmt.Sprintf("&side=%s", req.Side))
	sb.WriteString(fmt.Sprintf("&type=%s", req.OrderType))

	if req.PositionSide != "" {
		sb.WriteString(fmt.Sprintf("&positionSide=%s", req.PositionSide))
	}

	if req.TimeInForce != "" {
		sb.WriteString(fmt.Sprintf("&timeInForce=%s", req.TimeInForce))
	}

	if req.Quantity != "" {
		sb.WriteString(fmt.Sprintf("&quantity=%s", req.Quantity))
	}

	if req.ReduceOnly != "" {
		sb.WriteString(fmt.Sprintf("&reduceOnly=%s", req.ReduceOnly))
	}

	if req.LimitPrice != "" {
		sb.WriteString(fmt.Sprintf("&price=%s", req.LimitPrice))
	}

	if req.NewClientOrderID != "" {
		sb.WriteString(fmt.Sprintf("&newClientOrderId=%s", req.NewClientOrderID))
	}

	if req.StopPrice != "" {
		sb.WriteString(fmt.Sprintf("&stopPrice=%s", req.StopPrice))
	}

	if req.ClosePosition != "" {
		sb.WriteString(fmt.Sprintf("&closePosition=%v", req.ClosePosition))
	}

	if req.ActivationPrice != "" {
		sb.WriteString(fmt.Sprintf("&activationPrice=%v", req.ActivationPrice))
	}

	if req.CallbackRate != "" {
		sb.WriteString(fmt.Sprintf("&callbackRate=%v", req.CallbackRate))
	}

	if req.WorkingType != "" {
		sb.WriteString(fmt.Sprintf("&workingType=%s", req.WorkingType))
	}

	if req.PriceProtect != "" {
		sb.WriteString(fmt.Sprintf("&priceProtect=%s", req.PriceProtect))
	}

	if req.NewOrderRespType != "" {
		sb.WriteString(fmt.Sprintf("&newOrderRespType=%s", req.NewOrderRespType))
	}

	return sb.String()
}
