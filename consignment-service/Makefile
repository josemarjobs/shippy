build:
	protoc -I . ./proto/consignment/consignment.proto --go_out=plugins=grpc:./

build-image:
	GOOS=linux GOARCH=amd64 go build
	docker build -t josemarjobs/consignment-service .
run:
	docker container run -p 50051:50051 josemarjobs/consignment-service