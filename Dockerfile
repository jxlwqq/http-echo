# syntax=docker/dockerfile:1
FROM --platform=$TARGETPLATFORM golang:1.17-alpine AS builder
ARG TARGETARCH
ARG TARGETOS
WORKDIR /workspace
ADD go.mod go.mod
ADD go.sum go.sum
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -a -o http-echo main.go

FROM gcr.io/distroless/static:nonroot-$TARGETARCH
WORKDIR /
COPY --from=builder /workspace/http-echo .
EXPOSE 8080
ENTRYPOINT ["/http-echo"]