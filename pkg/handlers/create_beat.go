package handlers

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/timestamppb"

	"buf.build/gen/go/a-novel/proto/grpc/go/storystructure/v1/storystructurev1grpc"
	storystructurev1 "buf.build/gen/go/a-novel/proto/protocolbuffers/go/storystructure/v1"

	"github.com/a-novel/golib/grpc"
	"github.com/a-novel/golib/loggers/adapters"

	"github.com/a-novel/uservice-story-structure/pkg/services"
)

const CreateBeatServiceName = "create_beat"

type CreateBeat interface {
	storystructurev1grpc.CreateBeatServiceServer
}

type createBeatImpl struct {
	service services.CreateBeat
}

var handleCreateBeatError = grpc.HandleError(codes.Internal).
	Is(services.ErrInvalidCreateBeatRequest, codes.InvalidArgument).
	Handle

func (handler *createBeatImpl) Exec(
	ctx context.Context,
	request *storystructurev1.CreateBeatServiceExecRequest,
) (*storystructurev1.CreateBeatServiceExecResponse, error) {
	res, err := handler.service.Exec(ctx, &services.CreateBeatRequest{
		Name:      request.GetName(),
		Prompt:    request.GetPrompt(),
		CreatorID: request.GetCreatorId(),
	})
	if err != nil {
		return nil, handleCreateBeatError(err)
	}

	return &storystructurev1.CreateBeatServiceExecResponse{
		Id:        res.ID,
		CreatorId: res.CreatorID,
		Name:      res.Name,
		Prompt:    res.Prompt,
		CreatedAt: timestamppb.New(res.CreatedAt),
	}, nil
}

func NewCreateBeat(service services.CreateBeat, logger adapters.GRPC) CreateBeat {
	handler := &createBeatImpl{service: service}
	return grpc.ServiceWithMetrics(CreateBeatServiceName, handler, logger)
}
