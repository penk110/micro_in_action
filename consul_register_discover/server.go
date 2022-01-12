package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	uuid "github.com/satori/go.uuid"

	"github.com/penk110/micro_in_action/consul_register_discover/config"
	"github.com/penk110/micro_in_action/consul_register_discover/discover"
	"github.com/penk110/micro_in_action/consul_register_discover/endpoint"
	"github.com/penk110/micro_in_action/consul_register_discover/service"
	"github.com/penk110/micro_in_action/consul_register_discover/transport"
)

func init() {
	config.InitEnv()
}

func main() {
	var (
		svc             service.Service
		discoveryClient discover.DiscoveryClient
		endpoints       endpoint.DiscoveryEndpoints
		ctx             context.Context
		errChan         chan error
		err             error
	)

	ctx = context.Background()
	errChan = make(chan error)
	if discoveryClient, err = discover.NewKitDiscover(config.Cfg.ServerCfg.Addr, config.Cfg.ConsulCfg.Port); err != nil {
		config.Logger.Printf("get consul client failed, err: %s", err.Error())
		os.Exit(1)
	}
	svc = service.NewDiscoveryService(discoveryClient)
	endpoints = endpoint.DiscoveryEndpoints{
		PingPongEndpoint:    endpoint.MakePingPongEndpoint(svc),
		DiscoveryEndpoint:   endpoint.MakeDiscoveryEndpoint(svc),
		HealthCheckEndpoint: endpoint.MakeHealthCheckEndpoint(svc),
	}

	// TODO: 定义HTTP接口
	log.Println(endpoints)
	router := transport.MakeHttpHandler(ctx, endpoints, config.KitLogger)
	instanceId := config.Cfg.ServerCfg.Name + "-" + uuid.NewV4().String()

	// TODO: 注册、监听、取消注册
	go func() {
		config.Logger.Println("http server start at port:" + strconv.Itoa(config.Cfg.ServerCfg.Port))
		if !discoveryClient.Register(config.Cfg.ServerCfg.Name, instanceId, "/health", config.Cfg.ServerCfg.Addr, config.Cfg.ServerCfg.Port, nil, config.Logger) {
			config.Logger.Printf("string-service for service %s failed.", config.Cfg.ServerCfg.Name)
			os.Exit(1)
		}
		handler := router
		errChan <- http.ListenAndServe(config.Cfg.ServerCfg.Addr+":"+strconv.Itoa(config.Cfg.ServerCfg.Port), handler)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	err = <-errChan
	//服务退出取消注册
	discoveryClient.DeRegister(instanceId, config.Logger)
	log.Println(err)
}
