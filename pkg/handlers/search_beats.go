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

const SearchBeatsServiceName = "search_beats"

type SearchBeats interface {
	storystructurev1grpc.SearchBeatsServiceServer
}

type searchBeatsImpl struct {
	service services.SearchBeats
}

var handleSearchBeatsError = grpc.HandleError(codes.Internal).
	Is(services.ErrInvalidSearchBeatsRequest, codes.InvalidArgument).
	Handle

func (handler *searchBeatsImpl) Exec(
	ctx context.Context,
	request *storystructurev1.SearchBeatsServiceExecRequest,
) (*storystructurev1.SearchBeatsServiceExecResponse, error) {
	res, err := handler.service.Exec(
		ctx,
		&services.SearchBeatsRequest{
			Limit:         int(request.GetPagination().GetLimit()),
			Offset:        int(request.GetPagination().GetOffset()),
			CreatorIDs:    request.GetCreatorIds(),
			Sort:          entities.SortBeatConverter.FromProto(request.GetOrderBy()),
			SortDirection: grpc.SortDirectionConverter.FromProto(request.GetOrderDirection()),
		},
	)
	if err != nil {
		return nil, handleSearchBeatsError(err)
	}

	return &storystructurev1.SearchBeatsServiceExecResponse{Ids: res.IDs}, nil
}

func NewSearchBeats(service services.SearchBeats, logger adapters.GRPC) SearchBeats {
	handler := &searchBeatsImpl{service: service}
	return grpc.ServiceWithMetrics(SearchBeatsServiceName, handler, logger)
}
