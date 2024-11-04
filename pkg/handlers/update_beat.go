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

const UpdateBeatServiceName = "update_beat"

type UpdateBeat interface {
	storystructurev1grpc.UpdateBeatServiceServer
}

type updateBeatImpl struct {
	service services.UpdateBeat
}

var handleUpdateBeatError = grpc.HandleError(codes.Internal).
	Is(services.ErrInvalidUpdateBeatRequest, codes.InvalidArgument).
	Is(dao.ErrBeatNotFound, codes.NotFound).
	Handle

func (handler *updateBeatImpl) Exec(
	ctx context.Context,
	request *storystructurev1.UpdateBeatServiceExecRequest,
) (*storystructurev1.UpdateBeatServiceExecResponse, error) {
	res, err := handler.service.Exec(
		ctx,
		&services.UpdateBeatRequest{
			ID:        request.GetId(),
			Name:      request.GetName(),
			Prompt:    request.GetPrompt(),
			CreatorID: request.GetCreatorId(),
		},
	)
	if err != nil {
		return nil, handleUpdateBeatError(err)
	}

	return &storystructurev1.UpdateBeatServiceExecResponse{
		Id:        res.ID,
		Name:      res.Name,
		Prompt:    res.Prompt,
		CreatorId: res.CreatorID,
		CreatedAt: timestamppb.New(res.CreatedAt),
		UpdatedAt: grpc.TimestampOptional(res.UpdatedAt),
	}, nil
}

func NewUpdateBeat(service services.UpdateBeat, logger adapters.GRPC) UpdateBeat {
	handler := &updateBeatImpl{service: service}
	return grpc.ServiceWithMetrics(UpdateBeatServiceName, handler, logger)
}
