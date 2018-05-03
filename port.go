package main

import (
	"fmt"
	"net"
	"flag"
)

var ip *string
var port *int

func init() {

	ip = flag.String("ip", "127.0.0.1", "ip")
	port = flag.Int("port", 80, "port")
	flag.Parse()
}

func main() {
	ips := net.ParseIP(*ip)
	if CheckEnabled(&ips, *port) {
		fmt.Println(*port,"已经起用")
	} else {
		fmt.Println(*port,"没有起用")
	}
}
func CheckEnabled(ip *net.IP, p int) bool {
	/*默认本机ip*/
	if ip == nil {
		ip = &net.IP{127,0,0,1}
	}

	tcpAddr := net.TCPAddr{
		IP: *ip,
		Port: p,
	}

	tcpConn, err := net.DialTCP("tcp", nil, &tcpAddr)
	if err != nil {
		return false
	}
	defer tcpConn.Close()
	return true
}