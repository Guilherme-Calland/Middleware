package main

import(
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"strings"
)

type ClientManager struct{
	clients map[*Client]bool
	broadcast chan[] byte
	register chan *Client
	unregister chan *Client
}
/*client manager vai gerenciar todos os clientes*/
type Client struct{
	socket net.Conn
	data chan []byte
}
/**/
func (manager *ClientManager) start(){
	/*
	goroutine. se tiver dados no canal "register", uma nova conexão será
	estabelecida, se tiver dados passando pelo canal "unregister" e esses
	dados correspondem a alguma conexão ja existente essa conexão será
	fechada e removida da lista.
	caso o canal "broadcast tiver passando dados por ele então recebemos
	uma mensagem e a mensagem sera passada por todas as conexões"
	*/
	for{
		select{
		case connection := <- manager.register:
			manager.clients[connection] = true
			fmt.Println("New Connection!")
		case connection := <-manager.unregister:
			if _, ok := manager.clients[connection]; ok {
				close(connection.data)
				delete(manager.clients, connection)
				fmt.Println("A Connection Closed!")
			}
		case message := <- manager.broadcast:
			for connection := range manager.clients{
				select{
				case connection.data <- message:
				default:
					close(connection.data)
					delete(manager.clients, connection)
				}
			}
		}
	}
}

func (manager *ClientManager) receive(client *Client){
	/*
	funcao para o servidor servidor receber uma mensagem do cliente.
	se houver algum erro (if err) o cliente sera removido e a sua
	conexão fechará.
	se tudo der certo a mesnagem será addicionada ao "broadcast" e 
	distribuída para todos os clientes ClientManager manager
	*/
	for{
		message := make([]byte, 4096)
		length, err := client.socket.Read(message)
		if err != nil{
			manager.unregister <- client
			client.socket.Close()
			break
		}
		if length > 0{
			fmt.Println("Mensagem recebida: " + string(message))
			manager.broadcast <- message
		}
	}
}

func(client *Client) receive(){
	/*
	goroutine do cliente para receber dados.
	mensagem é recebida, se nao tiver erros, é imprimido
	*/
	for{
		message := make([]byte, 4096)
		length, err := client.socket.Read(message)
		if err != nil{
			client.socket.Close()
			break
		}
		if length > 0 {
			fmt.Println("Mensagem recebida: "+ string(message))
		}
	}
}

func (manager *ClientManager) send(client *Client){
	/*
	funcao que envia mensagens.
	se houver algum erro, saimos do loop (return) e a connexão 
	com o socket encerrará (defer client.socket.Close())
	*/
	defer client.socket.Close()
	
	for{
		select{
		case message, ok := <- client.data:
			if !ok {
				return
			}
			client.socket.Write(message)
		}
	}
}




func startServerMode(){
	/*
	aqui definimos em qual porta "escutaremos" por conexões.
	criaremos nosso gerenciador de clientes(manager) e rodaremos
	o goroutine do manager.
	entramos num loop escutando e aceitando conexões,
	caso uma conexão for aceita, será preparada para enviar e receber
	dados
	*/
	fmt.Println("Starting Server...")
	listener, error := net.Listen("tpc", ":12345")
	if error != nil {
		fmt.Println("Error!")
	}
	manager := ClientManager{
		clients: make(map[*Client]bool),
		broadcast: make(chan []byte),
		register: make(chan *Client),
		unregister: make(chan *Client),
	}
	go manager.start()
	for{
		connection , error := listener.Accept()
		if error != nil {
			fmt.Println("Error!")
		}
		client := &Client{socket: connection, data: make(chan []byte)}
		manager.register <- client
		go manager.receive(client)
		go manager.send(client)
	}
}


/*SE DER ERRO, MUDE OS ENDEREÇOS IP*/
func startClientMode(){
	/*
	se a conexão tiver sucesso, o cliente é criado e o cliente
	pode dar seu input quando o usuario der quebra de linha \n
	é sinal que a mensagem terminou e esta pronta pra envio
	*/
	fmt.Println("Starting client...")
	connection, error := net.Dial("tcp", ":12345")
	if error != nil {
		fmt.Print("Error!")
	}
	client := &Client{socket: connection}
	go client.receive()
	for{
		reader := bufio.NewReader(os.Stdin)
		message, _ := reader.ReadString('\n')
		connection.Write([]byte(strings.TrimRight(message, "\n")))
	}
}

//func startServerMode(){}
//func startClientMode(){}

func main(){
	flagMode := flag.String("mode", "server", "start in client or server mode")
	flag.Parse()
	if strings.ToLower(*flagMode) == "server" {
		startServerMode()
	} else {
		startClientMode()
	}
}
