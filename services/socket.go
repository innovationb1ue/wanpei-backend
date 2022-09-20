package services

import (
	"github.com/gorilla/websocket"
	"wanpei-backend/mapper"
)

type Socket struct {
	SocketMapper *mapper.Socket
}

func NewSocket(socketMapper *mapper.Socket) *Socket {
	return &Socket{
		SocketMapper: socketMapper,
	}
}

func (s *Socket) AppendSocket(ID uint, ws *websocket.Conn) {
	s.SocketMapper.AddSocket(ID, ws)
}
