package user

import (
	"context"

	pb "github.com/jupemara/ddd-guys/go/adapter/controller/grpc/user/find"
	"github.com/jupemara/ddd-guys/go/usecase/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GrpcUserFindController struct {
	pb.UnimplementedGrpcUserFindControllerServer
	usecase *user.UserFindUsecase
}

func NewFind(usecase *user.UserFindUsecase) *GrpcUserFindController {
	return &GrpcUserFindController{usecase: usecase}
}

func (c *GrpcUserFindController) Execute(ctx context.Context, req *pb.UserId) (*pb.Response, error) {
	id := req.GetValue()
	if len(id) <= 0 {
		return nil, status.Error(codes.InvalidArgument, "id was not found")
	}
	dto, err := c.usecase.Execute(id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "unexpected error occurred")
	}
	res := &pb.Response{
		Id:        dto.Id,
		Firstname: dto.FirstName,
		Lastname:  dto.LastName,
	}
	return res, nil
}
