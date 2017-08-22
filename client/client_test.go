package client

import (
	"testing"
	"fmt"
	"github.com/kyu-suke/fuinsareshi/setting"
	"github.com/jarcoal/httpmock"
)

func TestGetHeartBeat(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "http://localhost:8081/v1/heartbeat",
		httpmock.NewStringResponder(200, `{"group":"default"}`))

	setting.GroupName = "default"
	r, err := GetHeartBeat("localhost:8081")
	if err != nil {
		fmt.Println(err)
	}

	if r == false {
		t.Fatal("trueになるはずなのにfalseになっているよ！")
	}
	fmt.Println(r)
}
