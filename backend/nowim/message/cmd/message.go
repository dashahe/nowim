package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"nowim.message/internal/api"
	"nowim.message/internal/config"
	_ "nowim.message/internal/config"
	_ "nowim.message/internal/db"
	"nowim.message/pkg/message"
	_ "nowim.message/pkg/ulid"
)

func main() {
	host, port := config.Config().GRpc.Host, config.Config().GRpc.Port
	lis, err := net.Listen("tcp", host + ":" + port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	message.RegisterMessageServer(grpcServer, api.NewMessageServer())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("grpc start failed, err: %+v", err)
	}
}
