package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {

	CONNECT := "127.0.0.1:2020"

	// Establecimiendo de la conexion al servidor TCP
	c, err := net.Dial("tcp", CONNECT)

	// Validacion de la conexion
	if err != nil {
		fmt.Println(err)
		return
	}

	for {

		// Lectura en consola de los comandos del cliente 
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')

		// Enviar  mensaje hacia el servidor TCP
		fmt.Fprintf(c, text+"\n")

		// Lectura de la respeusta del servidor UDP
		message, _ := bufio.NewReader(c).ReadString('\n')

		fmt.Print("-> " + message)
		if strings.TrimSpace(string(text)) == "STOP" {
			fmt.Println("Desconectando cliente TCP...")
			return
		}
	}
}
