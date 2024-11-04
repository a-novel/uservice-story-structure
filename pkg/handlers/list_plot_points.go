package handlers

import (
	"context"

	"github.com/samber/lo"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/timestamppb"

	"buf.build/gen/go/a-novel/proto/grpc/go/storystructure/v1/storystructurev1grpc"
	storystructurev1 "buf.build/gen/go/a-novel/proto/protocolbuffers/go/storystructure/v1"

	"github.com/a-novel/golib/grpc"
	"github.com/a-novel/golib/loggers/adapters"

	"github.com/a-novel/uservice-story-structure/pkg/services"
)

const ListPlotPointsServiceName = "list_plot_points"

type ListPlotPoints interface {
	storystructurev1grpc.ListPlotPointsServiceServer
}

type listPlotPointsImpl struct {
	service services.ListPlotPoints
}

var handleListPlotPointsError = grpc.HandleError(codes.Internal).
	Is(services.ErrInvalidListPlotPointsRequest, codes.InvalidArgument).
	Handle

func plotPointToListElementProto(
	item *services.ListPlotPointsResponsePlotPoint, _ int,
) *storystructurev1.ListPlotPointsServiceExecResponseElement {
	return &storystructurev1.ListPlotPointsServiceExecResponseElement{
		Id:        item.ID,
		Name:      item.Name,
		Prompt:    item.Prompt,
		CreatorId: item.CreatorID,
		CreatedAt: timestamppb.New(item.CreatedAt),
		UpdatedAt: lo.TernaryF(
			item.UpdatedAt == nil,
			func() *timestamppb.Timestamp { return nil },
			func() *timestamppb.Timestamp { return timestamppb.New(*item.UpdatedAt) },
		),
	}
}

func (handler *listPlotPointsImpl) Exec(
	ctx context.Context,
	request *storystructurev1.ListPlotPointsServiceExecRequest,
) (*storystructurev1.ListPlotPointsServiceExecResponse, error) {
	res, err := handler.service.Exec(ctx, &services.ListPlotPointsRequest{IDs: request.GetIds()})
	if err != nil {
		return nil, handleListPlotPointsError(err)
	}

	beats := lo.Map(res.PlotPoints, plotPointToListElementProto)

	return &storystructurev1.ListPlotPointsServiceExecResponse{PlotPoints: beats}, nil
}

func NewListPlotPoints(service services.ListPlotPoints, logger adapters.GRPC) ListPlotPoints {
	handler := &listPlotPointsImpl{service: service}
	return grpc.ServiceWithMetrics(ListPlotPointsServiceName, handler, logger)
}
