# Proyecto 1 - Sistemas distribuidos

## Integrantes

- Nestor Angeles
- Fernando Garcia
- Felix Lopez

## Desarrollo del proyecto

- [x] Servidor y cliente UDP
- [x] Servidor y cliente TCP por procesos
- [ ] Servidor y cliente TCP por hilos
- [x] Servidor de cuentas RCP
- [ ] Consola remota
- [ ] Consola local

## Servidor y cliente UPD

### Servidor

El servidor corre en el puerto `2002`, para iniciarlo debemos cambiar el directorio a `./src/servers` y correr el comando `go run upd.go`

utiliza la libreria [net en el lenguaje de golang](https://golang.org/pkg/net/) para realizar la conexion

- La funcion `ResolveUDPAddr` devuelve una dirección de punto final UDP, lo cual sirve para establecer la conexion

- La funcion `ListenUDP` actúa como ListenPacket para redes UDP, el cual se utiliza para establecer el escucha de eventos

### Cliente

Se conecta automaticamente al puerto en el que esta corriendo el servidor, `2002`, para iniciarlo debemos cambiar el directorio a `./src/clients` y correr el comando `go run upd.go`

## Servidor y cliente TCP

### Servidor

El servidor corre en el puerto `2020`, para iniciarlo debemos cambiar el directorio a `./src/servers` y correr el comando `go run tcp.go`

Al igual que el servidor UPD utilza la libreria `net` para realiar la conexion especificamente la funcion `Listen`

### Cliente

Se conecta automaticamente al puerto en el que esta corriendo el servidor, `2020`, para iniciarlo debemos cambiar el directorio a `./src/clients` y correr el comando `go run tcp.go`

## Servidor de cuentas

Es un servidor RCP que corre en el puerto `4040` el cual utiliza la sublibreria [net/rcp](https://golang.org/pkg/net/rpc/) para realizar la conexion, para iniciarlo debemos cambiar el directorio a `./src/servers` y correr el comando `go run counter.go`

## Como interactuar con el servidor de cuentas a traves de un cliente (TCP o UDP)

Tenemos cuatro operaciones basicas `INCREMENT`, `DECREMENT`, `GET`, `RESET`

* INCREMENT: Incrementa el contador en el servidor de cuentas con un valor dado, para suminstrar el valor debemos escribir el comando separado de una coma de la siguiente manera: `INCREMENT,10`

* DECREMENT: Decrementa el contador en el servidor de cuentas con un valor dado, para suministrar el valodr debemos escribir el comando separado de una coma de la siguiente manera: `DECREMENT,5`

* GET: Obtiene el estado actual del contador

* RESET: Reinicia el contador a 0

Estos comand funcionan para ambos clientes 
