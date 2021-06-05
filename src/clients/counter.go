package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Item struct {
	Title string
	Body  string
}

func main() {
	var reply int

	// Con esta linea tu te puedes conectar
	client, err := rpc.DialHTTP("tcp", "localhost:4041")

	if err != nil {
		log.Fatal("Connection error: ", err)
	}

	client.Call("API.INCREMENT", 10, &reply)
	client.Call("API.INCREMENT", 30, &reply)
	client.Call("API.DECREMENT", 5, &reply)
	client.Call("API.GET", "", &reply)
	fmt.Println("Estado actual del contador", reply)
	client.Call("API.RESET", "", &reply)
	client.Call("API.GET", "", &reply)
	fmt.Println("Estado actual del contador", reply)
}
