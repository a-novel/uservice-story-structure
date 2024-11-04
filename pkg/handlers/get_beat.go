package handlers

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/timestamppb"

	"buf.build/gen/go/a-novel/proto/grpc/go/storystructure/v1/storystructurev1grpc"
	storystructurev1 "buf.build/gen/go/a-novel/proto/protocolbuffers/go/storystructure/v1"

	"github.com/a-novel/golib/grpc"
	"github.com/a-novel/golib/loggers/adapters"

	"github.com/a-novel/uservice-story-structure/pkg/dao"
	"github.com/a-novel/uservice-story-structure/pkg/services"
)

const GetBeatServiceName = "get_beat"

type GetBeat interface {
	storystructurev1grpc.GetBeatServiceServer
}

type getBeatImpl struct {
	service services.GetBeat
}

var handleGetBeatError = grpc.HandleError(codes.Internal).
	Is(services.ErrInvalidGetBeatRequest, codes.InvalidArgument).
	Is(dao.ErrBeatNotFound, codes.NotFound).
	Handle

func (handler *getBeatImpl) Exec(
	ctx context.Context,
	request *storystructurev1.GetBeatServiceExecRequest,
) (*storystructurev1.GetBeatServiceExecResponse, error) {
	res, err := handler.service.Exec(ctx, &services.GetBeatRequest{ID: request.GetId()})
	if err != nil {
		return nil, handleGetBeatError(err)
	}

	return &storystructurev1.GetBeatServiceExecResponse{
		Id:        res.ID,
		Name:      res.Name,
		Prompt:    res.Prompt,
		CreatorId: res.CreatorID,
		CreatedAt: timestamppb.New(res.CreatedAt),
		UpdatedAt: grpc.TimestampOptional(res.UpdatedAt),
	}, nil
}

func NewGetBeat(service services.GetBeat, logger adapters.GRPC) GetBeat {
	handler := &getBeatImpl{service: service}
	return grpc.ServiceWithMetrics(GetBeatServiceName, handler, logger)
}
