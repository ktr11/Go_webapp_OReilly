package main

import (
	"github.com/gorilla/websocket"
)

// client はチャットを行っている1人のユーザーを表す
type client struct {
	// socketはこのクライアントのためのwebsocketです。
	socket *websocket.Conn
	// sendはメッセージが送られるチャネルです。
	send chan []byte
	// roomはこのクライアントが参加しているチャットルームです。
	room *room
}
