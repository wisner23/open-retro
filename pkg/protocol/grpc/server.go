package grpc

import (
	"net"
	"google.golang.org/grpc"
	"github.com/wisner23/open-retro/pkg/api/v1"
	service "github.com/wisner23/open-retro/pkg/service/v1" 
)

func RunServer() error {
	server := grpc.NewServer()
	v1.RegisterEventServiceServer(server, service.NewEventServiceServer())
	
	listen, err := net.Listen("tcp", ":8093")

	if err != nil {
		return err
	}

	return server.Serve(listen)
}