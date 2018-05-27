build:
	GOOS=linux GOARCH=amd64 go build
	docker build -t josemarjobs/consignment-cli .

run:
	docker container run --net shippy --rm \
	  -e MICRO_REGISTRY=mdns \
	  josemarjobs/consignment-cli