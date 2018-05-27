package main

import (
	"log"

	"golang.org/x/net/context"

	pb "github.com/josemarjobs/shipper/consignment-service/proto/consignment"
	vesselProto "github.com/josemarjobs/shipper/vessel-service/proto/vessel"
	micro "github.com/micro/go-micro"
)

const (
	port = ":50051"
)

type Repository interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
	GetAll() []*pb.Consignment
}
type ConsignmentRepo struct {
	consignments []*pb.Consignment
}

func (r *ConsignmentRepo) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	r.consignments = append(r.consignments, consignment)
	return consignment, nil
}

func (r *ConsignmentRepo) GetAll() []*pb.Consignment {
	return r.consignments
}

// grpc shipping service
type service struct {
	repo         Repository
	vesselClient vesselProto.VesselServiceClient
}

func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {
	vesselResponse, err := s.vesselClient.FindAvailable(context.Background(), &vesselProto.Specification{
		MaxWeight: req.Weight,
		Capacity:  int32(len(req.Containers)),
	})
	if err != nil {
		return err
	}
	log.Printf("Found vessel: %v\n", vesselResponse.Vessel.Name)

	// set the vessel id as the vessel we got from vessel service
	req.VesselId = vesselResponse.Vessel.Id

	c, err := s.repo.Create(req)
	if err != nil {
		return err
	}
	res.Created = true
	res.Consignment = c
	return nil
}

func (s *service) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	res.Consignments = s.repo.GetAll()
	return nil
}
func main() {
	repo := new(ConsignmentRepo)

	srv := micro.NewService(
		// name must match the package name given in the protobuf definition
		micro.Name("go.micro.srv.consignment"),
		micro.Version("latest"),
	)

	vesselClient := vesselProto.NewVesselServiceClient("go.micro.srv.vessel", srv.Client())

	// init will parse the command line flags
	srv.Init()

	pb.RegisterShippingServiceHandler(srv.Server(), &service{repo, vesselClient})

	if err := srv.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
