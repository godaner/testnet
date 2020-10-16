package main

import (
	"flag"
	"fmt"
	"net"
	"time"
)

func main() {
	var addr string
	var nett string
	var udpto int64
	flag.StringVar(&addr, "addr", "addr", "addr")
	flag.StringVar(&nett, "net", "tcp", "net")
	flag.Int64Var(&udpto, "udpto", int64(time.Duration(2)*time.Second), "udp read timeout")
	flag.Parse()
	if addr == "" || nett == "" {
		flag.PrintDefaults()
		return
	}
	c, err := net.Dial(nett, addr)
	if err != nil {
		panic(err)
	}
	switch nett {
	case "udp":
		readErr := make(chan error, 0)
		go func() {
			bs := make([]byte, 1024, 1024)
			c.SetReadDeadline(time.Now().Add(time.Duration(udpto)))
			_, err := c.Read(bs)
			if err != nil {
				readErr <- err
				return
			}
			fmt.Println(string(bs))
			readErr <- nil
		}()
		_, err := c.Write([]byte("Hello , i am udp client."))
		if err != nil {
			panic(err)
		}
		err = <-readErr
		if err != nil {
			panic(err)
		}

	}

	fmt.Println(addr + " reachable!")
}
