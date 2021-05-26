package main

import (
	"context"
	"errors"
	"flag"
	"log"
	"net"
	"os"

	pb "github.com/AlastorTh/My_gRPC/my_gRPC"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedDatabusServiceServer
	operation string
}

func (s *server) Send(ctx context.Context, in *pb.SendRequest) (*pb.SendResponse, error) {
	var res float32
	switch s.operation {
	case "add":
		res = in.GetPrm1() + in.GetPrm2()
		log.Printf("%f + %f = %f\n", in.GetPrm1(), in.GetPrm2(), res)
		return &pb.SendResponse{Result: res}, nil
	case "sub":
		res = in.GetPrm1() - in.GetPrm2()
		log.Printf("%f - %f = %f\n", in.GetPrm1(), in.GetPrm2(), res)
		return &pb.SendResponse{Result: res}, nil
	case "mul":
		res = in.GetPrm1() * in.GetPrm2()
		log.Printf("%f * %f = %f\n", in.GetPrm1(), in.GetPrm2(), res)
		return &pb.SendResponse{Result: res}, nil
	case "div":
		if in.GetPrm2() == 0 {
			return nil, errors.New("Division by zero not allowed!")
		}
		res = in.GetPrm1() / in.GetPrm2()
		log.Printf("%f / %f = %f\n", in.GetPrm1(), in.GetPrm2(), res)
		return &pb.SendResponse{Result: in.GetPrm1() / in.GetPrm2()}, nil
	default:
		log.Fatal("Not a valid operation")
		return nil, errors.New("Type Error here")
	}

}

func main() {
	flag.Parse()
	if flag.NArg() < 2 { //checking the num of args
		log.Fatalln("Invalid num of args: <port> <arithm. operation>")
	}

	port := os.Args[1]
	operation := os.Args[2]

	lis, err := net.Listen("tcp", ":"+port) // listen for server
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	srv := &server{operation: operation} // create + initialize struct with operation cmd argument
	pb.RegisterDatabusServiceServer(grpcServer, srv)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
