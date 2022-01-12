package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/go-kit/kit/log"
	"github.com/hashicorp/consul/api"
)

const (
	consulHost = "127.0.0.1"
	consulPort = "8500"
)

func main() {
	var (
		errCh  chan error
		logger log.Logger
	)
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)

	// 创建consul api客户端
	consulConfig := api.DefaultConfig()
	consulConfig.Address = "http://" + consulHost + ":" + consulPort
	consulClient, err := api.NewClient(consulConfig)
	if err != nil {
		logger.Log("err", err)
		os.Exit(1)
	}

	proxy := NewReverseProxy(consulClient, logger)

	errCh = make(chan error)

	go func() {
		var (
			signalCh = make(chan os.Signal)
		)

		signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)
		errCh <- fmt.Errorf("%s", <-signalCh)
	}()

	//开始监听
	go func() {
		fmt.Println("transport", "HTTP", "addr", "8030")
		errCh <- http.ListenAndServe(":8030", proxy)
	}()
	logger.Log("exit", <-errCh)
}

func NewReverseProxy(client *api.Client, logger log.Logger) *httputil.ReverseProxy {
	// Director must be a function which modifies
	// the request into a new request to be sent
	// using Transport. Its response is then copied
	// back to the original client unmodified.
	// Director must not access the provided Request
	// after returning.
	// Director func(*http.Request)
	Director := func(req *http.Request) {
		var (
			reqPath     string
			pathArray   []string
			serviceName string
		)
		reqPath = req.URL.Path
		if reqPath == "" {
			return
		}

		// path: /serverName/path
		pathArray = strings.Split(reqPath, "/")
		serviceName = pathArray[1]
		fmt.Printf("[NewReverseProxy] reqPath: %s pathArray: %v serviceName: %s\n", reqPath, pathArray, pathArray[1])

		// 调用consul api查询serviceName的服务实例列表
		result, _, err := client.Catalog().Service(serviceName, "", nil)
		if err != nil {
			fmt.Printf("[NewReverseProxy] ReverseProxy failed, err: %s\n", err.Error())
			return
		}

		//重新组织请求路径，去掉服务名称部分
		destPath := strings.Join(pathArray[2:], "/")

		// 随机选择一个服务实例
		if len(result) == 0 {
			return
		}
		tgt := result[rand.Int()%len(result)]
		fmt.Printf("[NewReverseProxy] service id:  %s\n", tgt.ServiceID)

		// 设置代理服务地址信息
		req.URL.Scheme = "http"
		req.URL.Host = fmt.Sprintf("%s:%d", tgt.ServiceAddress, tgt.ServicePort)
		req.URL.Path = "/" + destPath

	}

	return &httputil.ReverseProxy{Director: Director}
}
