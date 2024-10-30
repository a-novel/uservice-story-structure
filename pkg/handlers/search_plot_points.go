package handlers

import (
	"context"

	"github.com/samber/lo"
	"google.golang.org/grpc/codes"

	"buf.build/gen/go/a-novel/proto/grpc/go/storystructure/v1/storystructurev1grpc"
	commonv1 "buf.build/gen/go/a-novel/proto/protocolbuffers/go/common/v1"
	storystructurev1 "buf.build/gen/go/a-novel/proto/protocolbuffers/go/storystructure/v1"

	"github.com/a-novel/golib/grpc"
	"github.com/a-novel/golib/loggers/adapters"

	"github.com/a-novel/uservice-story-structure/pkg/entities"
	"github.com/a-novel/uservice-story-structure/pkg/services"
)

const SearchPlotPointsServiceName = "search_plot_points"

type SearchPlotPoints interface {
	storystructurev1grpc.SearchPlotPointsServiceServer
}

type searchPlotPointsImpl struct {
	service services.SearchPlotPoints
	logger  adapters.GRPC

	report func(
		context.Context, *storystructurev1.SearchPlotPointsServiceExecRequest,
	) (*storystructurev1.SearchPlotPointsServiceExecResponse, error)
}

var handleSearchPlotPointsError = grpc.HandleError(codes.Internal).
	Is(services.ErrInvalidSearchPlotPointsRequest, codes.InvalidArgument).
	Handle

func (handler *searchPlotPointsImpl) exec(
	ctx context.Context,
	request *storystructurev1.SearchPlotPointsServiceExecRequest,
) (*storystructurev1.SearchPlotPointsServiceExecResponse, error) {
	res, err := handler.service.Exec(
		ctx,
		&services.SearchPlotPointsRequest{
			Limit:      int(request.GetPagination().GetLimit()),
			Offset:     int(request.GetPagination().GetOffset()),
			CreatorIDs: request.GetCreatorIds(),
			Sort: lo.Switch[storystructurev1.SortPlotPoints, entities.SortPlotPoint](request.GetOrderBy()).
				Case(storystructurev1.SortPlotPoints_SORT_PLOT_POINTS_BY_NAME, entities.SortPlotPointName).
				Case(storystructurev1.SortPlotPoints_SORT_PLOT_POINTS_BY_CREATED_AT, entities.SortPlotPointCreatedAt).
				Case(storystructurev1.SortPlotPoints_SORT_PLOT_POINTS_BY_UPDATED_AT, entities.SortPlotPointUpdatedAt).
				Default(entities.SortPlotPointNone),
			SortDirection: lo.Switch[commonv1.SortDirection, entities.SortDirection](request.GetOrderDirection()).
				Case(commonv1.SortDirection_SORT_DIRECTION_ASC, entities.SortDirectionAsc).
				Case(commonv1.SortDirection_SORT_DIRECTION_DESC, entities.SortDirectionDesc).
				Default(entities.SortDirectionNone),
		},
	)
	if err != nil {
		return nil, handleSearchPlotPointsError(err)
	}

	return &storystructurev1.SearchPlotPointsServiceExecResponse{Ids: res.IDs}, nil
}

func (handler *searchPlotPointsImpl) Exec(
	ctx context.Context,
	request *storystructurev1.SearchPlotPointsServiceExecRequest,
) (*storystructurev1.SearchPlotPointsServiceExecResponse, error) {
	return handler.report(ctx, request)
}

func NewSearchPlotPoints(service services.SearchPlotPoints, logger adapters.GRPC) SearchPlotPoints {
	handler := &searchPlotPointsImpl{service: service, logger: logger}
	handler.report = adapters.WrapGRPCCall(SearchPlotPointsServiceName, logger, handler.exec)
	return handler
}
