package main

import (
	"fmt"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"time"
)

// Funcion random
func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func main() {

	PORT := ":" + "2002"

	s, err := net.ResolveUDPAddr("udp4", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}

	connection, err := net.ListenUDP("udp4", s)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("UDP server corriendo en el puerto 2002")

	defer connection.Close()
	buffer := make([]byte, 1024)
	rand.Seed(time.Now().Unix())

	for {
		n, addr, err := connection.ReadFromUDP(buffer)
		fmt.Print("-> ", string(buffer[0:n-1]))

		var response = strings.TrimSpace(string(buffer[0:n]))
		fmt.Println("Comando:", response)

		if strings.TrimSpace(string(buffer[0:n])) == "STOP" {
			fmt.Println("Desconectando servidor UDP")
			return
		}

		data := []byte(strconv.Itoa(random(1, 1001)))
		fmt.Printf("data: %s\n", string(data))

		_, err = connection.WriteToUDP(data, addr)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
