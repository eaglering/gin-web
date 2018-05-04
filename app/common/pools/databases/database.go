package databases

import (
	"github.com/go-xorm/xorm"
	_ "github.com/go-sql-driver/mysql"
	. "gin-web/app/config"
	"sync"
	"time"
)

var (
	db *xorm.Engine
	once sync.Once
)

func Instance() *xorm.Engine {
	once.Do(func() {
		var err error
		db, err = xorm.NewEngine(Config.DB.Engine, Config.DB.Dsn)

		db.TZLocation, err = time.LoadLocation("Asia/Shanghai")
		if err != nil {
			panic(err)
		}

		db.SetMaxIdleConns(Config.DB.MaxIdle)
		db.SetMaxOpenConns(Config.DB.MaxOpen)

		if err != nil {
			panic(err)
		}
		if err := db.Ping(); err != nil {
			panic(err)
		}
	})
	return db
}
