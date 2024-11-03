package handlers

import (
	"context"

	"github.com/samber/lo"
	"google.golang.org/grpc/codes"

	"buf.build/gen/go/a-novel/proto/grpc/go/storystructure/v1/storystructurev1grpc"
	commonv1 "buf.build/gen/go/a-novel/proto/protocolbuffers/go/common/v1"
	storystructurev1 "buf.build/gen/go/a-novel/proto/protocolbuffers/go/storystructure/v1"

	"github.com/a-novel/golib/database"
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
	logger  adapters.GRPC

	report func(
		context.Context, *storystructurev1.SearchBeatsServiceExecRequest,
	) (*storystructurev1.SearchBeatsServiceExecResponse, error)
}

var handleSearchBeatsError = grpc.HandleError(codes.Internal).
	Is(services.ErrInvalidSearchBeatsRequest, codes.InvalidArgument).
	Handle

func (handler *searchBeatsImpl) exec(
	ctx context.Context,
	request *storystructurev1.SearchBeatsServiceExecRequest,
) (*storystructurev1.SearchBeatsServiceExecResponse, error) {
	res, err := handler.service.Exec(
		ctx,
		&services.SearchBeatsRequest{
			Limit:      int(request.GetPagination().GetLimit()),
			Offset:     int(request.GetPagination().GetOffset()),
			CreatorIDs: request.GetCreatorIds(),
			Sort: lo.Switch[storystructurev1.SortBeats, entities.SortBeat](request.GetOrderBy()).
				Case(storystructurev1.SortBeats_SORT_BEATS_BY_NAME, entities.SortBeatName).
				Case(storystructurev1.SortBeats_SORT_BEATS_BY_CREATED_AT, entities.SortBeatCreatedAt).
				Case(storystructurev1.SortBeats_SORT_BEATS_BY_UPDATED_AT, entities.SortBeatUpdatedAt).
				Default(entities.SortBeatNone),
			SortDirection: lo.Switch[commonv1.SortDirection, database.SortDirection](request.GetOrderDirection()).
				Case(commonv1.SortDirection_SORT_DIRECTION_ASC, database.SortDirectionAsc).
				Case(commonv1.SortDirection_SORT_DIRECTION_DESC, database.SortDirectionDesc).
				Default(database.SortDirectionNone),
		},
	)
	if err != nil {
		return nil, handleSearchBeatsError(err)
	}

	return &storystructurev1.SearchBeatsServiceExecResponse{Ids: res.IDs}, nil
}

func (handler *searchBeatsImpl) Exec(
	ctx context.Context,
	request *storystructurev1.SearchBeatsServiceExecRequest,
) (*storystructurev1.SearchBeatsServiceExecResponse, error) {
	return handler.report(ctx, request)
}

func NewSearchBeats(service services.SearchBeats, logger adapters.GRPC) SearchBeats {
	handler := &searchBeatsImpl{service: service, logger: logger}
	handler.report = adapters.WrapGRPCCall(SearchBeatsServiceName, logger, handler.exec)
	return handler
}
