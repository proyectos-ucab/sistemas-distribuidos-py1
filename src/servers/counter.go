package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type API int

var counter int

// Consultar contador
func (a *API) GetCounter(empty string, reply *int) error {
	*reply = counter
	return nil
}

// Reiniciar el contador
func (a *API) Restart(empty string, reply *int) error {
	counter = 0
	*reply = counter
	return nil
}

// Agregar un numero al contador
func (a *API) AddToCounter(number int, reply *int) error {
	counter += number
	*reply = counter
	fmt.Println("Se agregaron: ", number)
	return nil
}

// Restar un numero al contador
func (a *API) SubstractToCounter(number int, reply *int) error {
	counter -= number
	*reply = counter
	fmt.Println("Se restaron: ", number)
	return nil
}

func main() {
	api := new(API)
	err := rpc.Register(api)
	if err != nil {
		log.Fatal("error al registrar API", err)
	}

	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":4040")

	if err != nil {
		log.Fatal("Listener error", err)
	}
	log.Printf("RPC server corriendo en el puerto %d", 4040)
	http.Serve(listener, nil)

	if err != nil {
		log.Fatal("Error del servidor ", err)
	}

}
