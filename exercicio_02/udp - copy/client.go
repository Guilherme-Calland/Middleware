package main
import( 
	"fmt"
	"net"
	"bufio"
	"os"
	)
func main(){
	connection_01, _ := net.DialUDP("udp", nil, &net.UDPAddr{IP:[]byte{127,0,0,1},Port:1313,Zone:""})
	connection_02, _ := net.DialUDP("udp", nil, &net.UDPAddr{IP:[]byte{127,0,0,1},Port:8080,Zone:""})

	defer connection_01.Close()
	defer connection_02.Close()

	
	fmt.Println("Digite uma mensagem ao servidor_01: ")
	reader_01 := bufio.NewReader(os.Stdin)
	client_text_01, _ := reader_01.ReadString('\n')
	fmt.Fprint(connection_01, client_text_01)
	connection_01.Write([]byte(client_text_01))

	fmt.Println("Digite uma mensagem ao servidor_02: ")
	reader_02 := bufio.NewReader(os.Stdin)
	client_text_02, _ := reader_02.ReadString('\n')
	fmt.Fprintf(connection_02, client_text_02 )
	connection_02.Write([]byte(client_text_02))
	
}