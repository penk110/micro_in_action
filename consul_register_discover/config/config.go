package config

import (
	"encoding/json"
	"log"
	"os"
	"strconv"

	kitLog "github.com/go-kit/kit/log"
)

/*
	TODO: 先从环境变量读取
*/

var (
	Cfg       *Config
	Logger    *log.Logger
	KitLogger kitLog.Logger
)

func init() {
	Logger = log.New(os.Stderr, "", log.LstdFlags)
	KitLogger = kitLog.NewLogfmtLogger(os.Stderr)
	KitLogger = kitLog.With(KitLogger, "ts", kitLog.DefaultTimestampUTC)
	KitLogger = kitLog.With(KitLogger, "caller", kitLog.DefaultCaller)
}

func InitEnv() {
	var (
		envFrom string
		err     error
	)
	envFrom = os.Getenv("ENV_FROM")
	switch envFrom {
	case "CONSUL":
		panic("not support")
	case "ENV":
		err = initFromEnv()
	default:
		panic("envFrom is empty")
	}

	if err != nil {
		log.Panicf("init cfg failed, envFrom: %s err: %s\n", envFrom, err.Error())
	}
	s, _ := json.Marshal(Cfg)
	log.Println("cfg: "+string(s), Cfg.ServerCfg.Addr, Cfg.ServerCfg.Port)
}

type Config struct {
	ServerCfg *ServerCfg `json:"server_cfg"`
	ConsulCfg *ConsulCfg `json:"consul_cfg"`
}

type ServerCfg struct {
	Name string `json:"name"`
	Addr string `json:"addr"`
	Port int    `json:"port"`
}

type ConsulCfg struct {
	Addr string `json:"addr"`
	Port uint64 `json:"port"`
}

func initFromEnv() error {
	var (
		serverCfg *ServerCfg
		consulCfg *ConsulCfg
		err       error
	)
	serverCfg = &ServerCfg{}
	serverCfg.Name = os.Getenv("SERVER_NAME")
	serverCfg.Addr = os.Getenv("SERVER_ADDR")
	if serverCfg.Port, err = strconv.Atoi(os.Getenv("SERVER_PORT")); err != nil {
		return err
	}

	consulCfg = &ConsulCfg{}
	consulCfg.Addr = os.Getenv("CONSUL_ADDR")
	if consulCfg.Port, err = strconv.ParseUint(os.Getenv("CONSUL_PORT"), 10, 64); err != nil {
		return err
	}
	Cfg = &Config{
		ServerCfg: serverCfg,
		ConsulCfg: consulCfg,
	}

	return nil
}

func GetLogger() *log.Logger {
	return Logger
}

func GetKitLogger() kitLog.Logger {
	return KitLogger
}
