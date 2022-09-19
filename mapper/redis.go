package mapper

import (
	"context"
	"github.com/go-redis/redis/v9"
	"strconv"
	"wanpei-backend/models"
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
	return &Redis{Client: rdb}
}

func (r *Redis) AppendUserToMatchPool(user *models.User) {
	ctx := context.Background()
	r.Client.LPush(ctx, r.Settings.RedisMatchMakingUsersQueueName, user.ID)
}

func (r *Redis) CheckUserInMatchPool(user *models.User) bool {
	ctx := context.Background()
	res := r.Client.LPos(ctx, r.Settings.RedisMatchMakingUsersQueueName, strconv.Itoa(int(user.ID)), redis.LPosArgs{
		Rank:   1,
		MaxLen: 0,
	})
	return res.Val() == 0
}

func (r *Redis) RemoveUserFromMatchPool(user *models.User) {
	ctx := context.Background()
	r.Client.LRem(ctx, r.Settings.RedisMatchMakingUsersQueueName, 0, user.ID)
}
