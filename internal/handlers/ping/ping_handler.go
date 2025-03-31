package ping

import (
	"net/http"

	"github.com/gorilla/websocket"
)

type PingHandler struct {
}

const (
	socketBufferSize  = 1024
	messageBufferSize = 1024
)

var upgrader = &websocket.Upgrader{ReadBufferSize: socketBufferSize, WriteBufferSize: socketBufferSize}

type Client struct {
	conn    *websocket.Conn
	send    chan []byte
	channel *Channel
}

type Channel struct {
	clients map[*Client]bool
	join    chan *Client
	leave   chan *Client
	forward chan []byte
}

// TODO: implement this once things are working
type Message struct {
	from    *Client
	content []byte
}

func (ph *PingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("talking to the ws"))
	// conn, err := upgrader.Upgrade(w, r, nil)
	// if err != nil {
	// 	return
	// }

	// client := Client{conn: conn, send: make([]byte, messageBufferSize, channel: )}
}
