package worker

import (
	"context"
	"strconv"
	"time"
	"wanpei-backend/mapper"
)

type Match struct {
	UserMapper     *mapper.User
	UserGameMapper *mapper.UserGame
	SocketMgr      *mapper.Socket
	RedisMapper    *mapper.Redis
}

func NewMatch(UserMapper *mapper.User, UserGameMapper *mapper.UserGame, SocketMgr *mapper.Socket, RedisMapper *mapper.Redis) *Match {
	return &Match{
		UserMapper:     UserMapper,
		UserGameMapper: UserGameMapper,
		SocketMgr:      SocketMgr,
		RedisMapper:    RedisMapper,
	}
}

func MatchWorker(match *Match) {
	go match.MakeMatch(context.Background())
}

// MakeMatch query for all the users in the queue and make pairs at a fixed interval
func (m *Match) MakeMatch(ctx context.Context) {
	ticker := time.NewTicker(3 * time.Second)
	for {
		<-ticker.C
		// get all users in Redis
		userIDs := m.RedisMapper.Client.LRange(ctx, "match:users", 0, -1).Val()
		if len(userIDs) == 0 {
			continue
		}
		// use a map to hold UserID:TagString
		// when inserting all tags into the map, make pairs whenever duplicate tag is met.
		allTags := map[int]uint{}
		// iter through all users in queue
		for _, id := range userIDs {
			idInt, _ := strconv.Atoi(id)
			idUint := uint(idInt)
			UserGameTags := m.UserGameMapper.GetUserGames(idUint)
			// iter through tags of a single user
			for _, userGame := range UserGameTags {
				if allTags[userGame.GameID] != 0 {
					go m.matchSuccess(ctx, idUint, allTags[userGame.GameID])
					delete(allTags, userGame.GameID)
				} else {
					allTags[userGame.GameID] = userGame.ID
				}
			}
		}
	}
}

func (m *Match) matchSuccess(ctx context.Context, ID1 uint, ID2 uint) {
	//todo: finish the logic here after reading best practices.
	// need to think more about the concurrency and lifecycle problem.
	// What if a client quit when we iter through the queue list and got matched with another?

	// todo: delete userID from queue
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
