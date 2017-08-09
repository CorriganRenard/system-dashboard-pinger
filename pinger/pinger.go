package pinger

import (
	"fmt"
	"net"
	"os"
	"time"

	fastping "github.com/tatsushid/go-fastping"
)

func Ping(ips []string) map[string]bool {

	//type struct of type []ipList{} with fields of ip string and pingSuccess bool
	// type ipAddress struct {
	// 	ip          string
	// 	pingSuccess bool
	// }
	var ipmap map[string]bool
	ipmap = make(map[string]bool)

	p := fastping.NewPinger()

	//Get all the args and put them into a slice

	//range loop through the slice and resolve each ip address
	for _, v := range ips {

		ra, err := net.ResolveIPAddr("ip4:icmp", v)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		ipmap[ra.String()] = false
		p.AddIPAddr(ra)
	}

	p.OnRecv = func(addr *net.IPAddr, rtt time.Duration) {
		fmt.Printf("IP Addr: %s receive, RTT: %v\n", addr.String(), rtt)
		fmt.Println("success!")
		ipmap[addr.String()] = true
	}
	p.OnIdle = func() {
		fmt.Println("finish")
	}
	err := p.Run()
	if err != nil {
		fmt.Println(err)
	}

	return ipmap

}
