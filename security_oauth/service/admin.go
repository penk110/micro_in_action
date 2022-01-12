package service

type Admin struct {
}

func NewAdmin() *Admin {
	return &Admin{}
}

func (a *Admin) SimpleData(username string) string {
	return "hello " + username + " ,simple data, with simple authority"
}
