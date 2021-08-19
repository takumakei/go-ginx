package root

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

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
