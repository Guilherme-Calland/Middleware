package main

import(
	"fmt"
	"net"
	"bufio"
)

func main(){
	fmt.Println("inicando servidor...")

	ln, _ := net.Listen("tcp", ":8080")
	conn, _ := ln.Accept()

	for{
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Println("Mensagem do cliente: " + message)
		conn.Write([]byte("Ok\n"))
	}
	
	conn.Close()
}
