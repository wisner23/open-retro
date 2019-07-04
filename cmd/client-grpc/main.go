package main

import (
	"log"
	"time"
	"context"
	"google.golang.org/grpc"
	"github.com/wisner23/open-retro/pkg/api/v1"
)

const (
	apiVersion = "v1"
)

func main() {

	address := "localhost:8093"
	conn, err := grpc.Dial(address, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Did not connect: %v",err)
	}

	client := v1.NewEventServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	event, err := client.Create(ctx, &v1.CreateRequest{Api: "v1", Event: &v1.Event{Id: 37}})
	log.Printf("Added - id: %v", event.GetId())
}