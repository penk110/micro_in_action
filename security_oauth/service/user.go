package service

import (
	"context"
	"errors"
	"sync"

	"github.com/penk110/micro_in_action/security_oauth/model"
)

var (
	ErrUserNotExist = errors.New("user is not exist")
	ErrPassword     = errors.New("invalid password")
)

type User struct {
	users map[string]*model.UserDetails
	mutex sync.Mutex
}

func NewMemoryUser() *User {
	var (
		user *User
	)
	user = &User{
		users: make(map[string]*model.UserDetails),
		mutex: sync.Mutex{},
	}
	return user
}

func (u *User) GetUserDetailByUsername(ctx context.Context, username, password string) (*model.UserDetails, error) {
	var (
		user *model.UserDetails
		ok   bool
	)
	if user, ok = u.users[username]; ok {
		if user.Password == password {
			return user, nil
		} else {
			return nil, ErrPassword
		}
	}
	return nil, ErrUserNotExist
}
