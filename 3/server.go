package main

import (
	"fmt"
	"log"
	"net"
	"strings"
)

var DNS = []string{}

func agregarRegistro(dominio string, ipAddr string, TTL string, tipo string) {
	// var temp = []string{dominio, ipAddr, TTL, tipo}
	DNS = append(DNS, dominio, ipAddr, TTL, tipo)
}

func obtenerRegistro(dominio string) string {
	for i := 0; i < len(DNS); i = i + 4 {
		if DNS[i] == dominio {
			return DNS[i+1]
		}
	}
	return "No se pudo encontrar el dominio"
}

type DNS_list struct {
	dominio string
	ipAddr  string
	TTL     string
	tipo    string
}

func main() {

	conexion, err := net.ListenPacket("udp", "localhost:63420")
	if err != nil {
		log.Fatal(err)
	}

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
		// fmt.Printf("Mensaje recibido desde %s: %s\n", clientAddr, string(buffer[:n]))

		var dominio string
		var ip string
		var ttl string
		var tipo string

		msg := string(buffer[:n])

		if msg == "A" {
			mensaje := []byte("Listo para agregar un registro")
			_, err = conexion.WriteTo(mensaje, clientAddr)
			if err != nil {
				fmt.Println("Error al responder al cliente UDP:", err)
				continue
			}
		}

		if msg == "O" {
			mensaje := []byte("Listo para enviar la IP")
			_, err = conexion.WriteTo(mensaje, clientAddr)
			if err != nil {
				fmt.Println("Error al responder al cliente UDP:", err)
				continue
			}
		}

		n, clientAddr, err = conexion.ReadFrom(buffer)
		if err != nil {
			fmt.Println("Error al leer datos UDP:", err)
			continue
		}

		dominio = strings.Split(string(buffer[:n]), " ")[0]
		ip = strings.Split(string(buffer[:n]), " ")[1]
		ttl = strings.Split(string(buffer[:n]), " ")[2]
		tipo = strings.Split(string(buffer[:n]), " ")[3]

		if msg == "A" {
			agregarRegistro(dominio, ip, ttl, tipo)
		}
		if msg == "O" {
			mje := obtenerRegistro(dominio)
			mensaje := []byte(mje)
			_, err = conexion.WriteTo(mensaje, clientAddr)
			if err != nil {
				fmt.Println("Error al responder al cliente UDP:", err)
				continue
			}

		}

		// Obtener la dirección IP del cliente UDP
		// clientIP := strings.Split(clientAddr.String(), ":")[0]

		mensaje := []byte("Hola desde el server")
		_, err = conexion.WriteTo(mensaje, clientAddr)
		if err != nil {
			fmt.Println("Error al responder al cliente UDP:", err)
			continue
		}

		// Cerrar la conexión UDP
		// conexion.Close()
		break // Salir del bucle después de cerrar la conexión
	}
	fmt.Println("Conexión UDP cerrada")
	conexion.Close()
}
