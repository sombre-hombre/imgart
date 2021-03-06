BINARY_NAME=imgartapi

default: test vet

.PHONY: packages
packages:
	go list ./...

.PHONY: test
test: 
	go test -v ./...

.PHONY: vet
vet:
	go vet ./...

.PHONY: clean
clean:
	go clean
	rm $(BINARY_NAME)

.PHONY: deps
deps:
	go get -u github.com/golang/dep/cmd/dep
	dep ensure

.PHONY: build
build: deps
	go build -o $(BINARY_NAME) -v ./cmd/imgartapi

.PHONY: docker
docker:
	docker-compose up

.PHONY: docker-debug
docker-debug:
	docker-compose -f docker-compose.yml -f docker-compose.debug.yml up