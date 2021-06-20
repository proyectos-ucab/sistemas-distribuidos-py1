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

// Estructura de datos para recibir comandos
type Response struct {
	command string
	value   int
}

// Funcion para leer y obetner el valor de linea de comandos
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

	// Variable de respuesta para el servidor
	var reply int

	PORT := ":" + "2002"

	// Creacion del servidor UDP
	s, err := net.ResolveUDPAddr("udp4", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Conexion del servidor TCP  al servidor de cuentas
	client, err := rpc.DialHTTP("tcp", "localhost:4040")


	// Validacion de la conexion al servidor de cuentas
	if err != nil {
		log.Fatal("Connection error: ", err)
	}

	// Inicializacion del paquete de listeners para la conexiones UDP
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

		// Lectura del input del cliente
		n, addr, err := connection.ReadFromUDP(buffer)

		var request = strings.TrimSpace(string(buffer[0:n]))

		// Verificacion de comando para realizar operacions
		var res = scanResponse(request)

		// LLamadas con el comando a proceder en el servidor de cuentas
		if res.command != "" {
			if res.value > 0 {
				client.Call(res.command, res.value, &reply)
			} else {
				client.Call(res.command, "", &reply)
			}
		}


		// Condicion para cerrar el servidor TCP
		if strings.TrimSpace(string(buffer[0:n])) == "STOP" {
			fmt.Println("Desconectando servidor UDP")
			return
		}


		// Respuesta al cliente UPD
		data := []byte(strconv.Itoa(reply))

		_, err = connection.WriteToUDP(data, addr)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
