package info

import (
	"github.com/labstack/echo"
	"fmt"
)

func GetInfo(c echo.Context) error {
	fmt.Println("hoge")
	var e error
	return e
}
