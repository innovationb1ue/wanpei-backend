package worker

import (
	"context"
	"github.com/go-redis/redis/v9"
	"go.uber.org/fx"
	"strconv"
	"wanpei-backend/mapper"
)

type Match struct {
	fx.In
	UserMapper     *mapper.User
	UserGameMapper *mapper.UserGame
	SocketMgr      *mapper.Socket
}

// MakeMatch query for all the users in the queue and make pairs at a fixed interval
func (m *Match) MakeMatch(ctx context.Context, rdb *redis.Client) {
	// get all users in Redis
	userIDs := rdb.LRange(ctx, "match:users", 0, -1).Val()
	// use a map to hold UserID:TagString
	// when inserting all tags into the map, make pairs whenever duplicate tag is met.
	allTags := map[string]uint{}
	// iter through all users in queue
	for _, id := range userIDs {
		idInt, _ := strconv.Atoi(id)
		idUint := uint(idInt)
		UserGameTags := m.UserGameMapper.GetUserGames(idUint)
		// iter through tags of a single user
		for _, tag := range UserGameTags {
			if allTags[tag] != 0 {
				go m.matchSuccess(ctx, idUint, allTags[tag])
				delete(allTags, tag)
			}
		}
	}
}

func (m *Match) matchSuccess(ctx context.Context, ID1 uint, ID2 uint) {
	//todo: finish the logic here after reading best practices.
	// need to think more about the concurrency and lifecycle problem.
	// What if a client quit when we iter through the queue list and got matched with another?

	// first check whether socket is still alive by having them
	socket1, err := m.SocketMgr.GetSocket(ID1)
	if err != nil {
		return
	}
	socket2, err := m.SocketMgr.GetSocket(ID2)
	if err != nil {
		return
	}
	// send success message back to client
	err = socket1.WriteJSON(map[string]any{"action": "success", "data": nil})
	if err != nil {
		return
	}
	err = socket2.WriteJSON(map[string]any{"action": "success", "data": nil})
	if err != nil {
		return
	}
}
