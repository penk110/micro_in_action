package service

type Health struct {
}

func NewHealth() *Health {
	return &Health{}
}

func (h *Health) Check() bool {
	return true
}

func (h *Health) PingPong() string {
	return "pong"
}
