build:
	protoc -I . ./proto/user/user.proto --go_out=plugins=micro:./
	GOOS=linux GOARCH=amd64 go build -o user-service *.go
	docker build -t josemarjobs/user-service .

image:
	docker build -t josemarjobs/user-service .

run:
	docker container run --rm -p 50053:50051 --net shippy \
	  -e MICRO_SERVER_ADDRESS=:50051 \
	  -e MICRO_REGISTRY=mdns josemarjobs/user-service
	  