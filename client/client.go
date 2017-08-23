package client

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/kyu-suke/fuinsareshi/setting"
	"fmt"
	"time"
)

type Client struct {
	Group string
}

func GetHeartBeat(host string) (bool, error) {

	// request get
	httpClient := http.Client{
		Timeout: time.Duration(time.Second * 1),
	}
	r, err := httpClient.Get(fmt.Sprintf("http://%s/v1/heartbeat", host))
	if err != nil {
		return false, err
	}

	defer r.Body.Close()

	// response kakunin
	ba, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return false, err
	}

	// json decode
	heartBeatClient := new(Client)
	if err := json.Unmarshal(ba, heartBeatClient); err != nil {
		return false, err
	}

	// status code , group name check
	if r.StatusCode == 200 && heartBeatClient.Group == setting.GroupName {
		return true, nil
	}

	return false, nil
}
