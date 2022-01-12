package endpoint

import (
	"github.com/go-kit/kit/endpoint"
	"net/http"
)

type Oauth2Endpoints struct {
	TokenEndpoint       endpoint.Endpoint
	CheckTokenEndpoint  endpoint.Endpoint
}

// 构造方法

// 定义

type TokenReq struct {
	GrantType string
	Reader    *http.Request
}

type TokenResp struct {

}
