package main

import (
	"github.com/labstack/echo"
	"github.com/kyu-suke/fuinsareshi/router"
	"flag"
	"github.com/kyu-suke/fuinsareshi/setting"
	"github.com/kyu-suke/fuinsareshi/util/utilip"
	"github.com/kyu-suke/fuinsareshi/daemon"
)

func main() {
	flag.StringVar(&setting.GroupName, "group", "default", "group name")
	flag.StringVar(&setting.Subnet, "subnet", "default", "subnet hani")
	flag.Parse()
	e := echo.New()
	router.SetupV1(e)

	// setting.Ips にハートビート
	ips, _ := utilip.Hosts(setting.Subnet)
	go daemon.HeartBeat(ips)


	e.Logger.Fatal(e.Start(":8080"))
}

