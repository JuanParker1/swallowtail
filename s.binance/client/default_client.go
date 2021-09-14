package client

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"swallowtail/libraries/gerrors"
	"swallowtail/libraries/transport"
	"swallowtail/s.binance/client/signer"
	"swallowtail/s.binance/domain"
	"time"

	"github.com/monzo/slog"
	"github.com/monzo/terrors"
)

type binanceClient struct {
	http transport.HttpClient
}

func (c *binanceClient) GetLatestPrice(ctx context.Context, req *GetLatestPriceRequest) (*GetLatestPriceResponse, error) {
	url := fmt.Sprintf("%s/%s", binanceFuturesURL, "ticker/price")
	rspBody := &GetLatestPriceResponse{}

	qs := fmt.Sprintf("symbol=%s", req.Symbol)

	if err := c.do(ctx, http.MethodGet, url, qs, nil, rspBody, defaultCredentials); err != nil {
		return nil, gerrors.Augment(err, "failed_to_get_latest_price", nil)
	}

	return rspBody, nil
}

func (c *binanceClient) ListAllAssetPairs(ctx context.Context) (*ListAllAssetPairsResponse, error) {
	url := fmt.Sprintf("%s/%s", binanceAPIUrl, "exchangeInfo")
	rspBody := &ListAllAssetPairsResponse{}

	if err := c.http.Do(ctx, http.MethodGet, url, nil, rspBody); err != nil {
		return nil, terrors.Augment(err, "Failed to list all asset pairs", nil)
	}

	return rspBody, nil
}

func (c *binanceClient) ExecuteSpotTrade(ctx context.Context, trade *domain.Trade) error {
	return gerrors.Unimplemented("unimplemented.execute_spot_trade", nil)
}

func (c *binanceClient) ReadSpotAccount(ctx context.Context, in *ReadSpotAccountRequest) (*ReadSpotAccountResponse, error) {
	return nil, gerrors.Unimplemented("unimplemented.read_spot_account", nil)
}

func (c *binanceClient) ReadPerpetualFuturesAccount(ctx context.Context, _ *ReadPerpetualFuturesAccountRequest, credentials *Credentials) (*ReadPerpetualFuturesAccountResponse, error) {
	url := fmt.Sprintf("%s/%s", binanceFuturesURLV2, "balance")
	rspBody := &ReadPerpetualFuturesAccountResponse{}

	if err := c.doWithSignature(ctx, http.MethodGet, url, "", nil, rspBody, credentials); err != nil {
		return nil, gerrors.Augment(err, "failed_to_read_perpetual_futures_account.client", nil)
	}

	return rspBody, nil
}

func (c *binanceClient) ExecutePerpetualFuturesTrade(ctx context.Context, req *ExecutePerpetualFuturesTradeRequest, credentials *Credentials) (*ExecutePerpetualFuturesTradeResponse, error) {
	url := fmt.Sprintf("%s/%s", binanceFuturesURL, "order")
	rspBody := &ExecutePerpetualFuturesTradeResponse{}

	qs := buildQueryStringFromFuturesPerpetualTrade(req)

	if err := c.doWithSignature(ctx, http.MethodPost, url, qs, nil, rspBody, credentials); err != nil {
		slog.Warn(ctx, "Binance Perpetuals futures trade FAILED: %v", qs)
		return nil, gerrors.Augment(err, "failed_to_read_perpetual_futures_account.client", nil)
	}

	slog.Info(ctx, "Binance Perpetuals futures trade executed: %v", qs)

	return rspBody, nil
}

func (c *binanceClient) Ping(ctx context.Context) error {
	endpoint := fmt.Sprintf("%s/ping", binanceAPIUrl)
	rspBody := &PingResponse{}

	if err := c.http.Do(ctx, http.MethodGet, endpoint, nil, rspBody); err != nil {
		return terrors.Augment(err, "Failed to connect to the Binance API.", nil)
	}

	return nil
}

func (c *binanceClient) VerifyCredentials(ctx context.Context, credentials *Credentials) (*VerifyCredentialsResponse, error) {
	endpoint := fmt.Sprintf("%s/account/apiRestrictions", binanceSpotURL)
	rspBody := &VerifyCredentialsResponse{}

	if err := c.doWithSignature(ctx, http.MethodGet, endpoint, "", nil, rspBody, credentials); err != nil {
		return nil, gerrors.Augment(err, "client_request_failed.verify_credentials", map[string]string{
			"endpoint": endpoint,
		})
	}

	return rspBody, nil
}

