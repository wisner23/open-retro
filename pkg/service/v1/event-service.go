package v1

import (
	"log"
	"context"
	"github.com/wisner23/open-retro/pkg/api/v1"
)

const (
	apiVersion = "v1"
)

type eventServiceServer struct { }

func NewEventServiceServer() v1.EventServiceServer {
	return &eventServiceServer{}
}

func (s *eventServiceServer) Create(ctx context.Context, request *v1.CreateRequest) (*v1.CreateResponse, error) {
	log.Printf("Http Server: Create request -  ID:%v\n", request.GetEvent().GetId())
	return &v1.CreateResponse { 
		Api: apiVersion,
		Id: request.GetEvent().GetId(),
	}, nil
}