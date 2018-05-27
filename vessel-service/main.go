package main

import (
	"fmt"
	"log"
	"os"

	pb "github.com/josemarjobs/shipper/vessel-service/proto/vessel"
	micro "github.com/micro/go-micro"
)

const (
	defaultDbHost = "localhost:27017"
)

func createDummyData(repo Repository) {
	defer repo.Close()

	vessels := []*pb.Vessel{
		&pb.Vessel{Id: "vessel01", Name: "Titanic", MaxWeight: 2000000, Capacity: 500},
	}

	fmt.Println("Creating dummy data")
	for _, v := range vessels {
		repo.Create(v)
		fmt.Printf("created - %v\n", v)
	}
}

func main() {

	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = defaultDbHost
	}
	session, err := CreateSession(dbHost)
	if err != nil {
		log.Fatalf("unable to connect to mongo: %v", err)
	}
	defer session.Close()

	createDummyData(&VesselRepo{session.Copy()})

	srv := micro.NewService(
		micro.Name("go.micro.srv.vessel"),
		micro.Version("latest"),
	)

	srv.Init()

	pb.RegisterVesselServiceHandler(srv.Server(), &service{session})

	if err := srv.Run(); err != nil {
		log.Println(err)
	}
}
