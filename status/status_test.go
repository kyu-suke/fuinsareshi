package status

import (
	"testing"
)

func TestAliveIps_Add(t *testing.T) {
	ips := &AliveIps{Ip:[]string{"192.168.1.0", "192.168.1.1"}}

	addIp := "192.168.1.2"
	ips.Add(addIp)

	if isExist(*ips, addIp) == false {
		t.Fatal("trueになるはずなのにfalseになっているよ")
	}
}

func TestAliveIps_Del(t *testing.T) {
	ips := &AliveIps{Ip:[]string{"192.168.1.0", "192.168.1.1"}}

	delIp := "192.168.1.0"
	ips.Del(delIp)

	if isExist(*ips, delIp) == true {
		t.Fatal("falseになるはずなのにtrueになっているよ")
	}
}

func isExist(ips AliveIps, ip string) bool {
	for _, v := range ips.Ip {
		if v == ip {
			return true
		}

	}
	return false
}
