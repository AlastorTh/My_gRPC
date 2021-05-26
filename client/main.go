// Package main implements a client for Greeter service.
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
	if flag.NArg() < 3 {
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
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewDatabusServiceClient(conn)
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Send(ctx, &pb.SendRequest{Prm1: float32(prm1), Prm2: float32(prm2)})
	if err != nil {
		log.Fatalf("could not Send: %v", err)
	}
	log.Println(r.GetResult())

}
