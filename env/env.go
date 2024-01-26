package env

import (
	"github.com/caarlos0/env"
	// 自动解析引入境变量
	"log"

	_ "github.com/joho/godotenv/autoload"
)

const TokenExpire = 604800

func init() {
	var err error
	if err = env.Parse(&AliPayConf); err != nil {
		log.Println("parse alipay config error")
	}

	if err = env.Parse(&JwtConf); err != nil {
		log.Println("parse jwt config error")
	}
}
