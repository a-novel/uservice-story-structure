package handlers

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/emptypb"

	"buf.build/gen/go/a-novel/proto/grpc/go/storystructure/v1/storystructurev1grpc"
	storystructurev1 "buf.build/gen/go/a-novel/proto/protocolbuffers/go/storystructure/v1"

	"github.com/a-novel/golib/grpc"
	"github.com/a-novel/golib/loggers/adapters"

	"github.com/a-novel/uservice-story-structure/pkg/dao"
	"github.com/a-novel/uservice-story-structure/pkg/services"
)

const DeletePlotPointServiceName = "delete_plot_point"

type DeletePlotPoint interface {
	storystructurev1grpc.DeletePlotPointServiceServer
}

type deletePlotPointImpl struct {
	service services.DeletePlotPoint
}

var handleDeletePlotPointError = grpc.HandleError(codes.Internal).
	Is(services.ErrInvalidDeletePlotPointRequest, codes.InvalidArgument).
	Is(dao.ErrPlotPointNotFound, codes.NotFound).
	Handle

func (handler *deletePlotPointImpl) Exec(
	ctx context.Context,
	request *storystructurev1.DeletePlotPointServiceExecRequest,
) (*emptypb.Empty, error) {
	err := handler.service.Exec(ctx, &services.DeletePlotPointRequest{
		ID:        request.GetId(),
		CreatorID: request.GetCreatorId(),
	})
	if err != nil {
		return nil, handleDeletePlotPointError(err)
	}

	return new(emptypb.Empty), nil
}

func NewDeletePlotPoint(service services.DeletePlotPoint, logger adapters.GRPC) DeletePlotPoint {
	handler := &deletePlotPointImpl{service: service}
	return grpc.ServiceWithMetrics(DeletePlotPointServiceName, handler, logger)
}
