package main

import (
	"github.com/gorilla/websocket"
)

// read
// クライアントがWebSocketからReadMessageを使ってデータを読み込むために使われます。
func (c *client) read() {
	for {
		if _, msg, err := c.socket.ReadMessage(); err == nil {
			c.room.forward <- msg
		} else {
			break
		}
		c.socket.Close()
	}
}

// write
// 継続的にsendチャネルからメッセージを受け取り、WebSocketのWriteMessageメソッドを使ってこれを書き出します。
func (c *client) write() {
	for msg := range c.send {
		if err := c.socket.WriteMessage(websocket.TextMessage, msg); err != nil {
			break
		}
	}
	c.socket.Close()
}

// client はチャットを行っている1人のユーザーを表す
type client struct {
	// socketはこのクライアントのためのwebsocketです。
	socket *websocket.Conn
	// sendはメッセージが送られるチャネルです。
	send chan []byte
	// roomはこのクライアントが参加しているチャットルームです。
	room *room
}
