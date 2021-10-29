# http-echo

HTTP Echo is a go web server that serves the contents it was started with as an HTML page.

The default port is 8080, but this is configurable via the `--addr` flag:

```shell
# go build
make go-build
# start the server
./http-echo --text="hello world" --addr=:8080
```

Then visit http://localhost:8080/ in your browser.
