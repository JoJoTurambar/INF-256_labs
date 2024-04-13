package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"strings"
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

	fmt.Println("Servidor escuchando en localhost")

	// Buffer para almacenar los datos recibidos
	buffer := make([]byte, 1024)
	// Esperar a recibir datos
	for {
		// Leer datos del cliente UDP
		n, clientAddr, err := conexion.ReadFrom(buffer)
		if err != nil {
			fmt.Println("Error al leer datos UDP:", err)
			continue
		}

		// Imprime el mensaje recibido
		fmt.Printf("Mensaje recibido desde %s: %s\n", clientAddr, string(buffer[:n]))

		// Obtener la dirección IP del cliente UDP
		clientIP := strings.Split(clientAddr.String(), ":")[0]

		// Prepara la dirección IP para conexión TCP
		tcpAddr := clientIP + ":8081"

		// Responder al cliente UDP con la dirección IP para conexión TCP
		mensaje := []byte(tcpAddr)
		_, err = conexion.WriteTo(mensaje, clientAddr)
		if err != nil {
			fmt.Println("Error al responder al cliente UDP:", err)
			continue
		}

		// Cerrar la conexión UDP
		conexion.Close()
		break // Salir del bucle después de cerrar la conexión
	}

	// Crear el listener TCP
	tcpListener, err := net.Listen("tcp", ":8081")
	if err != nil {
		fmt.Println("Error al crear el listener TCP:", err)
		return
	}
	defer tcpListener.Close()
	//fmt.Println("Pasó por aqui don pirata?")

	// Aceptar la conexión TCP
	tcpConn, err := tcpListener.Accept()
	if err != nil {
		fmt.Println("Error al aceptar la conexión TCP:", err)
		return
	}
	defer tcpConn.Close()

	// Aquí puedes realizar las operaciones con la conexión TCP
	fmt.Println("Conexión TCP establecida con éxito.")
}
