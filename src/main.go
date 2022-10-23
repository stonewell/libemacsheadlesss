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

type Server struct {
	pb.UnimplementedHeadlessServer

	sc *C.ServerConfig
	clients map[int]Client
}

type Client struct {
	clientId int
	callback * C.ClientCallbacks
	stream pb.Headless_ConnectServer
}

func main() {
	StartServer(nil)
}

func (self *Server) Connect(stream pb.Headless_ConnectServer) error {
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

		switch req.Type {
		case pb.CmdType_Cmd_NewClient:
			self.OnNewClient(stream)
		case pb.CmdType_Cmd_ClientDisconnect:
		case pb.CmdType_Cmd_ClientInfo:
		case pb.CmdType_Cmd_Nope:
			log.Println("Nope command type:")
		default:
			log.Println("Invalid command type:", req.Type)
		}
	}
}

//export StartServer
func StartServer(sc *C.ServerConfig) {
	log.Info("Starting the server...")

	lis, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalf("Unable to listen on port 3000: %v", err)
	}

	s := grpc.NewServer()

	s_impl := Server {
		sc: sc,
		clients: make(map[int]Client),
	}

	pb.RegisterHeadlessServer(s, &s_impl)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

//export RegisterClientCallback
func RegisterClientCallback(clientId int, callback * C.ClientCallbacks) {
}

//export DisconnectClient
func DisconnectClient(clientId int) {
}

func (self * Server)OnNewClient(stream pb.Headless_ConnectServer) {
}
