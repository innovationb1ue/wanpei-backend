package services

import (
	"github.com/gorilla/websocket"
	"log"
	"wanpei-backend/mapper"
	"wanpei-backend/server"
)

type Socket struct {
	SocketMapper *mapper.Socket
	RedisMapper  *mapper.Redis
	MatchService *Match
	Settings     *server.Settings
}

func NewSocket(socketMapper *mapper.Socket, redisMapper *mapper.Redis, matchService *Match, settings *server.Settings) *Socket {
	return &Socket{
		SocketMapper: socketMapper,
		RedisMapper:  redisMapper,
		MatchService: matchService,
		Settings:     settings,
	}
}

func (s *Socket) AppendSocket(ID uint, ws *websocket.Conn) {
	s.SocketMapper.AddSocket(ID, ws)
}

func (s *Socket) StartHeartbeat(ID uint) {
	ws := s.SocketMapper.Sockets[ID]
	if ws == nil {
		return
	}
	// defer close
	defer func() {
		s.SocketMapper.DeleteSocket(ID)    // delete socket from map
		s.MatchService.RemoveFromQueue(ID) // delete user from queue
		err := ws.Close()                  // close websocket
		if err != nil {
			log.Println("error when closing websocket: ", err)
		}
		log.Println("Now queue = ", s.RedisMapper.GetAllFromQueue())
	}()
	// possible manually stop the heartbeat by close the [done] channel
	done := make(chan struct{})
	dead := make(chan struct{})
	// start heartbeats
	go ping(ws, done, dead, s.Settings)
	// block until heartbeat dead
	<-dead
	log.Println("a websocket died.")
	return
}
