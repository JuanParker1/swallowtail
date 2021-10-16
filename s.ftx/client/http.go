package client

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"swallowtail/libraries/gerrors"
	"time"
)

func (f *ftxClient) do(ctx context.Context, method, endpoint string, req, rsp interface{}, pagination *PaginationFilter, credentials *Credentials) error {
	url := fmt.Sprintf("%s%s", f.hostname, buildEndpoint(endpoint, pagination))

	var creds = credentials
	if creds == nil {
		creds = &Credentials{}
	}

	return f.http.DoWithEphemeralHeaders(ctx, method, url, req, rsp, creds.SubaccountAsHeaders())
}

func (f *ftxClient) signBeforeDo(ctx context.Context, method, endpoint string, req, rsp interface{}, pagination *PaginationFilter, credentials *Credentials) error {
	ts := strconv.FormatInt(time.Now().UTC().Unix()*1000, 10)
	preparedEndpoint := buildEndpoint(endpoint, pagination)

	reqBody, err := json.Marshal(req)
	if err != nil {
		return gerrors.Augment(err, "failed_to_sign_request.bad_request_body", nil)
	}

	signature, err := credentials.SignRequest(method, preparedEndpoint, ts, reqBody)
	if err != nil {
		return gerrors.Augment(err, "failed_to_send_request.failed_to_sign_request", nil)
	}

	url := fmt.Sprintf("%s%s", f.hostname, preparedEndpoint)

	return f.http.DoWithEphemeralHeaders(ctx, method, url, req, rsp, credentials.AsHeaders(signature, ts))
}

func buildEndpoint(base string, pagination *PaginationFilter) string {
	var endpoint = base
	if pagination != nil {
		endpoint = fmt.Sprintf("%s?%s", endpoint, pagination.ToQueryString())
	}

	return endpoint
}
