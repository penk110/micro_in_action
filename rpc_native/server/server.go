package server

import (
	"fmt"
	"log"
)

type Service interface {
	Fmt(req StringReq, ret *string) error
}

type StringReq struct {
	Fmt string
	A   string
	B   string
}

type StringServer struct {
}

func (ss StringServer) Fmt(s StringReq, ret *string) error {
	*ret = fmt.Sprintf(s.Fmt, s.A, s.B)
	log.Printf("fmt: %s\n", *ret)
	return nil
}
