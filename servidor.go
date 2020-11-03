package main

import (
	"fmt"
	"net"
)

func servidor()  {
	s, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		c, err := s.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleClient(c)
	}
}

func handleClient(c net.Conn)  {
	b := make([]byte, 100)
	bs, err := c.Read(b)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("Mensaje: ", string(b[:bs]))
		fmt.Println("Bytes: ", bs)
	}
}

func main()  {
	go servidor()

	var input string
	fmt.Scanln(&input)
}