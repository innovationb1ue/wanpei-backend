package mapper

import "time"

type Token struct {
	ToKens map[string]uint
}

func NewTokenManager() *Token {
	return &Token{ToKens: make(map[string]uint)}
}

// AddToken handle the full lifecycle of a websocket token.
func (t *Token) AddToken(token string, ID uint) {
	t.ToKens[token] = ID
	// set a timer to delete this key after certain duration
	go func() {
		interval := time.NewTimer(100 * time.Second)
		<-interval.C
		delete(t.ToKens, token)
	}()
}

func (t *Token) DeleteToken(token string) {
	delete(t.ToKens, token)
}

func (t *Token) GetUserIDByToken(token string) uint {
	id := t.ToKens[token]
	return id
}
