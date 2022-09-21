package mapper

import (
	"errors"
	"github.com/gorilla/websocket"
	"wanpei-backend/server"
)

type Socket struct {
	Redis    *Redis
	Sockets  map[uint]*websocket.Conn // memory storage. can be optimized to redis if necessary
	Settings *server.Settings
}

func NewSocketManager(redis *Redis, settings *server.Settings) *Socket {
	return &Socket{
		Redis:    redis,
		Sockets:  make(map[uint]*websocket.Conn),
		Settings: settings,
	}
}

func (s *Socket) AddSocket(ID uint, socket *websocket.Conn) {
	s.Sockets[ID] = socket
}

func (s *Socket) GetSocket(ID uint) (*websocket.Conn, error) {
	socket := s.Sockets[ID]
	if socket == nil {
		return nil, errors.New("socket not found")
	}
	return socket, nil
}

func (s *Socket) DeleteSocket(ID uint) {
	delete(s.Sockets, ID)
}
