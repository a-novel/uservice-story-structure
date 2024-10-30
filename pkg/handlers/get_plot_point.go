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

	"github.com/a-novel/uservice-story-structure/pkg/dao"
	"github.com/a-novel/uservice-story-structure/pkg/services"
)

const GetPlotPointServiceName = "get_plot_point"

type GetPlotPoint interface {
	storystructurev1grpc.GetPlotPointServiceServer
}

type getPlotPointImpl struct {
	service services.GetPlotPoint
	logger  adapters.GRPC

	report func(
		context.Context, *storystructurev1.GetPlotPointServiceExecRequest,
	) (*storystructurev1.GetPlotPointServiceExecResponse, error)
}

var handleGetPlotPointError = grpc.HandleError(codes.Internal).
	Is(services.ErrInvalidGetPlotPointRequest, codes.InvalidArgument).
	Is(dao.ErrPlotPointNotFound, codes.NotFound).
	Handle

func (handler *getPlotPointImpl) exec(
	ctx context.Context,
	request *storystructurev1.GetPlotPointServiceExecRequest,
) (*storystructurev1.GetPlotPointServiceExecResponse, error) {
	res, err := handler.service.Exec(ctx, &services.GetPlotPointRequest{ID: request.GetId()})
	if err != nil {
		return nil, handleGetPlotPointError(err)
	}

	return &storystructurev1.GetPlotPointServiceExecResponse{
		Id:        res.ID,
		Name:      res.Name,
		Prompt:    res.Prompt,
		CreatorId: res.CreatorID,
		CreatedAt: timestamppb.New(res.CreatedAt),
		UpdatedAt: lo.TernaryF(
			res.UpdatedAt == nil,
			func() *timestamppb.Timestamp { return nil },
			func() *timestamppb.Timestamp { return timestamppb.New(*res.UpdatedAt) },
		),
	}, nil
}

func (handler *getPlotPointImpl) Exec(
	ctx context.Context,
	request *storystructurev1.GetPlotPointServiceExecRequest,
) (*storystructurev1.GetPlotPointServiceExecResponse, error) {
	return handler.report(ctx, request)
}

func NewGetPlotPoint(service services.GetPlotPoint, logger adapters.GRPC) GetPlotPoint {
	handler := &getPlotPointImpl{service: service, logger: logger}
	handler.report = adapters.WrapGRPCCall(GetPlotPointServiceName, logger, handler.exec)
	return handler
}