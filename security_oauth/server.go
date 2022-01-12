package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	uuid "github.com/satori/go.uuid"

	"github.com/penk110/micro_in_action/common/discover"
	"github.com/penk110/micro_in_action/security_oauth/config"
	"github.com/penk110/micro_in_action/security_oauth/klog"
	"github.com/penk110/micro_in_action/security_oauth/transport"
)

func main() {
	var (
		instanceID string
		dcClient   discover.Client
		err        error
		errCI      error
		errCh      chan error
	)
	instanceID = config.Server.Name + "-" + uuid.NewV4().String()
	errCh = make(chan error)
	if dcClient, err = discover.NewKitDiscover(config.Consul.Host, config.Consul.Port); err != nil {
		klog.Logger.Printf(err.Error())
		os.Exit(-1)
	}

	//创建http.Handler32
	//r := transport.MakeHttpHandler(ctx, endpoints, tokenService, clientDetailsService, config.KitLogger)

	//http server
	go func() {
		klog.Logger.Printf("Http Server start at %s:%d\n", config.Server.Host, config.Server.Port)
		if !dcClient.Register(config.Server.Name, instanceID, "/health", config.Server.Host, config.Server.Port, nil, klog.Logger) {
			klog.Logger.Printf("use-string-service for service %s failed", config.Server.Host)
			os.Exit(-1)
		}
		errCh <- http.ListenAndServe(fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port), transport.GetMuxRouter())
	}()

	go func() {
		var signalCh chan os.Signal
		signalCh = make(chan os.Signal)
		signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)
		errCh <- fmt.Errorf("%s", <-signalCh)
	}()

	errCI = <-errCh
	dcClient.DeRegister(instanceID, klog.Logger)
	klog.Logger.Println(errCI)
}
