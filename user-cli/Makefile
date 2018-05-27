build:
	GOOS=linux GOARCH=amd64 go build -o user-cli *.go
	docker build -t josemarjobs/user-cli .
	
run:
	docker run -r MICRO_REGISTRY=mdns josemarjobs/user-cli

compose-run: 
	docker-compose run user-cli command \
	 --name="Peter Griffin" \
	 --email="peterg@gmail.com" \
	 --password="supersecrer" \
	 --company="Beer Bottle Inc"