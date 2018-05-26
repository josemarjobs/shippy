package main

import (
	"log"
	"net"

	"google.golang.org/grpc/reflection"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "github.com/josemarjobs/shipper/consignment-service/proto/consignment"
)

const (
	port = ":50051"
)

type IRepository interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
	GetAll() []*pb.Consignment
}
type Repository struct {
	consignments []*pb.Consignment
}

func (r *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	r.consignments = append(r.consignments, consignment)
	return consignment, nil
}

func (r *Repository) GetAll() []*pb.Consignment {
	return r.consignments
}

// grpc shipping service
type service struct {
	repo IRepository
}

func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment) (*pb.Response, error) {
	c, err := s.repo.Create(req)
	if err != nil {
		return nil, err
	}
	return &pb.Response{Created: true, Consignment: c}, nil
}

func (s *service) GetConsignments(ctx context.Context, req *pb.GetRequest) (*pb.Response, error) {
	return &pb.Response{Consignments: s.repo.GetAll()}, nil
}
func main() {
	repo := new(Repository)
	// setup gRPC server
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}

	s := grpc.NewServer()
	pb.RegisterShippingServiceServer(s, &service{repo})

	reflection.Register(s)
	log.Printf("server starting on port: %v\n", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
