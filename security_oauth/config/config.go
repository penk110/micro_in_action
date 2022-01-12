package config

var (
	Server *server
	Consul *consul
)

func init() {
	Server = &server{
		Name: "SECURITY_OAUTH",
		Host: "127.0.0.1",
		Port: 8000,
	}
	Consul = &consul{
		Host: "127.0.0.1",
		Port: 8500,
	}
}

type server struct {
	Name string
	Host string
	Port uint64
}

type consul struct {
	Host string
	Port uint64
}
