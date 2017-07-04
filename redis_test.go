package redisPool

import (
	"os"
	"fmt"
	"testing"

	"github.com/garyburd/redigo/redis"
	"github.com/stretchr/testify/assert"
)

var (
	rds	*Redis
)

var (
	testKey		string = "TestRedisHashKey"
	nonexistKey	string = "TestRedisNonExistHashKey"
	emptykey	string = ""

	testField1	string = "field1"
	testField2	string = "field2"
	testField3	string = "field3"
	nonexistField	string = "non-exist-field"

	testValue1	[]byte = []byte("value1")
	testValue2	[]byte = []byte("value2")
	testValue3	[]byte = []byte("value3")

	testData	map[string][]byte = map[string][]byte{
				"field1" : []byte("value1"),
				"field2" : []byte("value2"),
				"field3" : []byte("value3"),
			}
)

func init() {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Printf("Failed connect to redis: %s\n", err)
		os.Exit(1)
	}

	rds = &Redis{
		conn:conn,
	}
}

func TestRedis_Hget_Hmget(t *testing.T) {
	// Key not exist
	bytes, err := rds.Hget(nonexistKey, nonexistField)
	assert.Nil(t, bytes)
	assert.Nil(t, err)

	// Empty key
	bytes, err = rds.Hget(emptykey, nonexistField)
	assert.Nil(t, bytes)
	assert.Nil(t, err)

	// HSet
	err = rds.hset(testKey, testField1, testValue1)
	assert.Nil(t, err)

	err = rds.hset(testKey, testField2, testValue2)
	assert.Nil(t, err)

	err = rds.hset(testKey, testField3, testValue3)
	assert.Nil(t, err)

	// HGet
	bytes, err = rds.Hget(testKey, testField1)
	assert.Equal(t, bytes, testValue1)
	assert.Nil(t, err)

	bytes, err = rds.Hget(testKey, nonexistField)
	assert.Nil(t, bytes)
	assert.Nil(t, err)

	// HMGet
	data, err := rds.Hmget(testKey, []string{testField1, testField2, testField3})
	assert.Equal(t, data, testData)
	assert.Nil(t, err)

	data, err = rds.Hmget(testKey, []string{testField1, testField2, testField3, nonexistField})
	assert.Equal(t, data, testData)
	assert.Nil(t, err)

	// HDel
	err = rds.hdel(testKey, []string{testField1, testField2, testField3})
	assert.Nil(t, err)

	err = rds.hdel(nonexistKey, []string{testField1, testField2, testField3})
	assert.Nil(t, err)

	// HGet
	bytes, err = rds.Hget(testKey, testField1)
	assert.Nil(t, bytes)
	assert.Nil(t, err)
}
