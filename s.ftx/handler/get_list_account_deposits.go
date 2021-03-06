package handler

import (
	"context"
	"swallowtail/libraries/gerrors"
	"swallowtail/s.ftx/client"
	"swallowtail/s.ftx/marshaling"
	ftxproto "swallowtail/s.ftx/proto"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// ListAccountDeposits ...
func (s *FTXService) ListAccountDeposits(
	ctx context.Context, in *ftxproto.ListAccountDepositsRequest,
) (*ftxproto.ListAccountDepositsResponse, error) {
	switch {
	// TODO: validate start and end times.
	case in.ActorId == "":
		return nil, gerrors.BadParam("missing_param.actor_id", nil)
	}

	errParams := map[string]string{
		"actor_id": in.ActorId,
	}

	// Check the actor is authorized to make this request.
	ok, err := isValidActor(ctx, in.ActorId)
	if err != nil {
		return nil, gerrors.Augment(err, "failed_to_list_account_deposits", errParams)
	}

	if !ok {
		return nil, gerrors.Unauthenticated("failed_to_list_account_deposits.unauthorized_actor", errParams)
	}

	pagination := buildPaginationFilter(in.Start, in.End)

	rsp, err := client.ListAccountDeposits(ctx, &client.ListAccountDepositsRequest{}, pagination)
	if err != nil {
		return nil, gerrors.Augment(err, "failed_to_list_account_deposits", errParams)
	}

	protoDeposits := marshaling.DepositsDTOToProto(rsp.Deposits)

	return &ftxproto.ListAccountDepositsResponse{
		Deposits: protoDeposits,
	}, nil
}

func buildPaginationFilter(startTime, endTime int64) *client.PaginationFilter {
	if startTime == 0 && endTime == 0 {
		return nil
	}

	if endTime == 0 {
		return &client.PaginationFilter{
			Start: startTime,
			End:   timestamppb.Now().Seconds,
		}
	}

	if startTime == 0 {
		// It's non sensical to have an endtime but no starttime; just pass no pagination filter here rather than error.
		return nil
	}

	return &client.PaginationFilter{
		Start: startTime,
		End:   endTime,
	}
}
