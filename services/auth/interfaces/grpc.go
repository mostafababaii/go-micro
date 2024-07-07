package interfaces

import (
	"log"
	"net"

	auth "github.com/mostafababaii/go-micro/services/auth/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

func StartApplication() {
	grpclog.Println("Starting GRPC server")

	lis, err := net.Listen("tcp", ":8282")
	if err != nil {
		log.Fatal(err)
	}
	defer lis.Close()

	grpclog.Println("Listining on 127.0.0.1:8282")

	var ops []grpc.ServerOption
	grpcServer := grpc.NewServer(ops...)

	authServer := auth.NewGrpcServer()

	auth.RegisterAuthServer(grpcServer, authServer)

	if err = grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
