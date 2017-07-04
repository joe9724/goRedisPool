// Methods for writing operations are not exported, only use for test.

package redisPool

import (
	"log"

	"github.com/garyburd/redigo/redis"
)

type Redis struct {
	conn      redis.Conn
}

func NewRedis() *Redis {
	conn := pool.Get()
	log.Printf("Get conn from redis_conn_pool, active connections in the pool is %d", pool.ActiveCount())
	return &Redis{
		conn: conn,
	}
}

func (s *Redis) Close() {
	s.conn.Close()
	log.Printf("Return conn to redis_conn_pool, active connections in the pool is %d", pool.ActiveCount())
}

// Operations for Hash
func (s *Redis) Hget(key, field string) ([]byte, error) {
	log.Printf("Redis HGet: %s %s", key, field)
	value, err := s.conn.Do("HGET", key, field)
	if err != nil {
		return nil, err
	}
	if value == nil {
		return nil, nil
	}
	bytes, err := redis.Bytes(value, err)
	if err != nil {
		return nil, err
	}
	return bytes, err
}

func (s *Redis) Hmget(key string, fields []string) (map[string][]byte, error) {
	if len(fields) == 0 {
		return nil, nil
	}
	log.Printf("Redis HMGet: %s %#v", key, fields)

	args := make([]interface{}, len(fields)+1)
	args[0] = key
	for i, v := range fields {
		args[i+1] = v
	}
	values, err := redis.Values(s.conn.Do("HMGET", args...))
	if err != nil {
		return nil, err
	}
	if len(values) == 0 {
		return nil, nil
	}
	entries := make(map[string][]byte, len(fields))
	for i, value := range values {
		if value == nil {
			continue
		}
		bytes, err := redis.Bytes(value, err)
		if err != nil {
			return nil, err
		}
		entries[fields[i]] = bytes
	}
	return entries, nil
}

func (s *Redis) hset(key, field string, value []byte) error {
	log.Printf("Redis HSet: %s %s %#v", key, field, value)
	_, err := s.conn.Do("HSET", key, field, value)
	return err
}

func (s *Redis) hdel(key string, fields []string) error {
	if len(fields) == 0 {
		return nil
	}
	log.Printf("Redis HDel: %s %#v", key, fields)

	args := make([]interface{}, len(fields)+1)
	args[0] = key
	for i, v := range fields {
		args[i+1] = v
	}
	_, err := s.conn.Do("HDEL", args...)
	return err
}

// Operations for Key
func (s *Redis) delete(key string) error {
	log.Printf("Redis Delete: %s", key)
	_, err := s.conn.Do("DEL", key)
	return err
}
