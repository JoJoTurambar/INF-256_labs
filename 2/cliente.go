package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	// Resuelve la dirección UDP del servidor
	udpAddr, err := net.ResolveUDPAddr("udp", "localhost:8000")
	if err != nil {
		fmt.Println("Error al resolver la dirección del servidor:", err)
		os.Exit(1)
	}

	// Crea una conexión UDP
	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		fmt.Println("Error al crear la conexión:", err)
		os.Exit(1)
	}
	defer conn.Close()

	// Envía un mensaje al servidor
	message := []byte("¡Hola, servidor!")
	_, err = conn.Write(message)
	if err != nil {
		fmt.Println("Error al enviar el mensaje:", err)
		os.Exit(1)
	}

	fmt.Println("Mensaje enviado al servidor:", string(message))

	// Espera la respuesta del servidor
	buffer := make([]byte, 1024)
	n, _, err := conn.ReadFromUDP(buffer)
	if err != nil {
		fmt.Println("Error al recibir la respuesta del servidor:", err)
		os.Exit(1)
	}

	fmt.Println("Respuesta del servidor:", string(buffer[:n]))
	//realizar conexion TCP
	tcpAddr := string(buffer[:n])
	tcpConn, err := net.Dial("tcp", tcpAddr)
	if err != nil {
		fmt.Println("Error al establecer la conexión TCP:", err)
		os.Exit(1)
	}
	defer tcpConn.Close()
	fmt.Println("Conexión TCP establecida con éxito.")

	//muestra el mapa del jugador
	//....

	//while(no ganador)
	//pregunta en qué posicion quiere atacar
	var confirm string
	var ataque string
	fmt.Println("precione q para salir del juego y cualquier otra letra para continuar:")
	fmt.Scanln(&confirm)
	if confirm != "q" {
		fmt.Println("Ingrese la ubicacion donde desee lanzar un ataque (Ubicaciones posibles: A, B, C y D):")
		fmt.Scanln(&ataque)

		// Envía el ataque al servidor a través de la conexión TCP
		_, err = tcpConn.Write([]byte(ataque))
		if err != nil {
			fmt.Println("Error al enviar el ataque al servidor:", err)
			return
		}
		fmt.Println("Ataque enviado al servidor:", ataque)
	}
}
