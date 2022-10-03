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
		// timed task: fetch users in queue and match them
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
					// simultaneously handle matched users
					go m.matchSuccess(context.Background(), idUint, allTags[gameID])
					delete(allTags, gameID) // delete key in the game count map
					break                   // break, no need to check the rest of the tags
				} else {
					allTags[gameID] = idUint
				}
			}
		}
	}
}

func (m *Match) matchSuccess(ctx context.Context, ID1 uint, ID2 uint) {
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
	hub.AppendAvailableUser(ID1)
	hub.AppendAvailableUser(ID2)
	hubCtx, cancel := context.WithCancel(ctx)
	// start the hub broadcasting thread
	go hub.Run(cancel) // will cancel the whole context if the hub is deprecated when no one in it
	go func() {
		// register self to hub repo
		m.HubMapper.RegisterNewHub(hub)
		// block until the hub is done.
		<-hubCtx.Done()
		// unregister self to hub
		m.HubMapper.DeleteHub(hub.ID)
	}()

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
	_ = m.RedisMapper.RemoveUserFromMatchPool(ID1)
	_ = m.RedisMapper.RemoveUserFromMatchPool(ID2)

	// close sockets without handling any further error.
	// Users got matched so there is no need to preserve this result-informing websocket connections anymore.
	_ = socket1.Close()
	_ = socket2.Close()
}
