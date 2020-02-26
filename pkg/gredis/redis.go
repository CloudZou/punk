package gredis

import (
	"encoding/json"
	"github.com/gomodule/redigo/redis"
	"time"
)

type RedisClient struct {
	RedisPool *redis.Pool
}

type RedisConfig struct {
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

func NewRedisConn(cfg RedisConfig) (redisConn *redis.Pool, err error) {
	redisConn = &redis.Pool{
		MaxIdle:     cfg.MaxIdle,
		MaxActive:   cfg.MaxActive,
		IdleTimeout: cfg.IdleTimeout,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", cfg.Host)
			if err != nil {
				return nil, err
			}
			if cfg.Password != "" {
				if _, err := c.Do("AUTH", cfg.Password); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}

	return
}

// Set a key/value
func (r *RedisClient) Set(key string, data interface{}, time int) error {
	conn := r.RedisPool.Get()
	defer conn.Close()

	value, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = conn.Do("SET", key, value)
	if err != nil {
		return err
	}

	_, err = conn.Do("EXPIRE", key, time)
	if err != nil {
		return err
	}

	return nil
}

// Exists check a key
func (r *RedisClient) Exists(key string) bool {
	conn := r.RedisPool.Get()
	defer conn.Close()

	exists, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return false
	}

	return exists
}

// Get get a key
func (r *RedisClient) Get(key string) ([]byte, error) {
	conn := r.RedisPool.Get()
	defer conn.Close()

	reply, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return nil, err
	}

	return reply, nil
}

// Delete delete a kye
func (r *RedisClient) Delete(key string) (bool, error) {
	conn := r.RedisPool.Get()
	defer conn.Close()

	return redis.Bool(conn.Do("DEL", key))
}
