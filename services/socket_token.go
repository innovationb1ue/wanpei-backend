package services

import (
	"errors"
	"github.com/google/uuid"
	"wanpei-backend/mapper"
)

type Token struct {
	TokenMapper *mapper.Token
}

func NewToken(tokenMapper *mapper.Token) *Token {
	return &Token{TokenMapper: tokenMapper}
}

func (t *Token) IsIn(token string) bool {
	if t.TokenMapper.ToKens[token] != 0 {
		t.TokenMapper.DeleteToken(token)
		return true
	}
	return false
}

func (t *Token) GetUserID(token string) (uint, error) {
	id := t.TokenMapper.GetUserIDByToken(token)
	delete(t.TokenMapper.ToKens, token)
	if id > 0 {
		return id, nil
	} else {
		return 0, errors.New("token not exist")
	}
}

func (t *Token) GenerateRandom(ID uint) string {
	token := uuid.NewString()
	t.TokenMapper.AddToken(token, ID)
	return token
}
