package router

import (
	"github.com/labstack/echo"
	"github.com/kyu-suke/fuinsareshi/api/v1/info"
	"github.com/kyu-suke/fuinsareshi/api/v1/heartbeat"
)

func SetupV1(e *echo.Echo) {

	g := e.Group("/v1")
	g.GET("/info", info.GetInfo)
	g.GET("/heartbeat", heartbeat.GetHeartBeat)
}
