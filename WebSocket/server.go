package main

import (
    "fmt"
    "log"
    "net/http"

    "github.com/gorilla/websocket"
)

var addr = "localhost:8080"

var upgrader = websocket.Upgrader{}

func wsHello(w http.ResponseWriter, r *http.Request) {
    upgrader.CheckOrigin = func (r *http.Request) bool { return true }
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Print("upgrade error:", err)
        return
    }
    defer conn.Close()
    for {
        messageType, message, err := conn.ReadMessage()
        if err != nil {
            log.Println("read error:", err)
            break
        }
        log.Printf("recv: %s", message)
        response := []byte(fmt.Sprintf("Hello %s", message))
        err = conn.WriteMessage(messageType, response)
        if err != nil {
            log.Println("write error:", err)
            break
        }
    }
}

func main() {
	http.HandleFunc("/ws/hello", wsHello)
    log.Fatal(http.ListenAndServe(addr, nil))
}
