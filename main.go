package main

import (
	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	WriteBufferSize: 1024,
	ReadBufferSize:  1024,
}

func main() {
	envFile, _ := godotenv.Read(".env")
	listenPort := envFile["LISTEN_PORT"]

	http.HandleFunc("/wsping", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println("Upgrade error:", err)
			return
		}
		defer func(conn *websocket.Conn) {
			err := conn.Close()
			if err != nil {

			}
		}(conn)

		for {
			messageType, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("Read error:", err)
				break
			}
			log.Printf("Received: %s", message)

			err = conn.WriteMessage(messageType, message)
			if err != nil {
				log.Println("Write error:", err)
				break
			}
		}
	})

	log.Printf("Server started on :%s", listenPort)
	err := http.ListenAndServe(":"+listenPort, nil)
	if err != nil {
		log.Fatal("ListenAndServe error:", err)
	}
}
