package redisPool

type RedisConfig struct {
	RedisAddress		string		`yaml:"redis_address"`
	RedisPoolMaxIdle	int		`yaml:"redis_pool_max_idle"`
	RedisPoolMaxActive	int		`yaml:"redis_pool_max_active"`
	RedisPoolIdleTimeout	int		`yaml:"redis_pool_idle_timeout"`
}
