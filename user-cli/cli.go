package main

import (
	"log"
	"os"

	pb "github.com/josemarjobs/shipper/user-service/proto/user"
	"github.com/micro/cli"
	micro "github.com/micro/go-micro"
	microClient "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
	"golang.org/x/net/context"
)

func main() {
	cmd.Init()

	client := pb.NewUserServiceClient("go.micro.srv.user", microClient.DefaultClient)

	service := micro.NewService(
		micro.Flags(
			cli.StringFlag{Name: "name", Usage: "Your full name"},
			cli.StringFlag{Name: "email", Usage: "Your email"},
			cli.StringFlag{Name: "password", Usage: "Your password"},
			cli.StringFlag{Name: "company", Usage: "Your company"},
		),
	)

	service.Init(
		micro.Action(func(c *cli.Context) {
			name := c.String("name")
			email := c.String("email")
			password := c.String("password")
			company := c.String("company")

			r, err := client.Create(context.TODO(), &pb.User{
				Name:     name,
				Email:    email,
				Password: password,
				Company:  company,
			})
			if err != nil {
				log.Fatalf("could not create user: %v\n", err)
			}
			log.Printf("user created: %s\n", r.User.Id)

			getAll, err := client.GetAll(context.Background(), &pb.Request{})
			if err != nil {
				log.Fatalf("could not list users: %v\n", err)
			}
			log.Printf("Got %v users\n", len(getAll.Users))
			for _, u := range getAll.Users {
				log.Printf("%+v\n", u)
			}
			os.Exit(0)
		}),
	)

	if err := service.Run(); err != nil {
		log.Printf("Error running service: %v\n", err)
	}
}
