package servers

import (
	pb "awesomeProject/protoFiles/files"
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	pb.UnimplementedApiServer
}

func New() *Server {
	return &Server{}
}

func StartServer() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterApiServer(grpcServer, New())
	log.Println("server is running on port 8080...")
	if err = grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (s *Server) GenerateInvoiceFPS(ctx context.Context, req *pb.GenerateInvoiceRequestFPSPostman) (*pb.GenerateInvoiceResponseFPSPostman, error) {
	return s.InvoiceFPS(ctx, req), nil
}
