package test

import (
	"context"
	"github.com/go-redis/redis/v9"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"wanpei-backend/server"
)

func TestRedisConnSet(t *testing.T) {
	settings := server.NewSettings()
	rdb := redis.NewClient(&redis.Options{
		Addr:     settings.RedisAddr,
		Password: settings.RedisPassword,
		DB:       0,
	})
	ctx := context.Background()
	rdb.SAdd(ctx, "user:b1ue:game", 123)
	rdb.SAdd(ctx, "user:b1ue:game", 456)

	res1 := rdb.SMembers(ctx, "user:b1ue:game")
	cmd := rdb.SIsMember(ctx, "user:b1ue:game", 123)
	assert.True(t, cmd.Val())

	cmd = rdb.SIsMember(ctx, "user:b1ue:game", 789)
	assert.False(t, cmd.Val())

	log.Println(res1)

	rdb.FlushAll(ctx)

	rdb.LPush(ctx, "match:users", 1)
	matchUsers := rdb.LRange(ctx, "match:users", 0, -1)

	log.Println(matchUsers)

	rdb.FlushAll(ctx)
}
