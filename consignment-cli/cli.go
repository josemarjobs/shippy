package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"golang.org/x/net/context"

	pb "github.com/josemarjobs/shipper/consignment-service/proto/consignment"
	microclient "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
)

func parseFile(file string) (*pb.Consignment, error) {
	var consignment *pb.Consignment
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &consignment)
	return consignment, err
}

func main() {
	cmd.Init()

	client := pb.NewShippingServiceClient("go.micro.srv.consignment", microclient.DefaultClient)

	consignment, err := parseFile("consignment.json")
	fatalIfError(err, "could not parse file: %v\n", err)

	r, err := client.CreateConsignment(context.Background(), consignment)
	fatalIfError(err, "could not create consignment: %v\n", err)

	fmt.Printf("Created: %v\nConsignment: %+v\n", r.Created, r.GetConsignment())

	resp, err := client.GetConsignments(context.Background(), &pb.GetRequest{})
	fatalIfError(err, "could not list consignments: %v\n", err)
	fmt.Println("\n\nAll Consignments\n")
	for _, v := range resp.Consignments {
		fmt.Printf("%+v\n", v)
	}
}

func fatalIfError(err error, format string, v ...interface{}) {
	if err != nil {
		log.Fatalf(format, v)
	}
}
