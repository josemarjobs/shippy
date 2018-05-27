package main

import (
	"log"

	pb "github.com/josemarjobs/shipper/user-service/proto/user"
	micro "github.com/micro/go-micro"
)

func main() {
	db, err := createConnection()
	if err != nil {
		log.Fatalf("could not connect to DB: %v\n", err)
	}
	defer db.Close()

	db.AutoMigrate(&pb.User{})

	repo := &UserRepository{db}
	tokenService := &TokenService{repo}

	srv := micro.NewService(
		micro.Name("go.micro.srv.user"),
		micro.Version("latest"),
	)
	srv.Init()

	pb.RegisterUserServiceHandler(srv.Server(), &service{repo, tokenService})

	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
