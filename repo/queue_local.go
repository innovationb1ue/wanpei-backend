package repo

type QueueUserGame struct {
	UserGame map[uint][]int // store the selected game for each user in queue.
}

func NewQueueUserGame() *QueueUserGame {
	return &QueueUserGame{UserGame: map[uint][]int{}}
}
