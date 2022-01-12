package model

import "time"

type Oauth2Token struct {
	TokenType    string
	TokenValue   string
	RefreshToken *Oauth2Token
	ExpiresTime  *time.Time
}

func (token *Oauth2Token) IsExpire() bool {
	return token.ExpiresTime != nil && token.ExpiresTime.Before(time.Now())
}

// Oauth2Details oath2详情
type Oauth2Details struct {
	Client *ClientDetails
	User   *UserDetails
}
