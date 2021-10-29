# http-echo

HTTP Echo is a go web server that echos back the arguments given to it. Inspired by [hashicorp/http-echo](https://github.com/hashicorp/http-echo).

Building Multi-Architecture Docker Images With [Buildx](https://docs.docker.com/buildx/working-with-buildx/). Apple chip's users cheers :).

### Local
The default port is 8080, but this is configurable via the `--addr` flag:

```shell
# go build
make go-build
# start the server
./http-echo --text="hello world" --addr=:8080
```

Then visit http://localhost:8080/ in your browser.

### Docker

To run the server in docker:

```shell
# docker pull
docker pull jxlwqq/http-echo:latest
# docker run
docker run -p 8080:8080 jxlwqq/http-echo:latest --text="hello world"
```

Then visit http://localhost:8080/ in your browser.

### Kubernetes

To run the server in kubernetes:

```yaml
kubectl apply -f http-echo.yaml
kubectl port-forward services/http-echo 8080:8080
```

Then visit http://localhost:8080/ in your browser.