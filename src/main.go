package main

/*
   #include "headless.h"
*/
import "C"

import (
	"io"
	"net"
	"google.golang.org/grpc"
	pb "github.com/stonewell/emacsheadless/proto"
	log "github.com/sirupsen/logrus"
)

type server struct {
	pb.UnimplementedHeadlessServer
}

func main() {
	start_server(&C.ServerCallbacks{})
}

func (*server) Connect(stream pb.Headless_ConnectServer) error {
	log.Info("Connect Function")

	for {
		// Receive the request and possible error from the stream object
		req, err := stream.Recv()

		// If there are no more requests, we return
		if err == io.EOF {
			return nil
		}

		// Handle error from the stream object
		if err != nil {
			log.Println("Error when reading client request stream:", err)
			return err
		}

		log.Println("recv req:", req.Type)
	}
}

//export start_server
func start_server(callback * C.ServerCallbacks) {
	log.Info("Starting the server...")

	C.bridge_keyboard_input(callback.input, 42)

	lis, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalf("Unable to listen on port 3000: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterHeadlessServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
