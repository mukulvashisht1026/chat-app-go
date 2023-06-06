package main

import (
	"fmt"
	"log"
	"net/http"

	"io"

	"mukulvashisht1026/chat-app-go/services"

	"github.com/gorilla/mux"
	"golang.org/x/net/websocket"
)

// code for web sockets
type Server struct {
	conns map[*websocket.Conn]bool
}

func NewServer() *Server {
	return &Server{
		conns: make(map[*websocket.Conn]bool),
	}
}

func (s *Server) handleWS(ws *websocket.Conn) {
	fmt.Println("new incomming connection from sli: ", ws.RemoteAddr)
	s.conns[ws] = true
	s.readLoop(ws)
}

func (s *Server) readLoop(ws *websocket.Conn) {
	buf := make([]byte, 1024)
	for {
		n, err := ws.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("read error:", err)
			continue
		}
		msg := buf[:n]
		fmt.Println(string(msg))
		ws.Write([]byte("thank you for the message!!"))
	}
}

// code for rest API

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage endpoint is hit")

}

func handleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", services.AllArticles)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// code for rest API with gorilla MUX
func handleRequestWithMux() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage).Methods("GET")
	myRouter.HandleFunc("/articles", services.AllArticles).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {

	// code for web sockets
	/*
		fmt.Println("hello world")
		server := NewServer()
		http.Handle("/ws", websocket.Handler(server.handleWS))
		http.ListenAndServe(":3000", nil)
	*/
	// code for rest api
	handleRequestWithMux()

}
