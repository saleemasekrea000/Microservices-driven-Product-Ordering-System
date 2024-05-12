package grpc

import (
	"fmt"
	"net"
	"product/internal/ports"
	"product/proto_gen"

	"google.golang.org/grpc"
)

type GrpcAdapter struct {
	ApiPort ports.ApiPort
	port    int
	proto_gen.UnimplementedProductServiceServer
}

func New(port int, apiPort ports.ApiPort) *GrpcAdapter {
	return &GrpcAdapter{
		ApiPort: apiPort,
		port:    port,
	}
}

func (s *GrpcAdapter) Run() error {
	port := fmt.Sprintf(":%d", s.port)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}
	grpcServer := grpc.NewServer()
	proto_gen.RegisterProductServiceServer(grpcServer, s)
	err = grpcServer.Serve(lis)
	if err != nil {
		return err
	}
	return nil
}
