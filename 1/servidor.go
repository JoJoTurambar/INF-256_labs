package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
)

func main() {

	// tablero de Blackbeard
	blackbeard := [4]string{"A", "B", "C", "D"}

	//tablero del legendario pirata
	pirata := [4]string{"A", "B", "C", "D"}

	// X representa la ubicación del barco
	x_b := rand.Intn(4)
	blackbeard[x_b] = "X"

	x_p := rand.Intn(4)
	pirata[x_p] = "X"

	// fmt.Println(blackbeard, pirata)

	addr := net.UDPAddr{
		Port: 8080,
		IP:   net.ParseIP("LocalHost"),
	}

	conn, _ := net.ListenUDP("udp", &addr)

	buffer := make([]byte, 1024)

	for {
		n, addr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			log.Print(err)
			continue
		}

		// Pedir autorización al pirata
		fmt.Println("Hola, pirata", n, addr)

	}
}
