package handlers

import (
	"context"

	"google.golang.org/grpc/codes"

	"buf.build/gen/go/a-novel/proto/grpc/go/storystructure/v1/storystructurev1grpc"
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
}

var handleSearchPlotPointsError = grpc.HandleError(codes.Internal).
	Is(services.ErrInvalidSearchPlotPointsRequest, codes.InvalidArgument).
	Handle

func (handler *searchPlotPointsImpl) Exec(
	ctx context.Context,
	request *storystructurev1.SearchPlotPointsServiceExecRequest,
) (*storystructurev1.SearchPlotPointsServiceExecResponse, error) {
	res, err := handler.service.Exec(
		ctx,
		&services.SearchPlotPointsRequest{
			Limit:         int(request.GetPagination().GetLimit()),
			Offset:        int(request.GetPagination().GetOffset()),
			CreatorIDs:    request.GetCreatorIds(),
			Sort:          entities.SortPlotPointConverter.FromProto(request.GetOrderBy()),
			SortDirection: grpc.SortDirectionConverter.FromProto(request.GetOrderDirection()),
		},
	)
	if err != nil {
		return nil, handleSearchPlotPointsError(err)
	}

	return &storystructurev1.SearchPlotPointsServiceExecResponse{Ids: res.IDs}, nil
}

func NewSearchPlotPoints(service services.SearchPlotPoints, logger adapters.GRPC) SearchPlotPoints {
	handler := &searchPlotPointsImpl{service: service}
	return grpc.ServiceWithMetrics(SearchPlotPointsServiceName, handler, logger)
}
