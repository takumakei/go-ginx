package root

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func init() {
	R.GET("/", index)
}

func index(c *gin.Context) {
	v := &struct {
		Now string
	}{
		Now: time.Now().Format(time.RFC3339Nano),
	}
	c.PureJSON(http.StatusOK, v)
}
