package service

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	pb "github.com/stonewell/emacsheadless/proto"
	"google.golang.org/grpc"
	"io"
	"net"
)

type ServerConfigInterface interface {
	OnNewClient(clientId uint)
	OnClientDisconnect(clientId uint)
	GetPort() uint
	GetAddr() string
}

type ServerConfig struct {
	Port uint
	Addr string
}

type ClientCallbacks interface {
	KeyboardInput(keycode uint)
}

type Server struct {
	pb.UnimplementedHeadlessServer

	ServiceConfig ServerConfigInterface
	Clients       map[int32]Client

	ClientChannel chan ClientOP
}

type Client struct {
	ClientId int32
	Callback ClientCallbacks
	Stream   pb.Headless_ConnectServer
}

type ClientOP struct {
	Op     pb.CmdType
	Client Client
}

var G_ServImpl Server

func (self *ServerConfig) GetAddr() string {
	return self.Addr
}

func (self *ServerConfig) GetPort() uint {
	return self.Port
}

func (self *ServerConfig) OnNewClient(clientId uint) {
	log.Debug("On NewClient:", clientId)
}

func (self *ServerConfig) OnClientDisconnect(clientId uint) {
}

func StartServer(sc ServerConfigInterface) {
	addr := fmt.Sprintf("%s:%d", sc.GetAddr(), sc.GetPort())

	log.Debug("Starting the server...", addr)

	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Unable to listen on ", addr, err)
	}

	s := grpc.NewServer()

	G_ServImpl = Server{
		ServiceConfig: sc,
		Clients:       make(map[int32]Client),
		ClientChannel: make(chan ClientOP),
	}

	go ClientLifeCycleOp(&G_ServImpl)

	pb.RegisterHeadlessServer(s, &G_ServImpl)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: ", err)
	}
}

func (self *Server) Connect(stream pb.Headless_ConnectServer) error {
	log.Debug("Connect Function")

	for {
		// Receive the request and possible error from the stream object
		req, err := stream.Recv()

		// If there are no more requests, we return
		if err == io.EOF {
			return nil
		}

		// Handle error from the stream object
		if err != nil {
			log.Error("Error when reading client request stream:", err)
			return err
		}

		log.Debug("recv req:", req.Type)

		switch req.Type {
		case pb.CmdType_Cmd_NewClient:
			self.OnNewClient(stream)
		case pb.CmdType_Cmd_ClientDisconnect:
			self.ServiceConfig.OnClientDisconnect(uint(req.ClientInfo.ClientId))
		case pb.CmdType_Cmd_ClientInfo:
		case pb.CmdType_Cmd_Nope:
			log.Println("Nope command type:")
		default:
			log.Println("Invalid command type:", req.Type)
		}
	}
}

func (self *Server) OnNewClient(stream pb.Headless_ConnectServer) {
	log.Debug("Client Life Cycle:On NewClient")

	self.ClientChannel <- ClientOP{
		Op: pb.CmdType_Cmd_NewClient,
		Client: Client{
			ClientId: -1,
			Stream:   stream,
		},
	}
}

func ClientLifeCycleOp(server *Server) {
	for clientOP := range server.ClientChannel {
		log.Debug("Client Life Cycle OP", clientOP.Op)

		switch clientOP.Op {
		case pb.CmdType_Cmd_NewClient:
			clientCount := int32(len(server.Clients))
			newClient := clientOP.Client
			newClient.ClientId = clientCount
			server.Clients[clientCount] = newClient

			server.ServiceConfig.OnNewClient(uint(clientCount))
		case pb.CmdType_Cmd_ClientDisconnect:
			if client, ok := server.Clients[clientOP.Client.ClientId]; ok {
				client.Stream.Send(&pb.Cmd{
					Type: pb.CmdType_Cmd_ClientDisconnect,
				})

				delete(server.Clients, client.ClientId)
			}
		case pb.CmdType_Cmd_ClientInfo:
			if client, ok := server.Clients[clientOP.Client.ClientId]; ok {
				client.Callback = clientOP.Client.Callback

				client.Stream.Send(&pb.Cmd{
					Type: pb.CmdType_Cmd_ClientInfo,
					ClientInfo: &pb.ClientInfo{
						ClientId: client.ClientId,
					},
				})
			}
		}
	}
}
