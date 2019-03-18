package main

import(
    "net"
    "fmt"
)

func main(){
	server_connection, _ := net.ListenUDP("udp", &net.UDPAddr{IP:[]byte{0,0,0,0},Port:1313,Zone:""})
	defer server_connection.Close()
	buf := make([]byte, 1024)
	n, addr, _ := server_connection.ReadFromUDP(buf)
	fmt.Println("Mensagem do cliente: ", string(buf[0:n]), "do endereco IP:", addr)
	
}


