package service

import (
	"github.com/jarcoal/httpmock"
	"testing"
)

func TestProxy(t *testing.T) {
	Proxy()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "http://localhost:8081/v1/heartbeat",
		httpmock.NewStringResponder(200, `{"group":"default"}`))
	// どっか適当なページにリクエスト送る httpmock
	// prosy経由で:9090にリクエスト送る
	// レスポンス確認する
}
