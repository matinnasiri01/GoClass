package main

import (
    "log"
    "github.com/gorilla/websocket"
)


func main() {
    url := "ws://localhost:8080/ws/hello"
    log.Printf("connecting to %s", url)

    // 1. Dial 
    connection, _, err := websocket.DefaultDialer.Dial(url, nil)
    if err != nil {
	    log.Fatal("dial error:", err)
	}
	// 2. Close
	defer connection.Close()
	
	sendMessage := "Quera college"
	// 3. Write mess
	err = connection.WriteMessage(websocket.TextMessage, []byte(sendMessage))
	if err != nil {
		log.Println("write error:", err)
		} else {
			log.Println("send message:", sendMessage)
		}
		// 3. Read mess
		_, message, err := connection.ReadMessage()
    if err != nil {
        log.Println("read error:", err)
        return
    }
    log.Printf("received message: %s", message)
}