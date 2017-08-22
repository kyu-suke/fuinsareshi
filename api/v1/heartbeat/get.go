package heartbeat

import (
	"github.com/labstack/echo"
	"net/http"
	"github.com/kyu-suke/fuinsareshi/setting"
)

type GetResponse struct {
	Group string `json:"group"`
}

func GetHeartBeat(c echo.Context) error {
	// 200とグループ情報を返す
	r := &GetResponse{
		Group: setting.GroupName,
	}
	return c.JSON(http.StatusOK, r)
}