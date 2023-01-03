package services

import (
	"context"
	"errors"
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
	// flush all in the queue when restarting the match service
	p.RedisMapper.FlushAll()
	return &Match{
		UserMapper:   p.UserMapper,
		RedisMapper:  p.RedisMapper,
		SocketMapper: p.SocketMapper,
		Settings:     p.Settings,
	}
}

func (m *Match) CheckUserInMatchPool(ID uint) (bool, error) {
	ctx := context.Background()
	_, err := m.RedisMapper.Client.LPos(ctx, m.Settings.RedisMatchMakingUsersQueueName, strconv.Itoa(int(ID)), redis.LPosArgs{
		Rank:   1,
		MaxLen: 0,
	}).Result()
	if err == redis.Nil {
		return false, nil
	}
	if err != nil {
		return false, errors.New("unknown error with redis. check redis connection or version. ")
	}
	return true, nil
}

// AppendToQueue append a user ID to Redis queue if it does not exist in queue
func (m *Match) AppendToQueue(ID uint) (bool, error) {
	isExist, err := m.CheckUserInMatchPool(ID)
	if err != nil {
		return false, err
	}
	if isExist {
		log.Println("user already in match pool")
		return false, nil
	}
	m.RedisMapper.AppendUserToMatchPool(ID)
	return true, nil
}

// RemoveFromQueue remove the user from match making queue. If user do not exist, it is a no-op.
func (m *Match) RemoveFromQueue(ID uint) {
	_ = m.RedisMapper.RemoveUserFromMatchPool(ID)
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

func (m *Match) QueueLength() (uint, error) {
	ctx := context.Background()
	res, err := m.RedisMapper.Client.LLen(ctx, m.Settings.RedisMatchMakingUsersQueueName).Result()
	if err == redis.Nil {
		return 0, nil
	}
	if err != nil {
		return 0, errors.New("unknown error when getting queue length. ")
	}
	return uint(res), nil

}
