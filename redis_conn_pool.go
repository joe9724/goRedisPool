// Pool maintains a pool of connections. The application calls the Get method to get a connection from the pool
// and the connection's Close method to return the connection's resources to the pool.

package redisPool

import (
	"time"
	"log"

	"github.com/garyburd/redigo/redis"
)

func newPool(config *RedisConfig) *redis.Pool {
	return &redis.Pool{
		MaxIdle: config.RedisPoolMaxIdle,
		MaxActive: config.RedisPoolMaxActive,
		IdleTimeout: time.Duration(config.RedisPoolIdleTimeout) * time.Second,
		Dial: func () (redis.Conn, error) { return redis.Dial("tcp", config.RedisAddress) },
	}
}

var (
	pool *redis.Pool
)

func InitRedisPool(config *RedisConfig) {
	log.Printf("Init Redis Connection pool with params: RedisAddress = %s, MaxIdle = %d, MaxActive = %d, IdleTimeout = %d(s)",
		config.RedisAddress, config.RedisPoolMaxIdle, config.RedisPoolMaxActive, config.RedisPoolIdleTimeout)
	pool = newPool(config)
}
