package model

type UserDetails struct {
	// UserId 用户ID
	UserId int64
	// Username 用户名 唯一
	Username string
	// Password 用户密码
	Password string
	// Authorities 用户具有的权限
	Authorities []string // 具备的权限
}
