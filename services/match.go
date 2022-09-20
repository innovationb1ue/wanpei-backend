package services

import (
	"context"
	"github.com/go-redis/redis/v9"
	"github.com/gorilla/websocket"
	"go.uber.org/fx"
	"log"
	"strconv"
	"time"
	"wanpei-backend/mapper"
	"wanpei-backend/server"
)

type MatchIn struct {
	fx.In
	UserMapper   *mapper.User
	RedisMapper  *mapper.Redis
	SocketMapper *mapper.Socket
	Settings     *server.Settings
}

type Match struct {
	UserMapper   *mapper.User
	RedisMapper  *mapper.Redis
	SocketMapper *mapper.Socket
	Settings     *server.Settings
}

func NewMatch(p MatchIn) *Match {
	return &Match{
		UserMapper:   p.UserMapper,
		RedisMapper:  p.RedisMapper,
		SocketMapper: p.SocketMapper,
		Settings:     p.Settings,
	}
}

func (m *Match) CheckUserInMatchPool(ID uint) bool {
	ctx := context.Background()
	_, err := m.RedisMapper.Client.LPos(ctx, m.Settings.RedisMatchMakingUsersQueueName, strconv.Itoa(int(ID)), redis.LPosArgs{
		Rank:   1,
		MaxLen: 0,
	}).Result()
	if err == redis.Nil {
		return false
	}
	return true
}

// AppendToQueue append a user ID to Redis queue if it does not exist in queue
func (m *Match) AppendToQueue(ID uint) bool {
	isExist := m.CheckUserInMatchPool(ID)
	if isExist {
		log.Println("user already in match pool")
		return false
	}
	m.RedisMapper.AppendUserToMatchPool(ID)
	return true
}
func (m *Match) RemoveFromQueue(ID uint) {
	m.RedisMapper.RemoveUserFromMatchPool(ID)
}

func (m *Match) StartHeartbeat(ID uint) {
	ws := m.SocketMapper.Sockets[ID]
	if ws == nil {
		return
	}
	// defer close
	defer func() {
		m.SocketMapper.DeleteSocket(ID) // delete socket from map
		m.RemoveFromQueue(ID)           // delete user from queue
		err := ws.Close()               // close websocket
		if err != nil {
			log.Println("error when closing websocket: ", err)
		}
		log.Println("Now queue = ", m.RedisMapper.GetAllFromQueue())
	}()
	// possible manually stop the heartbeat by close the [done] channel
	done := make(chan struct{})
	dead := make(chan struct{})
	// start heartbeats
	go ping(ws, done, dead, m.Settings)
	// block until heartbeat dead
	<-dead
	log.Println("a websocket died.")
	return
}

func ping(ws *websocket.Conn, done chan struct{}, dead chan struct{}, settings *server.Settings) {
	ticker := time.NewTicker(settings.HeartbeatPeriod)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			if err := ws.WriteControl(websocket.PingMessage, []byte{}, time.Now().Add(settings.WriteWait)); err != nil {
				log.Println("ping:", err)
				close(dead)
				return
			}
		case <-done:
			return
		}
	}
}
