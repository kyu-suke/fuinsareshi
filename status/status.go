package status

//import "fmt"

type AliveIps struct {
	Ip []string
}

var Ips AliveIps

func (a *AliveIps) Add (ip string) {
	for _, v := range a.Ip {
		if v == ip {
			return
		}
	}
	a.Ip = append(a.Ip, ip)
}

func (a *AliveIps) Del (ip string) {
	for k, v := range a.Ip {
		if v == ip {
			a.Ip = del(a.Ip, k)
		}
	}
}

func del (s []string, i int) []string {
	//新しいスライスを用意することがポイント
	s = append(s[:i], s[i+1:]...)
	n := make([]string, len(s))
	copy(n, s)
	return n
}

