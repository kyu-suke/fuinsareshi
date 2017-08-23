package daemon

import (
	"fmt"
	"time"
	"github.com/kyu-suke/fuinsareshi/client"
	"github.com/kyu-suke/fuinsareshi/status"
)

func HeartBeat(ips []string) {
	for i := 0; i < len(ips); i++ {

		fmt.Print(ips[i])

		time.Sleep(1 * time.Second)

		r, err := client.GetHeartBeat(ips[i] + ":8080")
		if err != nil {
			panic(err)
		}

		if r {
			status.Ips.Add(ips[i])

		} else {
			status.Ips.Del(ips[i])
		}

		if i == len(ips) - 1 {
			i = 0
		}
	}
}