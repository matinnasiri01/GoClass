package main

import (
	"fmt"
	"net/http"
	"slices"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var upgrader = websocket.Upgrader{}

type User struct {
	Username string
	Ws *websocket.Conn
}
var rooms = make(map[string][]User)
var roomMu sync.Mutex

func wsChatRoom(c echo.Context) error {
	room := c.Param("roomId")
	username := c.Param("username")

	if room == "" && username == "" {
		c.Logger().Error("room and username are required")
		return nil
	}

	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, er := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if er != nil {
		c.Logger().Error(er)
		return nil
	}
	defer func (){
		for i,user := range rooms[room]{
			if user.Username == username {
				rooms[room] = slices.Delete(rooms[room],i,i+1)
				break
			}
		}
		ws.Close()
	}()

	roomMu.Lock()
	rooms[room] = append(rooms[room], User{Username: username,Ws: ws})
	roomMu.Unlock()


	for {
		_, message, err := ws.ReadMessage()
		if err != nil {
			c.Logger().Print("read error:", err)
			break
		}
		for _,user := range rooms[room]{
			if user.Username == username {
				continue
			}
			user.Ws.WriteMessage(websocket.TextMessage,[]byte(fmt.Sprintf("%s: %s", username, message)))
			
		}
	}

	return nil
}

func main() {
	e := echo.New()
	e.GET("/ws/chat/:roomId/user/:username", wsChatRoom)
	e.Logger.Fatal(e.Start(":8080"))
}