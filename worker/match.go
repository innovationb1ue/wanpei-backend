package worker

import (
	"context"
	"strconv"
	"time"
	"wanpei-backend/mapper"
	"wanpei-backend/models"
	"wanpei-backend/repo"
)

type Match struct {
	UserMapper     *mapper.User
	UserGameMapper *mapper.UserGame
	SocketMgr      *mapper.Socket
	RedisMapper    *mapper.Redis
	HubMapper      *mapper.Hub
	QueueUserGame  *repo.QueueUserGame
}

func NewMatch(UserMapper *mapper.User, UserGameMapper *mapper.UserGame,
	SocketMgr *mapper.Socket, RedisMapper *mapper.Redis, hubMapper *mapper.Hub, queueUserGame *repo.QueueUserGame) *Match {
	return &Match{
		UserMapper:     UserMapper,
		UserGameMapper: UserGameMapper,
		SocketMgr:      SocketMgr,
		RedisMapper:    RedisMapper,
		HubMapper:      hubMapper,
		QueueUserGame:  queueUserGame,
	}
}

func MatchWorker(match *Match) {
	go match.MakeMatch(context.Background())
}

// MakeMatch query for all the users in the queue and make pairs at a fixed interval
// should be started once at the App level
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
			//UserGameTags := m.UserGameMapper.GetUserGames(idUint) // this get user game from database
			userGameIDs := m.QueueUserGame.UserGame[idUint]
			// iter through tags of a single user
			for _, gameID := range userGameIDs {
				if allTags[gameID] != 0 {
					go m.matchSuccess(ctx, idUint, allTags[gameID])
					delete(allTags, gameID)
				} else {
					allTags[gameID] = idUint
				}
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

	// make a room for users & run broadcast routine
	hub := models.NewHub()
	go hub.Run()
	m.HubMapper.RegisterNewHub(hub)
	// todo: handle hub destroy

	// send success message back to client
	err = socket1.WriteJSON(map[string]any{"action": "success", "data": map[string]string{"ID": hub.ID}})
	if err != nil {
		return
	}
	err = socket2.WriteJSON(map[string]any{"action": "success", "data": map[string]string{"ID": hub.ID}})
	if err != nil {
		return
	}

	// move them our ot the match Pool
	m.RedisMapper.RemoveUserFromMatchPool(ID1)
	m.RedisMapper.RemoveUserFromMatchPool(ID2)

	// close sockets without handling any further error.
	// Users got matched so there is no need to preserve this result-informing websocket connections anymore.
	_ = socket1.Close()
	_ = socket2.Close()
}
