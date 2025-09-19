all: build run

build: 
	go build ./cmd/catalog-service-fiap

run:
	./catalog-service-fiap

test: 
	go test -covermode=atomic -coverprofile=coverage.out `go list ./... | grep -v mocks | grep -v cmd | grep -v testdata`

cov: test
	go tool cover -html=coverage.out

gen: 
	go generate ./...

swagger:
	swagger generate spec -o ./swagger.yaml --scan-models

docker-up:
	docker-compose up -d --build

docker-down:
	docker-compose down -v