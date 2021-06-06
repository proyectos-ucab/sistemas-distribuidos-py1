package main

import (
	"bufio"
	"fmt"
	"net"
	"net/rpc"
	"strconv"
	"strings"
)

type ResponseTCP struct {
	command string
	value   int64
}

func scanResponseTCP(res string) *ResponseTCP {

	var stringValues []string
	r := ResponseTCP{command: "", value: 0}

	if strings.Contains(res, "INCREMENT") {
		stringValues = strings.Split(res, ",")
		r.command = "API." + stringValues[0]
		value, _ := strconv.ParseInt(stringValues[1], 10, 64)
		r.value = value

	} else if strings.Contains(res, "DECREMENT") {
		stringValues = strings.Split(res, ",")
		r.command = "API." + stringValues[0]
		value, _ := strconv.ParseInt(stringValues[1], 10, 64)
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
	//arguments := os.Args
	/* if len(arguments) == 1 {
	        fmt.Println("Ingresa  el numero del puerto: ")
	        return
	} */

	PORT := ":2020" //+ arguments[1]
	l, err := net.Listen("tcp", PORT)

	client, err := rpc.DialHTTP("tcp", "localhost:4040")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	fmt.Println("[TCP Server]: Corriendo en el puerto: " + PORT)

	c, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}

	for {

		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		if strings.TrimSpace(string(netData)) == "STOP" {
			fmt.Println("Cerrando el Servidor TCP")
			return
		}

		var res = scanResponseTCP(strings.TrimSpace(string(netData)))

		if res.command != "" {
			if res.value > 0 {
				client.Call(res.command, res.value, &reply)
			} else {
				client.Call(res.command, "", &reply)
			}
		}

		data := "Estado del contador:" + strconv.Itoa(reply) + "\n"

		c.Write([]byte(data))
	}
}
