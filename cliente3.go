package main

import (
	"fmt"
	"net"
	"encoding/gob"
)

type Persona struct {
	Nombre string
	Email []string
}

func cliente(persona Persona) {
	c, err := net.Dial("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = gob.NewEncoder(c).Encode(persona)
	if err != nil {
		fmt.Println(err)
	}
	c.Close()
}

func main()  {
	persona := Persona {
		Nombre: "Michel Dávalos",
		Email: []string{
			"michel.davalos@academicos.udg.mx",
			"michel.davalos@red.cucei.udg.mx",
			"1234456677",
		},
	}
	go cliente(persona)

	var input string
	fmt.Scanln(&input)
}