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
func (a *API) GET(empty string, reply *int) error {
	*reply = counter
	fmt.Println("Contador en:", counter)
	return nil
}

// Reiniciar el contador
func (a *API) RESET(empty string, reply *int) error {
	counter = 0
	*reply = counter
	fmt.Println("Se reinicio el contador a 0")
	return nil
}

// Agregar un numero al contador
func (a *API) INCREMENT(number int, reply *int) error {
	counter += number
	*reply = counter
	fmt.Println("Se agregaron: ", number)
	return nil
}

// Restar un numero al contador
func (a *API) DECREMENT(number int, reply *int) error {
	counter -= number
	*reply = counter
	fmt.Println("Se restaron: ", number)
	return nil
}

func main() {

	// API Register para realizar llamados RCP
	api := new(API)
	err := rpc.Register(api)
	if err != nil {
		log.Fatal("error al registrar API", err)
	}

	rpc.HandleHTTP()

	// Creacion del servidor de cuentas con RCP
	listener, err := net.Listen("tcp", ":4040")

	// Validacion de errores
	if err != nil {
		log.Fatal("Listener error", err)
	}
	log.Printf("RPC server corriendo en el puerto %d", 4040)

	// Lectura de las llamadas del servidor
	http.Serve(listener, nil)

	if err != nil {
		log.Fatal("Error del servidor ", err)
	}

}
