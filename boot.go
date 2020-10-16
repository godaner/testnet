package main

import (
	"flag"
	"fmt"
	"net"
)

func main() {
	var addr string
	var nett string
	flag.StringVar(&addr, "addr", "addr", "addr")
	flag.StringVar(&nett, "net", "tcp", "net")
	flag.Parse()
	if addr == "" || nett == "" {
		flag.PrintDefaults()
		return
	}
	_, err := net.Dial(nett, addr)
	if err != nil {
		panic(err)
	}
	fmt.Println(addr + " reachable!")
}
