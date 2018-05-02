package config

import (
	"time"
	"gin-web/app/common/helpers"
)

var (
	Config		config
	configFile = "app/config/config.toml"
)

type config struct {
	ReleaseMode bool   `toml:"release_mode"`
	LogLevel    string `toml:"log_level"`

	SessionStore string `toml:"session_store"`
	CacheStore   string `toml:"cache_store"`

	// 应用配置
	App app

	// 模板
	Tmpl tmpl

	Server server

	// MySQL
	DB database `toml:"database"`

	// 静态资源
	Static static

	// Redis
	Redis redis

	// Cache
	Cache cache
}

type app struct {
	Name string `toml:"name"`
}

type server struct {
	Graceful bool   `toml:"graceful"`
	Address     string `toml:"address"`

	DomainApi    string `toml:"domain_api"`

	ApiSecret string `toml:"api_secret"`
}

type static struct {
	Type string `toml:"type"`
}

type tmpl struct {
	Type   string `toml:"type"`   // PONGO2,TEMPLATE(TEMPLATE Default)
	Data   string `toml:"data"`   // BINDATA,FILE(FILE Default)
	Dir    string `toml:"dir"`    // PONGO2(template/pongo2),TEMPLATE(template)
	Suffix string `toml:"suffix"` // .html,.tpl
}

type database struct {
	Engine   string `toml:"engine"`
	Dsn		 string `toml:"dsn"`
	MaxIdle  int	`toml:"max_idle"`
	MaxOpen  int	`toml:"max_open"`
}

type cache struct {
	Address string `toml:"address"`
	Password    string `toml:"password"`
	Expiration time.Duration	`toml:"expiration"`
}

type redis struct {
	Address string `toml:"address"`
	Password    string `toml:"password"`
	MaxIdle int    `toml:"max_idle"`
}

func InitConfig() error {
	Config = config{
		ReleaseMode: false,
		LogLevel:    "debug",
	}

	return helper.YAML(configFile, &Config)
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