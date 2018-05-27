package main

import (
	"log"
	"os"

	pb "github.com/josemarjobs/shipper/consignment-service/proto/consignment"
	vesselProto "github.com/josemarjobs/shipper/vessel-service/proto/vessel"
	micro "github.com/micro/go-micro"
)

const (
	defaultDbHost = "localhost:27017"
)

func main() {
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = defaultDbHost
	}

	session, err := CreateSession(dbHost)
	if err != nil {
		log.Panicf("could not connect to datastore with host %v: %v\n", dbHost, err)
	}
	defer session.Close()

	srv := micro.NewService(
		// name must match the package name given in the protobuf definition
		micro.Name("go.micro.srv.consignment"),
		micro.Version("latest"),
	)

	vesselClient := vesselProto.NewVesselServiceClient("go.micro.srv.vessel", srv.Client())

	// init will parse the command line flags
	srv.Init()

	pb.RegisterShippingServiceHandler(srv.Server(), &service{session, vesselClient})

	if err := srv.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
