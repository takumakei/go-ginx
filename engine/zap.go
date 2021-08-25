package engine

import (
	"io"

	"github.com/gin-gonic/gin"
	"github.com/openshift/installer/pkg/lineprinter"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	DefaultWriter io.Writer
	DefaultLevel  zapcore.Level = zapcore.DebugLevel
	DefaultMsg    string        = "GIN"
	DefaultKey    string        = "L"

	DefaultErrorWriter io.Writer
	DefaultErrorLevel  zapcore.Level = zapcore.ErrorLevel
	DefaultErrorMsg    string        = "GIN"
	DefaultErrorKey    string        = "L"
)

func init() {
	DefaultWriter = printer(func(line interface{}) {
		if ce := zap.L().Check(DefaultLevel, DefaultMsg); ce != nil {
			ce.Write(zap.Any(DefaultKey, line))
		}
	})

	DefaultErrorWriter = printer(func(line interface{}) {
		if ce := zap.L().Check(DefaultErrorLevel, DefaultErrorMsg); ce != nil {
			ce.Write(zap.Any(DefaultErrorKey, line))
		}
	})

	gin.DefaultWriter = DefaultWriter
	gin.DefaultErrorWriter = DefaultErrorWriter
}

func printer(p func(interface{})) io.Writer {
	return &lineprinter.LinePrinter{
		Print: (&lineprinter.Trimmer{
			WrappedPrint: func(args ...interface{}) {
				for _, a := range args {
					p(a)
				}
			},
		}).Print,
	}
}
