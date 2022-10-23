package main

import (
	"context"
	"fmt"
	"time"

	pb "github.com/stonewell/emacsheadless/proto"
	log "github.com/sirupsen/logrus"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Client Stream...")
	fmt.Println()

	opts := grpc.WithInsecure()
	con, err := grpc.Dial("localhost:3000", opts)
	if err != nil {
		log.Fatalf("Error connecting: %v \n", err)
	}

	defer con.Close()
	c := pb.NewHeadlessClient(con)

	stream, err := c.Connect(context.Background())
	if err != nil {
		log.Fatalf("Error when getting stream object: %v", err)
		return
	}

	stream.Send(&pb.Cmd {
		Type: pb.CmdType_Cmd_NewClient,
	})

	time.Sleep(1 * time.Second)

	stream.CloseSend()
}
