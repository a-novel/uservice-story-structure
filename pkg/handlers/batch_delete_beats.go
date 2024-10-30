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

const BatchDeleteBeatsServiceName = "batch_delete_beats"

type BatchDeleteBeats interface {
	storystructurev1grpc.BatchDeleteBeatsServiceServer
}

type batchDeleteBeatsImpl struct {
	service services.BatchDeleteBeats
	logger  adapters.GRPC

	report func(context.Context, *storystructurev1.BatchDeleteBeatsServiceExecRequest) (*emptypb.Empty, error)
}

var handleBatchDeleteBeatsError = grpc.HandleError(codes.Internal).
	Is(services.ErrInvalidBatchDeleteBeatsRequest, codes.InvalidArgument).
	Handle

func (handler *batchDeleteBeatsImpl) exec(
	ctx context.Context,
	request *storystructurev1.BatchDeleteBeatsServiceExecRequest,
) (*emptypb.Empty, error) {
	err := handler.service.Exec(ctx, &services.BatchDeleteBeatsRequest{
		IDs:       request.GetIds(),
		CreatorID: request.GetCreatorId(),
	})
	if err != nil {
		return nil, handleBatchDeleteBeatsError(err)
	}

	return new(emptypb.Empty), nil
}

func (handler *batchDeleteBeatsImpl) Exec(
	ctx context.Context,
	request *storystructurev1.BatchDeleteBeatsServiceExecRequest,
) (*emptypb.Empty, error) {
	return handler.report(ctx, request)
}

func NewBatchDeleteBeats(service services.BatchDeleteBeats, logger adapters.GRPC) BatchDeleteBeats {
	handler := &batchDeleteBeatsImpl{service: service, logger: logger}
	handler.report = adapters.WrapGRPCCall(BatchDeleteBeatsServiceName, logger, handler.exec)
	return handler
}
