.PHONY: go-build
go-build:
	go build -o http-echo main.go

.PHONY: docker-buildx
docker-buildx:
	docker buildx build -t jxlwqq/http-echo --platform linux/amd64,linux/arm64 --push .
