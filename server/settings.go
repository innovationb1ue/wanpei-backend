package server

import "time"

type Settings struct {
	Secret     string
	Addr       string
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

func NewSettings() *Settings {
	// define all the possible user settings here
	return &Settings{
		Secret:                         "ggbob",
		Addr:                           ":8096",
		Sha256Salt:                     "salt123",
		RedisAddr:                      "localhost:6379",
		RedisPassword:                  "",
		RedisMatchMakingUsersQueueName: "match:users",
		PongWait:                       5 * time.Second,
		WriteWait:                      5 * time.Second,
		MaxMessageSize:                 8192,
		HeartbeatPeriod:                3 * time.Second,
		MaxAge:                         60 * 60 * 24 * time.Second, // 24hrs expire
	}
}