func (c *binanceClient) GetStatus(ctx context.Context) (*GetStatusResponse, error) {
	endpoint := fmt.Sprintf("%s/time", binanceAPIUrl)
	rspBody := &GetStatusResponse{}

	if err := c.http.Do(ctx, http.MethodGet, endpoint, nil, rspBody); err != nil {
		return nil, gerrors.Augment(err, "client_request_failed.get_status.time", map[string]string{
			"endpoint": endpoint,
		})
	}

	// Convert to millisecond time
	rspBody.ServerTime /= 1_000

	return rspBody, nil
}

func (c *binanceClient) do(ctx context.Context, method, endpoint, queryString string, reqBody, rspBody interface{}, credentials *Credentials) error {
	formattedEndpoint := fmt.Sprintf("%s?%s", endpoint, queryString)
	if credentials == nil {
		return c.http.Do(ctx, method, formattedEndpoint, reqBody, rspBody)
	}

	return c.http.DoWithEphemeralHeaders(ctx, method, formattedEndpoint, reqBody, rspBody, credentials.AsHeaders())
}

func (c *binanceClient) doWithSignature(ctx context.Context, method, endpoint, queryString string, reqBody, rspBody interface{}, credentials *Credentials) error {
	errParams := map[string]string{
		"method":   method,
		"endpoint": endpoint,
	}

	// First check that credentials have indeed been passed correctly.
	switch {
	case credentials == nil:
		return gerrors.FailedPrecondition("cannot_sign_binance_request.nil_credentials", errParams)
	case credentials.SecretKey == "":
		return gerrors.FailedPrecondition("cannot_sign_binance_request.empty_secret_key", errParams)
	}

	// Sign our request with the secret key passed.
	signedRequest, err := c.signRequest(credentials.SecretKey, queryString, reqBody)
	if err != nil {
		return gerrors.Augment(err, "failed_do_request.signature_failure", errParams)
	}

	formattedEndpoint := fmt.Sprintf("%s?%s", endpoint, signedRequest)

	return c.http.DoWithEphemeralHeaders(ctx, method, formattedEndpoint, reqBody, rspBody, credentials.AsHeaders())
}

func (c *binanceClient) signRequest(secret, queryString string, reqBody interface{}) (string, error) {
	// converts to unix nano time to that of millisecond precision; this is all that we need.
	now := time.Now().UnixNano() / 1_000_000

	// sign the request
	hmac, err := signer.Sha256HMAC(secret, queryString, now, reqBody)
	if err != nil {
		return "", gerrors.Augment(err, "failed_to_sign_request", nil)
	}

	// return the new converted querystring with timestamp & signature appended.
	switch {
	case queryString == "":
		return fmt.Sprintf("timestamp=%d&signature=%s", now, hmac), nil
	default:
		return fmt.Sprintf("%s&timestamp=%d&signature=%s", queryString, now, hmac), nil
	}
}

// Binance :)
// TODO: this has to be fixed
func buildQueryStringFromFuturesPerpetualTrade(req *ExecutePerpetualFuturesTradeRequest) string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("symbol=%s", req.Symbol))
	sb.WriteString(fmt.Sprintf("&side=%s", req.Side))
	sb.WriteString(fmt.Sprintf("&type=%s", req.Type))

	if req.PositionSide != "" {
		sb.WriteString(fmt.Sprintf("&positionSide=%s", req.PositionSide))
	}

	if req.TimeInForce != "" {
		sb.WriteString(fmt.Sprintf("&timeInForce=%s", req.TimeInForce))
	}

	if req.Quantity != 0 {
		sb.WriteString(fmt.Sprintf("&quantity=%.1f", req.Quantity))
	}

	if req.ReduceOnly != "" {
		sb.WriteString(fmt.Sprintf("&reduceOnly=%s", req.ReduceOnly))
	}

	if req.Price != 0 {
		sb.WriteString(fmt.Sprintf("&price=%.3f", req.Price))
	}

	if req.NewClientOrderID != "" {
		sb.WriteString(fmt.Sprintf("&newClientOrderId=%s", req.NewClientOrderID))
	}

	if req.StopPrice != 0 {
		sb.WriteString(fmt.Sprintf("&stopPrice=%.3f", req.StopPrice))
	}

	if req.ClosePosition != "" {
		sb.WriteString(fmt.Sprintf("&closePosition=%v", req.ClosePosition))
	}

	if req.ActivationPrice != 0 {
		sb.WriteString(fmt.Sprintf("&activationPrice=%v", req.ActivationPrice))
	}

	if req.CallbackRate != 0 {
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
