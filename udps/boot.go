package main

import (
	"flag"
	"fmt"
	"net"
	"strconv"
	"strings"
)

func main() {
	var addr string
	flag.StringVar(&addr, "addr", "addr", "addr")
	flag.Parse()
	if addr == "" {
		flag.PrintDefaults()
		return
	}
	sp := strings.Split(addr, ":")
	ip := sp[0]
	port, _ := strconv.ParseInt(sp[1], 10, 64)
	if ip == "" {
		ip = "0.0.0.0"
	}
	fmt.Println(ip + ":" + fmt.Sprint(port))
	c, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.ParseIP(ip),
		Port: int(port),
		Zone: "",
	})
	if err != nil {
		panic(err)
	}
	bs := make([]byte, 1024, 1024)
	for {
		n, ra, err := c.ReadFromUDP(bs)
		if err != nil {
			panic(err)
		}
		_, err = c.WriteToUDP([]byte("Hello , i am udp server.I got your message : "+string(bs[0:n])), ra)
		if err != nil {
			panic(err)
		}
	}

}
