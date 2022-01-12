package config

var (
	MySQL  = &mySQL{}
	Server = &server{}
)

type server struct {
	Env     string
	Address string
	Port    int
}

type mySQL struct {
	DBName   string
	Password string
	Addr     string
	Port     int
	ChartSet string
}
