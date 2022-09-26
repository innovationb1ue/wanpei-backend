package mapper

import (
	"context"
	"github.com/go-redis/redis/v9"
	"log"
	"strconv"
	"wanpei-backend/server"
)

type Redis struct {
	Client   *redis.Client
	Settings *server.Settings
}

// NewRedis provide the redis Client for the App
func NewRedis(settings *server.Settings) *Redis {
	rdb := redis.NewClient(&redis.Options{
		Addr:     settings.RedisAddr,
		Password: settings.RedisPassword,
		DB:       0,
	})
	return &Redis{
		Client:   rdb,
		Settings: settings,
	}
}

func (r *Redis) AppendUserToMatchPool(ID uint) {
	ctx := context.Background()
	res := r.Client.LPush(ctx, r.Settings.RedisMatchMakingUsersQueueName, ID)
	if res.Err() != nil {
		log.Fatal("error when appending user to matchmaking pool")
		return
	}
}

func (r *Redis) RemoveUserFromMatchPool(ID uint) error {
	ctx := context.Background()
	res := r.Client.LRem(ctx, r.Settings.RedisMatchMakingUsersQueueName, 0, ID)
	if res.Err() != nil {
		return res.Err()
	} else {
		return nil
	}
}

func (r *Redis) GetAllFromQueue() []uint {
	ctx := context.Background()
	UserIDs := r.Client.LRange(ctx, r.Settings.RedisMatchMakingUsersQueueName, 0, -1).Val()
	var UserIDsUint []uint
	for _, id := range UserIDs {
		idInt, _ := strconv.Atoi(id)
		UserIDsUint = append(UserIDsUint, uint(idInt))
	}
	return UserIDsUint
}

func (r *Redis) FlushAll() {
	r.Client.FlushAll(context.Background())
}
