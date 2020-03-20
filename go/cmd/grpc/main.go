package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"strconv"
	"time"

	controller "github.com/jupemara/ddd-guys/go/adapter/controller/grpc/user"
	pb_find "github.com/jupemara/ddd-guys/go/adapter/controller/grpc/user/find"
	pb_register "github.com/jupemara/ddd-guys/go/adapter/controller/grpc/user/register"
	pb_update "github.com/jupemara/ddd-guys/go/adapter/controller/grpc/user/update"
	repository "github.com/jupemara/ddd-guys/go/adapter/repository/user"
	usecase "github.com/jupemara/ddd-guys/go/usecase/user"
	"google.golang.org/grpc"
)

const (
	port = 10000
)

func grpcRegister(ctx context.Context, conn *grpc.ClientConn, firstname, lastname string) {
	c := pb_register.NewGrpcUserRegisterControllerClient(conn)
	_, err := c.Execute(ctx, &pb_register.User{
		Firstname: firstname,
		Lastname:  lastname,
	})
	if err != nil {
		log.Fatalf("could not register: %v", err)
	}
	log.Printf("Registered: %s %s", firstname, lastname)
}

func grpcFind(ctx context.Context, conn *grpc.ClientConn, id string) {
	c := pb_find.NewGrpcUserFindControllerClient(conn)
	user, err := c.Execute(ctx, &pb_find.UserId{
		Value: id,
	})
	if err != nil {
		log.Fatalf("could not find user: %v", err)
	}
	log.Printf("found an user: %v", user)
}

func grpcUpdate(ctx context.Context, conn *grpc.ClientConn, id, firstname, lastname string) {
	c := pb_update.NewGrpcUserUpdateControllerClient(conn)
	_, err := c.Execute(ctx, &pb_update.User{
		Id:        id,
		Firstname: firstname,
		Lastname:  lastname,
	})
	if err != nil {
		log.Fatalf("could not update user: %v", err)
	}
	log.Printf("updated an user: %v", id)
}

func main() {
	var (
		server    = flag.Bool("server", false, "run as a grpc server")
		generate  = flag.Bool("generate", false, "generate users")
		update    = flag.Bool("update", false, "update an user")
		find      = flag.Bool("find", false, "find an user")
		id        = flag.String("id", "", "id of the user to find/update")
		firstname = flag.String("firstname", "", "firstname for the user to be updated")
		lastname  = flag.String("lastname", "", "lastname for the user to be updated")
	)
	flag.Parse()

	if *server {
		lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		var opts []grpc.ServerOption
		grpcServer := grpc.NewServer(opts...)
		pb_find.RegisterGrpcUserFindControllerServer(
			grpcServer, controller.NewFind(usecase.NewUserFindUsecase(repository.NewCsvRepository())))
		pb_register.RegisterGrpcUserRegisterControllerServer(
			grpcServer, controller.NewRegister(usecase.NewUserRegisterUsecase(repository.NewUuidIdProvider(), repository.NewCsvRepository())))
		pb_update.RegisterGrpcUserUpdateControllerServer(
			grpcServer, controller.NewUpdate(usecase.NewUserUpdateUsecase(repository.NewCsvRepository())))
		err = grpcServer.Serve(lis)
		if err != nil {
			log.Fatalln(err)
		}
	} else {
		conn, err := grpc.Dial("localhost:"+strconv.Itoa(port), grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Fatalf("failed to connect grpc server: %v", err)
		}
		defer conn.Close()
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		if *generate {
			grpcRegister(ctx, conn, "John", "Doh")
			grpcRegister(ctx, conn, "Rob", "Pike")
			grpcRegister(ctx, conn, "Ken", "Thompson")
		} else if *update {
			grpcUpdate(ctx, conn, *id, *firstname, *lastname)
		} else if *find {
			grpcFind(ctx, conn, *id)
		}
	}
}
