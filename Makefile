.PHONY: go-build
go-build:
	go build -o http-echo main.go

.PHONY: docker-build
docker-build:
	docker build -t jxlwqq/http-echo .

.PHONY: go-test
go-test:
	go test -v ./...