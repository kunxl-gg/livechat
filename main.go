package main

import (
	"fmt"
	"golang.org/x/net/websocket"
)

type Server struct{
	connections map[string]bool
}

func NewSever() *Server{
	return &Server{
		connections: make(map[string]bool),
	}
}

func (serverInstance *Server) handleWS(ws *websocket.Conn){

	fmt.Println("New incoming connection from the client")
	serverInstance.connections["connection"] = true

	serverInstance.readLoop(ws)

} 

func (serverInstance *Server) readLoop(ws *websocket.Conn){
	buf := make([]byte, 1024);
	for {
		n, err := ws.Read(buf);
		if err != nil {
			fmt.Println("There was an error, ", err)
			continue 
		}

		message := buf[:n]
		fmt.Println(message);
	}
}

func main(){

	fmt.Print("This is working");

}