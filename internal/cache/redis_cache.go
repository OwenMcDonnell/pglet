package cache

import (
	"github.com/gomodule/redigo/redis"
	log "github.com/sirupsen/logrus"
)

type redisCache struct {
	redisAddr     string
	redisPassword string
	conn          redis.Conn
}

func newRedisCache(redisAddr string, redisPassword string) cacher {

	log.Println("Connecting to Redis server", redisAddr)

	conn, err := redis.Dial("tcp", redisAddr)
	if err != nil {
		log.Fatal(err)
	}

	if redisPassword != "" {
		if _, err := conn.Do("AUTH", redisPassword); err != nil {
			conn.Close()
			log.Fatal(err)
		}
	}

	return &redisCache{
		redisAddr:     redisAddr,
		redisPassword: redisPassword,
		conn:          conn,
	}
}

func (c *redisCache) exists(key string) bool {
	return false
}

func (c *redisCache) getString(key string) string {
	value, err := redis.String(c.conn.Do("GET", key))
	if err != nil {
		log.Fatal(err)
	}
	return value
}

func (c *redisCache) setString(key string, value string, expireSeconds int) {
	args := []interface{}{key, value}
	if expireSeconds > 0 {
		args = append(args, "EX", expireSeconds)
	}
	_, err := c.conn.Do("SET", args...)
	if err != nil {
		log.Fatal(err)
	}
}

func (c *redisCache) inc(key string, by int) int {
	return 0
}

func (c *redisCache) hashSet(key string, fields ...string) {

}

func (c *redisCache) hashGet(key string, field string) string {
	return ""
}

func (c *redisCache) hashGetAll(key string) map[string]string {
	return nil
}

func (c *redisCache) hashRemove(key string, fields ...string) {

}

func (c *redisCache) setGet(key string) []string {
	return nil
}

func (c *redisCache) setAdd(key string, value string) {

}

func (c *redisCache) setRemove(key string, value string) {

}

func (c *redisCache) remove(key string) {

}

//
// PubSub
// =============================

func (c *redisCache) subscribe(channel string) chan []byte {
	return nil
}

func (c *redisCache) unsubscribe(ch chan []byte) {

}

func (c *redisCache) send(channel string, message []byte) {

}
