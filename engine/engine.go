// Package engine implements the factory function of *gin.Engine writing logs using zap.
package engine

import (
	"os"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// New returns the *gin.Engine writing logs using zap.
func New() *gin.Engine {
	router := gin.New()

	router.Use(
		gin.LoggerWithWriter(DefaultWriter),
		ginzap.Ginzap(zap.L(), time.RFC3339Nano, false),
		ginzap.RecoveryWithZap(zap.L(), true),
	)

	return router
}

// SetDefaultMode changes the mode of gin to defaultMode if the environment
// variable GIN_MODE (a.k.a gin.EnvGinMode) is not set, otherwise set to
// GIN_MODE.
func SetDefaultMode(defaultMode string) {
	mode := os.Getenv(gin.EnvGinMode)
	if mode == "" {
		mode = defaultMode
	}
	gin.SetMode(mode)
}
