protobuf:
	protoc -I . ./proto/consignment/consignment.proto --go_out=plugins=micro:./

micro:
	protoc -I . ./proto/consignment/consignment.proto --micro_out=. --go_out=.

build:
	GOOS=linux GOARCH=amd64 go build -o consignment-service *.go

image:
	docker build -t josemarjobs/consignment-service .

run:
	docker container run --rm -p 50051:50051 --net shippy \
	  -e MICRO_SERVER_ADDRESS=:50051 \
	  -e MICRO_REGISTRY=mdns josemarjobs/consignment-service

run-consul:
	docker run -d --rm --net shippy --name consul -e CONSUL_BIND_INTERFACE=eth0 consul