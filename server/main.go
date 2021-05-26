package main

import (
	"context"
	"errors"
	"flag"
	"log"
	"net"
	"os"

	My_gRPC "github.com/AlastorTh/My_gRPC/my_gRPC"
	pb "github.com/AlastorTh/My_gRPC/my_gRPC"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedDatabusServiceServer
	operation string
}

func (s *server) Send(ctx context.Context, in *pb.SendRequest) (*pb.SendResponse, error) {
	switch s.operation {
	case "add":
		log.Printf("%f + %f\n", in.GetPrm1(), in.GetPrm2())
		return &pb.SendResponse{Result: in.GetPrm1() + in.GetPrm2()}, nil
	case "sub":
		log.Printf("%f - %f\n", in.GetPrm1(), in.GetPrm2())
		return &pb.SendResponse{Result: in.GetPrm1() - in.GetPrm2()}, nil
	case "mul":
		log.Printf("%f * %f\n", in.GetPrm1(), in.GetPrm2())
		return &pb.SendResponse{Result: in.GetPrm1() * in.GetPrm2()}, nil
	case "div":
		log.Printf("%f / %f\n", in.GetPrm1(), in.GetPrm2())
		return &pb.SendResponse{Result: in.GetPrm1() / in.GetPrm2()}, nil
	default:
		log.Println("Not a valid operation")
		return nil, errors.New("Operation not valid; please use add, sub, mul, div")
	}

}

func main() {
	flag.Parse()
	if flag.NArg() < 2 {
		log.Fatalln("Invalid num of args: <port> <arithm. operation>")
	}

	port := os.Args[1]
	operation := os.Args[2]

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	srv := &server{operation: operation}
	My_gRPC.RegisterDatabusServiceServer(grpcServer, srv)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
