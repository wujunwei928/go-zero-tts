package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf

	// 自定义配置
	// cors 跨域origin配置列表
	CorsOriginList []string `json:"corsOriginList"`
}
