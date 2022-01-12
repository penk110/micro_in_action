package model

type ClientDetails struct {
	// ClientID 客户端ID
	ClientID string
	// ClientSecret 客户端秘钥
	ClientSecret string
	// AccessTokenValiditySeconds 访问令牌有效时间 秒
	AccessTokenValiditySeconds int64
	// RefreshTokenValiditySeconds 刷新令牌有效时间
	RefreshTokenValiditySeconds int64
	// RegisteredRedirectUri 重定向地址，授权码类型中使用
	RegisteredRedirectUri string
	// AuthorizedGrantTypes 可使用的授权类型
	AuthorizedGrantTypes []string // TODO: 配置定义
}
