go-ginx
======================================================================

- Package github.com/takumakei/go-ginx/engine
  > Package engine implements the factory function of *gin.Engine writing logs using zap.
- Package github.com/takumakei/go-ginx/routes
  > Package routes provides a way to place route registrations near their implementation.

### example main

```go
func main() {
	enzap.ReplaceGlobals()
	r := engine.New()
	R.Mount(r)
	r.Run(":8080)
}
```

### example handler

```go
var R = routes.New()

func init() {
	R.GET("/hello/:name", hello)
}

func hello(c *gin.Context) {
	v := &struct {
		Name string
	}{
		Name: c.Param("name"),
	}
	c.PureJSON(http.StatusOK, v)
}
```

### ./examples/simple

```text
bash-5.1$ ZAP_DEVELOPMENT=1 GIN_MODE=debug go run ./examples/simple
19T09:12:06.633575+0900 DEBUG   engine/zap.go:26        GIN     {"line": "[GIN-debug] [WARNING] Running in \"debug\" mode. Switch to \"release\" mode in production."}
19T09:12:06.634126+0900 DEBUG   engine/zap.go:26        GIN     {"line": " - using env:\texport GIN_MODE=release"}
19T09:12:06.634133+0900 DEBUG   engine/zap.go:26        GIN     {"line": " - using code:\tgin.SetMode(gin.ReleaseMode)"}
19T09:12:06.634138+0900 DEBUG   engine/zap.go:26        GIN     {"line": ""}
19T09:12:06.634283+0900 DEBUG   engine/zap.go:26        GIN     {"line": "[GIN-debug] GET    /hello/:name              --> github.com/takumakei/go-ginx/examples/simple/root.hello (4 handlers)"}
19T09:12:06.634349+0900 DEBUG   engine/zap.go:26        GIN     {"line": "[GIN-debug] GET    /                         --> github.com/takumakei/go-ginx/examples/simple/root.index (4 handlers)"}
19T09:12:06.634357+0900 DEBUG   engine/zap.go:26        GIN     {"line": "[GIN-debug] Listening and serving HTTP on :8800"}
19T09:12:11.384544+0900 INFO    zap@v0.0.1/zap.go:46    /       {"status": 200, "method": "GET", "path": "/", "query": "", "ip": "::1", "user-agent": "HTTPie/2.4.0", "time": "2021-08-19T09:12:11.384535+09:00", "latency": "650.238µs"}
19T09:12:11.384660+0900 DEBUG   engine/zap.go:26        GIN     {"line": "[GIN] 2021/08/19 - 09:12:11 | 200 |     708.258µs |             ::1 | GET      \"/\""}
^Csignal: interrupt
bash-5.1$
```

See also
----------------------------------------------------------------------

https://github.com/takumakei/go-enzap
> Getting zap.Logger out-of-the-box
