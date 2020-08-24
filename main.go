package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// WebSocketAddress has websocket properties
type WebSocketAddress struct {
	ip   string
	port int
}

var upgrader = websocket.Upgrader{}

func main() {
	ip := flag.String("i", "localhost", "Weboscket IP address")
	port := flag.Int("p", 4000, "Websocket Port")

	// below variables defines Mode of the program
	var serverMode bool
	var clientMode bool
	flag.BoolVar(&serverMode, "server", false, "using server mode")
	flag.BoolVar(&clientMode, "client", false, "using client mode")

	flag.Parse()
	WSdata := WebSocketAddress{ip: *ip, port: *port}

	// it decides to execute which mode according to the flags
	switch {
	case serverMode:
		server(WSdata)
	case clientMode:
		client(WSdata)
	default:
		fmt.Println("at first you have to define mode, use `-h` flag for more information")
	}
}

func client(WSdata WebSocketAddress) {
	wsa := fmt.Sprintf("ws://%s:%d", WSdata.ip, WSdata.port)
	fmt.Println(wsa)
}

func server(WSdata WebSocketAddress) {
	log.Printf("running in server mode: %v\n", WSdata)

	http.HandleFunc("/", echo)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", WSdata.ip, WSdata.port), nil))

}

// functions below is for usage of upper functions
func echo(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer conn.Close()
	for {
		mt, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		err = conn.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}
