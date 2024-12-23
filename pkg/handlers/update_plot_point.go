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

const UpdatePlotPointServiceName = "update_plot_point"

type UpdatePlotPoint interface {
	storystructurev1grpc.UpdatePlotPointServiceServer
}

type updatePlotPointImpl struct {
	service services.UpdatePlotPoint
}

var handleUpdatePlotPointError = grpc.HandleError(codes.Internal).
	Is(services.ErrInvalidUpdatePlotPointRequest, codes.InvalidArgument).
	Is(dao.ErrPlotPointNotFound, codes.NotFound).
	Handle

func (handler *updatePlotPointImpl) Exec(
	ctx context.Context,
	request *storystructurev1.UpdatePlotPointServiceExecRequest,
) (*storystructurev1.UpdatePlotPointServiceExecResponse, error) {
	res, err := handler.service.Exec(
		ctx,
		&services.UpdatePlotPointRequest{
			ID:        request.GetId(),
			Name:      request.GetName(),
			Prompt:    request.GetPrompt(),
			CreatorID: request.GetCreatorId(),
		},
	)
	if err != nil {
		return nil, handleUpdatePlotPointError(err)
	}

	return &storystructurev1.UpdatePlotPointServiceExecResponse{
		Id:        res.ID,
		Name:      res.Name,
		Prompt:    res.Prompt,
		CreatorId: res.CreatorID,
		CreatedAt: timestamppb.New(res.CreatedAt),
		UpdatedAt: grpc.TimestampOptional(res.UpdatedAt),
	}, nil
}

func NewUpdatePlotPoint(service services.UpdatePlotPoint, logger adapters.GRPC) UpdatePlotPoint {
	handler := &updatePlotPointImpl{service: service}
	return grpc.ServiceWithMetrics(UpdatePlotPointServiceName, handler, logger)
}
