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

	//addr := net.UDPAddr{
	//	Port: 8080,
	//	IP:   net.ParseIP("LocalHost"),
	//}

	//conn, _ := net.ListenUDP("udp", &addr)

	//buffer := make([]byte, 1024)

	//for {
	//	n, addr, err := conn.ReadFromUDP(buffer)
	//	if err != nil {
	//		log.Print(err)
	//		continue
	//	}

	// Pedir autorización al pirata
	//	fmt.Println("Hola, pirata", n, addr)

	//	}
	//}
	conexion, err := net.ListenPacket("udp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conexion.Close() // close connection when finished

	fmt.Println("Servidor escuchando en localhost") //para probar que este funcionando la conexión UDP

	// Buffer para almacenar los datos recibidos
	buffer := make([]byte, 1024)
	// Esperar a recibir datos
	for {
		// Leer datos del cliente
		n, addr, err := conexion.ReadFrom(buffer)
		if err != nil {
			fmt.Println("Error al leer datos:", err)
			continue
		}

		fmt.Printf("Recibido mensaje de %s: %s\n", addr.String(), string(buffer[:n]))
		// Responder al cliente
		mensaje := []byte("Mensaje recibido correctamente") // Cmbiar mensaje y entregar los datos para la conexion TCP
		_, err = conexion.WriteTo(mensaje, addr)
		if err != nil {
			fmt.Println("Error al responder al cliente:", err)
			continue
		}
	}
}
