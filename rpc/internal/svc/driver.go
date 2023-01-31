package svc

import (
	"database/sql"
	"time"

	"ariga.io/entcache"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/go-redis/redis/v8"
	"github.com/hewenyu/gozero-ent/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/logx"

	_ "github.com/go-sql-driver/mysql"
)

func GetCacheDriver(c *config.Config) *entcache.Driver {
	db, err := sql.Open("mysql", c.Mysql.DataSource)
	logx.Must(err)

	// db.SetMaxOpenConns(100)
	driver := entsql.OpenDB("mysql", db)

	rdb := redis.NewClient(&redis.Options{
		Addr:     c.Redis.Host,
		Password: c.Redis.Pass,
		DB:       0,
	})

	cacheDrv := entcache.NewDriver(
		driver,
		entcache.TTL(time.Second),
		entcache.Levels(
			entcache.NewLRU(256),
			entcache.NewRedis(rdb),
		),
	)

	return cacheDrv
}

func GetCacheTTLDriver(c *config.Config) *entcache.Driver {
	db, err := sql.Open("mysql", c.Mysql.DataSource)
	logx.Must(err)

	db.SetMaxOpenConns(100)

	driver := entsql.OpenDB("mysql", db)

	// 也可以使用redis 作为公共的缓存
	// rdb := redis.NewClient(&redis.Options{
	// 	Addr:     c.Redis.Host,
	// 	Password: c.Redis.Pass,
	// 	DB:       0,
	// })

	cacheDrv := entcache.NewDriver(
		driver,
		entcache.TTL(time.Second),
	)

	return cacheDrv
}

func GetDriver(c *config.Config) *entsql.Driver {
	db, err := sql.Open("mysql", c.Mysql.DataSource)
	db.SetMaxOpenConns(100)
	logx.Must(err)

	return entsql.OpenDB("mysql", db)
}
