package dao

import (
	"context"
	"time"

	"github.com/CloudZou/punk/pkg/gredis"
	"github.com/CloudZou/punk/pkg/conf/paladin"
	"github.com/CloudZou/punk/pkg/sync/pipeline/fanout"
	xtime "github.com/CloudZou/punk/pkg/time"
	"github.com/jinzhu/gorm"

	"github.com/google/wire"
)

var Provider = wire.NewSet(newDao, NewDB, NewRedisClient)

// Dao Dao.
type Dao struct {
	db          *gorm.DB
	redis       *gredis.RedisClient
	cache *fanout.Fanout
	demoExpire int32
}


func newDao(r *gredis.RedisClient, db *gorm.DB) (d *Dao, cf func(), err error) {
	var cfg struct{
		DemoExpire xtime.Duration
	}
	if err = paladin.Get("application.toml").UnmarshalTOML(&cfg); err != nil {
		return
	}
	d = &Dao{
		db: db,
		redis: r,
		cache: fanout.New("cache"),
		demoExpire: int32(time.Duration(cfg.DemoExpire) / time.Second),
	}
	cf = d.Close
	return
}

// Close close the resource.
func (d *Dao) Close() {
	d.cache.Close()
}

// Ping ping the resource.
func (d *Dao) Ping(ctx context.Context) (err error) {
	return nil
}
