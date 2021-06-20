package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {

	CONNECT := "127.0.0.1:2002"

	// Establecimiendo de la conexion al servidor UDP
	s, err := net.ResolveUDPAddr("udp4", CONNECT)
	c, err := net.DialUDP("udp4", nil, s)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Conectado al servidor UDP %s\n", c.RemoteAddr().String())
	defer c.Close()

	for {
		// Lectura en consola de los comandos del cliente 
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')

		// Enviar  mensaje hacia el servidor UDP`
		data := []byte(text + "\n")
		_, err = c.Write(data)


		if strings.TrimSpace(string(data)) == "STOP" {
			fmt.Println("Desconectando del cliente")
			return
		}

		if err != nil {
			fmt.Println(err)
			return
		}


		// Lectura de la respuesta del servidor servidor UDP`
		buffer := make([]byte, 1024)
		n, _, err := c.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println(err)
			return
		}


		fmt.Printf("Estado del contador: %s\n", string(buffer[0:n]))
	}
}
