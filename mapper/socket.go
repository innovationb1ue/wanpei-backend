package mapper

import (
	"errors"
	"github.com/gorilla/websocket"
)

type SocketManager struct {
	Redis   *Redis
	Sockets map[string]*websocket.Conn // one websocket for one user is probably enough
}

func NewSocketManager(redis *Redis) *SocketManager {
	return &SocketManager{
		Redis:   redis,
		Sockets: make(map[string]*websocket.Conn),
	}
}

func (s *SocketManager) AddSocket(name string, socket *websocket.Conn) {
	s.Sockets[name] = socket
}

func (s *SocketManager) GetSocket(name string) (*websocket.Conn, error) {
	Socket := s.Sockets[name]
	if Socket == nil {
		return nil, errors.New("socket not found")
	}
	return Socket, nil
}
