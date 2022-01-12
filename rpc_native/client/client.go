package main

import (
	"log"
	"net/rpc"

	"github.com/penk110/micro_in_action/rpc_native/server"
)

func main() {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:8030")
	if err != nil {
		log.Fatalf("dial err: %s", err.Error())

	}

	stringReq1 := &server.StringReq{
		Fmt: "a: %s, b: %s",
		A:   "A",
		B:   "B",
	}
	// 同步调用
	var reply1 string
	err = client.Call("StringServer.Fmt", stringReq1, &reply1)
	if err != nil {
		log.Fatalf("req err: %s", err.Error())
	}
	log.Printf("StringService Fmt: %s A: %s B: %s ret: %s\n", stringReq1.Fmt, stringReq1.A, stringReq1.B, reply1)

	var reply2 string
	stringReq2 := &server.StringReq{
		Fmt: "aa: %s, bb: %s",
		A:   "AA",
		B:   "BB",
	}
	// 异步调用
	call := client.Go("StringServer.Fmt", stringReq2, &reply2, nil)
	done := <-call.Done
	log.Printf("StringService2 Fmt: %s A: %s B: %s ret: %s done: %v\n", stringReq2.Fmt, stringReq2.A, stringReq2.B, reply2, done)
}
