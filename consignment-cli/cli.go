package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	pb "github.com/josemarjobs/shipper/consignment-service/proto/consignment"
	"google.golang.org/grpc"
)

const (
	address         = "localhost:50051"
	defaultFilename = "consignment.json"
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
	file := flag.String("file", defaultFilename, "default json file to post")
	host := flag.String("server", address, "default server url")
	flag.Parse()

	conn, err := grpc.Dial(*host, grpc.WithInsecure())
	fatalIfError(err, "failed to connect: %v\n", err)
	defer conn.Close()

	client := pb.NewShippingServiceClient(conn)

	consignment, err := parseFile(*file)
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
