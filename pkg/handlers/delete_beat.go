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

const DeleteBeatServiceName = "delete_beat"

type DeleteBeat interface {
	storystructurev1grpc.DeleteBeatServiceServer
}

type deleteBeatImpl struct {
	service services.DeleteBeat
	logger  adapters.GRPC

	report func(context.Context, *storystructurev1.DeleteBeatServiceExecRequest) (*emptypb.Empty, error)
}

var handleDeleteBeatError = grpc.HandleError(codes.Internal).
	Is(services.ErrInvalidDeleteBeatRequest, codes.InvalidArgument).
	Is(dao.ErrBeatNotFound, codes.NotFound).
	Handle

func (handler *deleteBeatImpl) exec(
	ctx context.Context,
	request *storystructurev1.DeleteBeatServiceExecRequest,
) (*emptypb.Empty, error) {
	err := handler.service.Exec(ctx, &services.DeleteBeatRequest{
		ID:        request.GetId(),
		CreatorID: request.GetCreatorId(),
	})
	if err != nil {
		return nil, handleDeleteBeatError(err)
	}

	return new(emptypb.Empty), nil
}

func (handler *deleteBeatImpl) Exec(
	ctx context.Context,
	request *storystructurev1.DeleteBeatServiceExecRequest,
) (*emptypb.Empty, error) {
	return handler.report(ctx, request)
}

func NewDeleteBeat(service services.DeleteBeat, logger adapters.GRPC) DeleteBeat {
	handler := &deleteBeatImpl{service: service, logger: logger}
	handler.report = adapters.WrapGRPCCall(DeleteBeatServiceName, logger, handler.exec)
	return handler
}
