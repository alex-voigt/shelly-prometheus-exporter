help:
	@echo "Available commands:"
	@echo "	run                - runs the exporter"
	@echo "	watch              - runs the exporter with hot reload"
	@echo "	test               - runs the tests"
	@echo "	docker-build       - builds the docker container"
	@echo "	docker-run         - runs the docker container"
	@echo ""

.PHONY: run
run:
	go run *.go

.PHONY: watch
watch:
	go get -u github.com/cosmtrek/air
	air

.PHONY: test
test:
	go test ./... -timeout 30s -v -cover

.PHONY: docker-build
docker-build:
	docker build -t shelly-exporter .

.PHONY: docker-run
docker-run:
	 docker run --name shelly-exporter -v "$$(pwd)"/config.yaml:/app/config.yaml --rm -it shelly-exporter:latest