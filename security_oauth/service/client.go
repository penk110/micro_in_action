package service

import (
	"context"
	"errors"
	"sync"

	"github.com/penk110/micro_in_action/security_oauth/model"
)

var (
	ErrClientNotExist = errors.New("ClientID is not exist")
	ErrClientSecret   = errors.New("invalid ClientSecret")
)

type Client struct {
	ms    map[string]*model.ClientDetails
	mutex sync.Mutex
}

// NewMemoryClient 构建内存client，以便测试
func NewMemoryClient(clientDetailsList []*model.ClientDetails) *Client {
	var (
		client *Client
	)
	client = &Client{
		ms:    make(map[string]*model.ClientDetails),
		mutex: sync.Mutex{},
	}
	if clientDetailsList != nil {
		for _, value := range clientDetailsList {
			client.ms[value.ClientID] = value
		}
	}

	return client
}

func (c *Client) GetByID(ctx context.Context, clientID string, clientSecret string) (*model.ClientDetails, error) {
	var (
		clientDetails *model.ClientDetails
		ok            bool
	)
	// TODO: 内存查询改成DB查询？
	// ctx -> timeout
	if clientDetails, ok = c.ms[clientID]; ok {
		if clientDetails.ClientSecret != clientSecret {
			return nil, ErrClientSecret
		}
		return clientDetails, nil
	}
	return nil, ErrClientNotExist
}
