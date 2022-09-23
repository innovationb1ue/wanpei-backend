package repo

type QueueUserGame struct {
	UserGame map[uint][]int
}

func NewQueueUserGame() *QueueUserGame {
	return &QueueUserGame{UserGame: map[uint][]int{}}
}
