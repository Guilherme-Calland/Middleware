package main

import(
	"fmt"
	"net"
	"bufio"
	"os"
)


func main(){

	//SERVER 01
	connection_01, error_01 := net.Dial("tcp", ":1001")
	if error_01 != nil {
		fmt.Println("Erro ao se conectar no servidor_01 !")
	}
	
	//SERVER 02
	connection_02, error_02 := net.Dial("tcp",":1313")
	if error_02 != nil{
		fmt.Println("Erro ao se conectar no servidor_02 !")
	}

	for{
		//le uma msg do teclado
		reader_01 := bufio.NewReader(os.Stdin)
		fmt.Print("\nDigite uma mensagem ao servidor_01 : ")
		client_text_01, _ := reader_01.ReadString('\n')
		//envia uma mensagem para o servidor
		fmt.Fprintf(connection_01, client_text_01 + "\n")
		//escuta por uma resposta do servidor
		server_text_01, _ := bufio.NewReader(connection_01).ReadString('\n')
		fmt.Print("Mensagem do Servidor_01: " + server_text_01)

		reader_02 := bufio.NewReader(os.Stdin)
		fmt.Print("\nDigite uma mensagem ao servidor_02 : ")
		client_text_02, _ := reader_02.ReadString('\n')
		fmt.Fprintf(connection_02, client_text_02 + "\n")
		//escuta por uma resposta do servidor
		server_text_02, _ := bufio.NewReader(connection_02).ReadString('\n')
		fmt.Print("Mensagem do Servidor_02: " + server_text_02)
	}

	connection_01.Close()
	connection_02.Close()
}
