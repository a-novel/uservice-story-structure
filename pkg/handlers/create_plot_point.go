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

const CreatePlotPointServiceName = "create_plot_point"

type CreatePlotPoint interface {
	storystructurev1grpc.CreatePlotPointServiceServer
}

type createPlotPointImpl struct {
	service services.CreatePlotPoint
}

var handleCreatePlotPointError = grpc.HandleError(codes.Internal).
	Is(services.ErrInvalidCreatePlotPointRequest, codes.InvalidArgument).
	Handle

func (handler *createPlotPointImpl) Exec(
	ctx context.Context,
	request *storystructurev1.CreatePlotPointServiceExecRequest,
) (*storystructurev1.CreatePlotPointServiceExecResponse, error) {
	res, err := handler.service.Exec(ctx, &services.CreatePlotPointRequest{
		Name:      request.GetName(),
		Prompt:    request.GetPrompt(),
		CreatorID: request.GetCreatorId(),
	})
	if err != nil {
		return nil, handleCreatePlotPointError(err)
	}

	return &storystructurev1.CreatePlotPointServiceExecResponse{
		Id:        res.ID,
		CreatorId: res.CreatorID,
		Name:      res.Name,
		Prompt:    res.Prompt,
		CreatedAt: timestamppb.New(res.CreatedAt),
	}, nil
}

func NewCreatePlotPoint(service services.CreatePlotPoint, logger adapters.GRPC) CreatePlotPoint {
	handler := &createPlotPointImpl{service: service}
	return grpc.ServiceWithMetrics(CreatePlotPointServiceName, handler, logger)
}
