package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

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
	flag.BoolVar(&serverMode, "server", false, "using server mode")
	var clientMode bool
	flag.BoolVar(&clientMode, "client", false, "using client mode")

	flag.Parse()
	WSdata := WebSocketAddress{ip: *ip, port: *port}

	// it decides to execute which mode according to the flags
	switch {
	case serverMode:
		Server(WSdata)
	case clientMode:
		Client(WSdata)
	default:
		fmt.Println("at first you have to define mode, use `-h` flag for more information")
	}
}

func Client(WSdata WebSocketAddress) {
	wsa := fmt.Sprintf("wss://%s:%d", WSdata.ip, WSdata.port)
	fmt.Println(wsa)
}

func Server(WSdata WebSocketAddress) {

	wsa := fmt.Sprintf("%s:%d/", WSdata.ip, WSdata.port)

	fmt.Println(wsa)
	fmt.Println("server")
	http.HandleFunc(wsa, echo)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", WSdata.port), nil))

}

// functions below is for usage of upper functions
func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}
