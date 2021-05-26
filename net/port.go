package main

import (
	"fmt"
	"net"
)

func main() {
	for i := 0; i < 65535; i++ {
		// 1, 2, ..4, 99
		// sitio:1, sitio:2, sitio:99, ...,
		// 1 -> Open,  2 -> Closed, ...
		conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", "scanme.nmap.org", i))
		if err != nil {
			continue
		}
		conn.Close()
		fmt.Printf("Port %d is open\n", i)
	}
}
