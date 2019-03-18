package main

import(
	"fmt"
	"net"
	"bufio"
)

func main(){
	fmt.Println("inicando servidor...")

	ln, _ := net.Listen("tcp", ":1313")
	conn, _ := ln.Accept()

	for{
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Println("Mensagem do cliente: " + message)
		new_message := "Ok"
		conn.Write([]byte(new_message ))
	}
	conn.Close()
}