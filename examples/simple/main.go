package main

import (
	"flag"

	"github.com/gin-gonic/gin"
	"github.com/takumakei/go-enzap"
	"github.com/takumakei/go-exit"
	"github.com/takumakei/go-ginx/engine"
	"github.com/takumakei/go-ginx/examples/simple/root"
)

var (
	FlagAddr string = ":8800"
)

func main() {
	flag.StringVar(&FlagAddr, "a", FlagAddr, "address to listen")
	flag.Parse()
	enzap.ReplaceGlobals()
	engine.SetDefaultMode(gin.ReleaseMode)
	r := engine.New()
	root.R.Mount(r)
	exit.Exit(r.Run(FlagAddr))
}
