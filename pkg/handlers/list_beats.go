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

const ListBeatsServiceName = "list_beats"

type ListBeats interface {
	storystructurev1grpc.ListBeatsServiceServer
}

type listBeatsImpl struct {
	service services.ListBeats
	logger  adapters.GRPC

	report func(
		context.Context, *storystructurev1.ListBeatsServiceExecRequest,
	) (*storystructurev1.ListBeatsServiceExecResponse, error)
}

var handleListBeatsError = grpc.HandleError(codes.Internal).
	Is(services.ErrInvalidListBeatsRequest, codes.InvalidArgument).
	Handle

func beatToListElementProto(
	item *services.ListBeatsResponseBeat, _ int,
) *storystructurev1.ListBeatsServiceExecResponseElement {
	return &storystructurev1.ListBeatsServiceExecResponseElement{
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

func (handler *listBeatsImpl) exec(
	ctx context.Context,
	request *storystructurev1.ListBeatsServiceExecRequest,
) (*storystructurev1.ListBeatsServiceExecResponse, error) {
	res, err := handler.service.Exec(ctx, &services.ListBeatsRequest{IDs: request.GetIds()})
	if err != nil {
		return nil, handleListBeatsError(err)
	}

	beats := lo.Map(res.Beats, beatToListElementProto)

	return &storystructurev1.ListBeatsServiceExecResponse{Beats: beats}, nil
}

func (handler *listBeatsImpl) Exec(
	ctx context.Context,
	request *storystructurev1.ListBeatsServiceExecRequest,
) (*storystructurev1.ListBeatsServiceExecResponse, error) {
	return handler.report(ctx, request)
}

func NewListBeats(service services.ListBeats, logger adapters.GRPC) ListBeats {
	handler := &listBeatsImpl{service: service, logger: logger}
	handler.report = adapters.WrapGRPCCall(ListBeatsServiceName, logger, handler.exec)
	return handler
}
