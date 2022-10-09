package server

import (
	"time"
)

type Settings struct {
	CookieSecret string
	// server listen address
	Addr  string
	https bool
	// pwd sha256 salt
	Sha256Salt string
	//Redis settings
	RedisAddr     string
	RedisPassword string

	// match making settings
	RedisMatchMakingUsersQueueName string

	// socket settings
	PongWait        time.Duration
	WriteWait       time.Duration
	MaxMessageSize  int
	HeartbeatPeriod time.Duration

	// session settings
	MaxAge time.Duration
}

func NewSettings(env *Env) *Settings {
	//todo: refactor this to all env configs
	if env.Vars["APP_TARGET"] == "prod" {
		// production settings
		return &Settings{
			CookieSecret:                   "ggbob",
			Addr:                           ":8096",
			https:                          true,
			Sha256Salt:                     "salt123",
			RedisAddr:                      "redis:6379",
			RedisPassword:                  "",
			RedisMatchMakingUsersQueueName: "match:users",
			PongWait:                       5 * time.Second,
			WriteWait:                      5 * time.Second,
			MaxMessageSize:                 8192,
			HeartbeatPeriod:                3 * time.Second,
			MaxAge:                         60 * 60 * 24 * time.Second, // 24hrs expire
		}
	} else {
		// test settings
		return &Settings{
			CookieSecret:                   "ggbob",
			Addr:                           ":8096",
			https:                          false,
			Sha256Salt:                     "salt123",
			RedisAddr:                      "127.0.0.1:6379",
			RedisPassword:                  "",
			RedisMatchMakingUsersQueueName: "match:users",
			PongWait:                       5 * time.Second,
			WriteWait:                      5 * time.Second,
			MaxMessageSize:                 8192,
			HeartbeatPeriod:                3 * time.Second,
			MaxAge:                         60 * 60 * 24 * time.Second, // 24hrs expire
		}
	}
}
