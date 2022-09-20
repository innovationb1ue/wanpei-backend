package worker

import (
	"context"
	"github.com/go-redis/redis/v9"
	"go.uber.org/fx"
	"log"
	"strconv"
	"wanpei-backend/mapper"
	"wanpei-backend/models"
)

type Match struct {
	fx.In
	UserMapper *mapper.User
}

// MakeMatch query for all the users in the queue and make pairs at a fixed interval
func (m *Match) MakeMatch(rdb *redis.Client) {
	ctx := context.Background()
	var users []*models.User
	// get all users
	for _, id := range rdb.LRange(ctx, "match:users", 0, -1).Val() {
		idInt, err := strconv.ParseInt(id, 10, 4)
		if err != nil {
			log.Fatal("can't parse id string", idInt, "to int")
		}
		user, err := m.UserMapper.GetUserById(ctx, uint(idInt))
		users = append(users, user)
	}
	// todo: iter through users to match

}
