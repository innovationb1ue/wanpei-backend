package test

import (
	"context"
	"github.com/go-redis/redis/v9"
	"github.com/stretchr/testify/assert"
	"log"
	"strconv"
	"testing"
	"wanpei-backend/server"
)

func TestRedisConnSet(t *testing.T) {
	env := server.GetEnv()
	settings := server.NewSettings(env)
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

func TestLPos(t *testing.T) {
	env := server.GetEnv()
	settings := server.NewSettings(env)
	rdb := redis.NewClient(&redis.Options{
		Addr:     settings.RedisAddr,
		Password: settings.RedisPassword,
		DB:       0,
	})
	ctx := context.Background()
	rdb.LPush(ctx, settings.RedisMatchMakingUsersQueueName, 999)
	rdb.LPush(ctx, settings.RedisMatchMakingUsersQueueName, 1)
	rdb.LPush(ctx, settings.RedisMatchMakingUsersQueueName, 2)
	res := rdb.LPos(ctx, settings.RedisMatchMakingUsersQueueName, strconv.Itoa(999), redis.LPosArgs{
		Rank:   1,
		MaxLen: 0,
	})

	t.Log(res.Val())
	assert.NotNil(t, res.Val())

	UserIds := rdb.LRange(ctx, settings.RedisMatchMakingUsersQueueName, 0, -1).Val()
	assert.Contains(t, UserIds, "999")
	t.Log(UserIds)
	rdb.FlushAll(ctx)
}
