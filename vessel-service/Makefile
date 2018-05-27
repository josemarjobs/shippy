build:
	protoc -I . ./proto/vessel/vessel.proto --go_out=plugins=micro:./

build-image:
	docker build -t josemarjobs/vessel-service .

run:
	docker run -p 50052:50051 --rm --net shippy \
	  -e MICRO_SERVER_ADDRESS=:50051 \
	  -e MICRO_REGISTRY=mdns josemarjobs/vessel-service