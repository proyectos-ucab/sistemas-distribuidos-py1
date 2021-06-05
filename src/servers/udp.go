package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/rpc"
	"strconv"
	"strings"
	"time"
)

type Response struct {
	command string
	value   int
}

// Funcion random
func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func scanResponse(res string) *Response {

	var stringValues []string
	r := Response{command: "", value: 0}

	if strings.Contains(res, "INCREMENT") {
		stringValues = strings.Split(res, ",")
		r.command = "API." + stringValues[0]
		value, _ := strconv.Atoi(stringValues[1])
		r.value = value
	} else if strings.Contains(res, "DECREMENT") {
		stringValues = strings.Split(res, ",")
		r.command = "API." + stringValues[0]
		value, _ := strconv.Atoi(stringValues[1])
		r.value = value
	} else if strings.Contains(res, "GET") {
		r.command = "API.GET"
	} else if strings.Contains(res, "RESET") {
		r.command = "API.RESET"
	}

	return &r
}

func main() {

	var reply int

	PORT := ":" + "2002"

	s, err := net.ResolveUDPAddr("udp4", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Con esta linea tu te puedes conectar
	client, err := rpc.DialHTTP("tcp", "localhost:4040")

	if err != nil {
		log.Fatal("Connection error: ", err)
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

		var request = strings.TrimSpace(string(buffer[0:n]))

		var res = scanResponse(request)

		if res.command != "" {
			if res.value > 0 {
				client.Call(res.command, res.value, &reply)
			} else {
				client.Call(res.command, "", &reply)
			}
		}

		if strings.TrimSpace(string(buffer[0:n])) == "STOP" {
			fmt.Println("Desconectando servidor UDP")
			return
		}

		data := []byte(strconv.Itoa(reply))

		_, err = connection.WriteToUDP(data, addr)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
