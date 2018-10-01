package main

import (
	"net"
	"log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"golang.org/x/net/context"

	 pb "github.com/emeve89/microgo/consignment-service/proto/consignment"
)

const (
	port = ":50051"
)

type IRepository interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
}

type Repository struct {
	consignments []*pb.Consignment
}

func (repo *Repository) Create(consignent *pb.Consignment) (*pb.Consignment, error) {
	updated := append(repo.consignments, consignent)
	repo.consignments = updated
	return consignent, nil
}

type service struct {
	repo IRepository
}

func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment) (*pb.Response, error) {
	consignment, err := s.repo.Create(req)
	if err != nil {
		return nil, err
	}

	return &pb.Response{Created: true, Consignment: consignment}, nil
}

func main() {
	repo := &Repository{}

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	pb.RegisterShippingServiceServer(s, &service{repo})

	reflection.Register(s)
	if err != s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}