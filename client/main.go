// Package main implements a client for DataBus service.
package main

import (
	"context"
	"flag"
	"log"
	"strconv"
	"time"

	pb "github.com/AlastorTh/My_gRPC/my_gRPC"

	"google.golang.org/grpc"
)

func main() {

	flag.Parse()
	if flag.NArg() < 3 { // check if got enough args
		log.Fatal("Invalid args number: <address> <num1> <num2>")
	}

	address := flag.Arg(0)

	prm1, err := strconv.Atoi(flag.Arg(1))
	if err != nil {
		log.Fatalln(err.Error())
	}

	prm2, err := strconv.Atoi(flag.Arg(2))
	if err != nil {
		log.Fatalln(err.Error())
	}

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock()) // establishing connection to the server
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	c := pb.NewDatabusServiceClient(conn) // client init
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()
	r, err := c.Send(ctx, &pb.SendRequest{Prm1: float32(prm1), Prm2: float32(prm2)}) // calling rpc method, passing prms
	if err != nil {
		log.Fatalf("could not Send: %v", err)
	}
	log.Println(r.GetResult()) // retrieving res from client

}
