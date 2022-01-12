package main

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/spf13/viper"

	"github.com/penk110/micro_in_action/config_viper/config"
)

var (
	BasePath string
)

func init() {
	var (
		err error
	)
	if BasePath, err = os.Getwd(); err != nil {
		log.Printf("get work dir failed, err: %s\n", err.Error())
		return
	}

	viper.AutomaticEnv()
	viper.SetConfigName("config")
	// 读取的配置文件路径
	viper.AddConfigPath(path.Join(BasePath, "config/"))
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("ReadInConfig failed, err:%s\n", err.Error())
	}
	if err := subConfig("MySQL", &config.MySQL); err != nil {
		log.Fatal("Fail to parse config", err)
	}
}

func subConfig(key string, value interface{}) error {
	log.Printf("config prefix：%v\n", key)
	sub := viper.Sub(key)
	sub.AutomaticEnv()
	sub.SetEnvPrefix(key)
	return sub.Unmarshal(value)
}

func parseYaml(v *viper.Viper) {

}

func main() {
	fmt.Println(config.MySQL)
}
