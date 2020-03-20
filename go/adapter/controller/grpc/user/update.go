package user

import (
	"context"

	pb "github.com/jupemara/ddd-guys/go/adapter/controller/grpc/user/update"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jupemara/ddd-guys/go/usecase/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GrpcUserUpdateController struct {
	pb.UnimplementedGrpcUserUpdateControllerServer
	usecase *user.UserUpdateUsecase
}

func NewUpdate(usecase *user.UserUpdateUsecase) *GrpcUserUpdateController {
	return &GrpcUserUpdateController{usecase: usecase}
}

func (c *GrpcUserUpdateController) Execute(ctx context.Context, req *pb.User) (*empty.Empty, error) {
	command := user.NewCommand(req.GetId(), req.GetFirstname(), req.GetLastname())
	err := c.usecase.Execute(command)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update user")
	}
	return &empty.Empty{}, nil
}
