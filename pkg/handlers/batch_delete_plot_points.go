package handlers

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/emptypb"

	"buf.build/gen/go/a-novel/proto/grpc/go/storystructure/v1/storystructurev1grpc"
	storystructurev1 "buf.build/gen/go/a-novel/proto/protocolbuffers/go/storystructure/v1"

	"github.com/a-novel/golib/grpc"
	"github.com/a-novel/golib/loggers/adapters"

	"github.com/a-novel/uservice-story-structure/pkg/services"
)

const BatchDeletePlotPointsServiceName = "batch_delete_plot_points"

type BatchDeletePlotPoints interface {
	storystructurev1grpc.BatchDeletePlotPointsServiceServer
}

type batchDeletePlotPointsImpl struct {
	service services.BatchDeletePlotPoints
}

var handleBatchDeletePlotPointsError = grpc.HandleError(codes.Internal).
	Is(services.ErrInvalidBatchDeletePlotPointsRequest, codes.InvalidArgument).
	Handle

func (handler *batchDeletePlotPointsImpl) Exec(
	ctx context.Context,
	request *storystructurev1.BatchDeletePlotPointsServiceExecRequest,
) (*emptypb.Empty, error) {
	err := handler.service.Exec(ctx, &services.BatchDeletePlotPointsRequest{
		IDs:       request.GetIds(),
		CreatorID: request.GetCreatorId(),
	})
	if err != nil {
		return nil, handleBatchDeletePlotPointsError(err)
	}

	return new(emptypb.Empty), nil
}

func NewBatchDeletePlotPoints(service services.BatchDeletePlotPoints, logger adapters.GRPC) BatchDeletePlotPoints {
	handler := &batchDeletePlotPointsImpl{service: service}
	return grpc.ServiceWithMetrics(BatchDeletePlotPointsServiceName, handler, logger)
}
