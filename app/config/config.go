package config

import (
	"time"
	"gin-web/app/common/helpers"
)

var (
	Config		config
	configFile = "app/config/config.yaml"
)

type config struct {
	ReleaseMode bool   `yaml:"release_mode"`
	LogLevel    string `yaml:"log_level"`

	SessionStore string `yaml:"session_store"`
	CacheStore   string `yaml:"cache_store"`

	// 应用配置
	App app

	// 模板
	Tmpl tmpl

	Server server

	// MySQL
	DB database `yaml:"database"`

	// 静态资源
	Static static

	// Redis
	Redis redis

	// Cache
	Cache cache

	// WebSocket
	WebSocket websocket
}

type app struct {
	Name string `yaml:"name"`
}

type server struct {
	Graceful 	bool   `yaml:"graceful"`
	Address     string `yaml:"address"`
	ApiSecret 	string `yaml:"api_secret"`
}

type static struct {
	Type string `yaml:"type"`
}

type tmpl struct {
	Type   string `yaml:"type"`   // PONGO2,TEMPLATE(TEMPLATE Default)
	Data   string `yaml:"data"`   // BINDATA,FILE(FILE Default)
	Dir    string `yaml:"dir"`    // PONGO2(template/pongo2),TEMPLATE(template)
	Suffix string `yaml:"suffix"` // .html,.tpl
}

type database struct {
	Engine   string `yaml:"engine"`
	Dsn		 string `yaml:"dsn"`
	MaxIdle  int	`yaml:"max_idle"`
	MaxOpen  int	`yaml:"max_open"`
}

type cache struct {
	Address string `yaml:"address"`
	Password    string `yaml:"password"`
	Expiration time.Duration	`yaml:"expiration"`
}

type redis struct {
	Address string `yaml:"address"`
	Password    string `yaml:"password"`
	MaxIdle int    `yaml:"max_idle"`
}

type websocket struct {
	PongTimeout int	`yaml:"pong_timeout"`
}

func InitConfig() error {
	Config = config{
		ReleaseMode: false,
		LogLevel:    "debug",
	}

	return helpers.YAML(configFile, &Config)
}

const (
	// Template Type
	PONGO2   = "PONGO2"
	TEMPLATE = "TEMPLATE"

	// Bindata
	BINDATA = "BINDATA"

	// File
	FILE = "FILE"

	// Redis
	REDIS = "REDIS"

	// memcached
	MEMCACHED = "MEMCACHED"

	// Cookie
	COOKIE = "COOKIE"

	// In Memory
	IN_MEMORY = "IN_MEMARY"
)