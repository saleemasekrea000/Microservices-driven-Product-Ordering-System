package grpc

import (
	"fmt"
	"net"
	"order/internal/ports"
	"order/proto_gen"

	GrpcServer "google.golang.org/grpc"
)

type GrpcAdapter struct {
	apiPort ports.APIPort
	port    int
	proto_gen.UnimplementedOrderServiceServer
}

func NewAdatpter(port int, apiPort ports.APIPort) *GrpcAdapter {
	return &GrpcAdapter{
		port:    port,
		apiPort: apiPort,
	}
}

func (s *GrpcAdapter) Run() error {
	port := fmt.Sprintf(":%d", s.port)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}
	grpcServer := GrpcServer.NewServer()
	proto_gen.RegisterOrderServiceServer(grpcServer, s)
	err = grpcServer.Serve(lis)
	if err != nil {
		return err
	}
	return nil
}
