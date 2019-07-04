package main

import (
	"fmt"
	"os"
	"github.com/wisner23/open-retro/pkg/protocol/grpc"
)

func main() {
	if err := grpc.RunServer(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
