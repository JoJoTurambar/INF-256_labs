package main

import (
	"fmt"
	"net"
)

func main() {
	conexion, err := net.Dial("udp", "localhost:8000")
	if err != nil {
		fmt.Println("Error al conectar al servidor:", err)
		return
	}
	defer conexion.Close()
	mensaje := []byte("Hola desde el cliente") //modificar el mensaje para pedir los datos de la conexión TCP
	// Enviar mensaje al servidor
	_, err = conexion.Write(mensaje)
	if err != nil {
		fmt.Println("Error al enviar mensaje:", err)
		return
	}
	fmt.Println("Mensaje enviado al servidor:", string(mensaje))
	// Buffer para almacenar la respuesta del servidor
	buffer := make([]byte, 1024)
	// Leer respuesta del servidor
	n, err := conexion.Read(buffer)
	if err != nil {
		fmt.Println("Error al recibir respuesta del servidor:", err)
		return
	}
	fmt.Println("Respuesta del servidor:", string(buffer[:n]))
}
