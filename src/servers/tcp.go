package main

import (
        "bufio"
        "fmt"
        "net"
        //"os"
        "strings"
        "time"
		"strconv"
		"net/rpc"
)

type Response struct {
	command string
	value   int64
}

func scanResponse(res string) *Response {

	var stringValues []string
	r := Response{command: "", value: 0}

	if strings.Contains(res, "INCREMENT") {
		stringValues = strings.Split(res, ",")
		fmt.Println(stringValues[1])
		r.command = "API." + stringValues[0]
		value, _ := strconv.ParseInt(stringValues[1],10,64)
		r.value = value
		fmt.Println(value)

	} else if strings.Contains(res, "DECREMENT") {
		stringValues = strings.Split(res, ",")
		r.command = "API." + stringValues[0]
		value, _ := strconv.ParseInt(stringValues[1],10,64)
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

				var res = scanResponse(string(netData))

				fmt.Println(res.command)
				fmt.Println(res.value)

				if res.command != "" {
					if res.value > 0 {
						client.Call(res.command, res.value, &reply)
					} else {
						client.Call(res.command, "", &reply)
					}
				}

                fmt.Print("-> ", string(netData))
                t := time.Now()
                myTime := t.Format(time.RFC3339) + "\n"
                c.Write([]byte(myTime))
        }
}
    
