# Importo la libreria socket para manejar conexiones
import socket as skt

# Se declara dirección del servidor y puerto
serverAddr = 'localhost'
serverPort = 63420

# Se crea un socket para hacer el manejo de la conexión
clientSocket = skt.socket(skt.AF_INET, skt.SOCK_DGRAM)

# Solicito mensajes y luego ejecuto las funciones para conectar al servidor y realizar el intercambio de mensajes.
toSend = input("Para agregar un registro pulse A. Para obtener pulse O. Para salir escriba STOP\n>")
toSend = toSend.upper()
if toSend == "STOP":
	clientSocket.sendto(toSend.encode(), (serverAddr,serverPort))
	print("Programa Finalizado")

elif toSend == "A":
	clientSocket.sendto(toSend.encode(), (serverAddr,serverPort))
	msg, addr = clientSocket.recvfrom(1024)
	print(msg.decode())

	toSend = input("Agregue el dominio, ip, ttl y el tipo separados por espacios\n")
	clientSocket.sendto(toSend.encode(), (serverAddr,serverPort))

elif toSend == "O":
	clientSocket.sendto(toSend.encode(), (serverAddr,serverPort))
	msg, addr = clientSocket.recvfrom(1024)
	print(msg.decode())

	toSend = input("Escriba el dominio del cuál desea obtener la IP\n")
	clientSocket.sendto(toSend.encode(), (serverAddr, serverPort))
	msg, addr = clientSocket.recvfrom(1024)
	print(msg.decode())
else:
	clientSocket.sendto(toSend.encode(), (serverAddr,serverPort))

	msg, addr = clientSocket.recvfrom(1024)
	print(msg.decode()) 