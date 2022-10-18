package main

import (
	"fmt"
	pb "github.com/stonewell/emacsheadless/proto"
)

type server struct {
	pb.UnimplementedHeadlessServer
}


//test
func main() {
	fmt.Println("This is a golang project with cmake.")
}
