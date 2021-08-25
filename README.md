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
26T08:15:32.119967+0900 DEBUG   engine/zap.go:26        GIN     {"L": "[GIN-debug] [WARNING] Running in \"debug\" mode. Switch to \"release\" mode in production."}
26T08:15:32.120564+0900 DEBUG   engine/zap.go:26        GIN     {"L": " - using env:\texport GIN_MODE=release"}
26T08:15:32.120572+0900 DEBUG   engine/zap.go:26        GIN     {"L": " - using code:\tgin.SetMode(gin.ReleaseMode)"}
26T08:15:32.120576+0900 DEBUG   engine/zap.go:26        GIN     {"L": ""}
26T08:15:32.120711+0900 DEBUG   engine/zap.go:26        GIN     {"L": "[GIN-debug] GET    /hello/:name              --> github.com/takumakei/go-ginx/examples/simple/root.hello (4 handlers)"}
26T08:15:32.120764+0900 DEBUG   engine/zap.go:26        GIN     {"L": "[GIN-debug] GET    /                         --> github.com/takumakei/go-ginx/examples/simple/root.index (4 handlers)"}
26T08:15:32.120772+0900 DEBUG   engine/zap.go:26        GIN     {"L": "[GIN-debug] Listening and serving HTTP on :8800"}
26T08:15:36.539925+0900 INFO    zap@v0.0.1/zap.go:46    /       {"status": 200, "method": "GET", "path": "/", "query": "", "ip": "::1", "user-agent": "HTTPie/2.4.0", "time": "2021-08-26T08:15:36.539917+09:00", "latency": "726.851µs"}
26T08:15:36.540045+0900 DEBUG   engine/zap.go:26        GIN     {"L": "[GIN] 2021/08/26 - 08:15:36 | 200 |     783.272µs |             ::1 | GET      \"/\""}
^Csignal: interrupt
bash-5.1$
```

See also
----------------------------------------------------------------------

https://github.com/takumakei/go-enzap
> Getting zap.Logger out-of-the-box
