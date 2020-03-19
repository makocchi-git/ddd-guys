package user

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	pb "github.com/jupemara/ddd-guys/go/adapter/controller/grpc/user/register"
	"github.com/jupemara/ddd-guys/go/usecase/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GrpcUserRegisterController struct {
	pb.UnimplementedGrpcUserRegisterControllerServer
	usecase *user.UserRegister
}

func NewRegister(usecase *user.UserRegister) *GrpcUserRegisterController {
	return &GrpcUserRegisterController{usecase: usecase}
}

func (c *GrpcUserRegisterController) Execute(ctx context.Context, req *pb.User) (*empty.Empty, error) {
	err := c.usecase.Execute(req.GetFirstname(), req.GetLastname())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid argument")
	}
	return &empty.Empty{}, nil
}
