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
	client, err := rpc.DialHTTP("tcp", "localhost:4040")

	if err != nil {
		log.Fatal("Connection error: ", err)
	}

	client.Call("API.AddToCounter", 10, &reply)
	client.Call("API.AddToCounter", 30, &reply)
	client.Call("API.SubstractToCounter", 5, &reply)
	client.Call("API.GetCounter", "", &reply)
	fmt.Println("Estado actual del contador", reply)
	client.Call("API.Restart", "", &reply)
	client.Call("API.GetCounter", "", &reply)
	fmt.Println("Estado actual del contador", reply)
}
