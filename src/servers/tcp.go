package main

import (
	"bufio"
	"fmt"
	"net"
	"net/rpc"
	"strconv"
	"strings"
)

// Estructura de datos para recibir comandos
type ResponseTCP struct {
	command string
	value   int64
}

// Funcion para leer y obetner el valor de linea de comandos
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

	// Variable de respuesta del servidor
	var reply int


	// Establecimiento del puerto y Creacion del servidor TCP
	PORT := ":2020" 

	l, err := net.Listen("tcp", PORT)

	// Conexion del servidor TCP  al servidor de cuentas
	client, err := rpc.DialHTTP("tcp", "localhost:4040")


	// Validacion de conexion y manejo de errores
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
		// Lectura del input en del cliente
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		// Condicion para cerrar el servidor TCP
		if strings.TrimSpace(string(netData)) == "STOP" {
			fmt.Println("Cerrando el Servidor TCP")
			return
		}

		// Verificacion de comando para realizar operacions
		var res = scanResponseTCP(strings.TrimSpace(string(netData)))


		// LLamadas con el comando a proceder en el servidor de cuentas
		if res.command != "" {
			if res.value > 0 {
				client.Call(res.command, res.value, &reply)
			} else {
				client.Call(res.command, "", &reply)
			}
		}

		// Respuesta al cliente TCP
		data := "Estado del contador:" + strconv.Itoa(reply) + "\n"

		c.Write([]byte(data))
	}
}
