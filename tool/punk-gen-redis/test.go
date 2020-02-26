package main

import "context"

type _redis interface {
	// redis: -key=keyArt -type=get
	CacheUser(c context.Context, id int64) (*Demo, error)
	// redis: -key=keyArt -expire=d.demoExpire
	AddCacheUser(c context.Context, id int64, art *Demo) (err error)
	// redis: -key=keyArt
	DeleteUser(c context.Context, id int64) (err error)
}
