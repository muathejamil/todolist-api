package redis

import (
	json2 "encoding/json"
	"errors"
	"github.com/go-redis/redis/v7"
	"strconv"
	"time"
	"todolist-api/contracts"
)

// Cache cache struct.
type Cache struct {
	host    string
	db      int
	expires time.Duration
}

// NewRedisCache creates new redis cache.
func NewRedisCache(host string, db int, exp time.Duration) *Cache {
	return &Cache{
		host:    host,
		db:      db,
		expires: exp,
	}
}

// getClient gets new redis client.
func (cache *Cache) getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cache.host,
		Password: "",
		DB:       cache.db,
	})
}

// Set sets the todo to the redis cache.
func (cache *Cache) Set(key uint64, value contracts.TodoDTO) error {
	client := cache.getClient()
	json, err := json2.Marshal(value)
	if err != nil {
		return errors.New("unable to set todo to cache")
	}
	// set value to key, and it will expire in the amount of seconds
	formatUint := strconv.FormatUint(key, 10)
	client.Set(formatUint, json, cache.expires*time.Minute)
	return nil
}

// Get gets the todo from the redis cache.
func (cache *Cache) Get(key uint64) (contracts.TodoDTO, error) {
	client := cache.getClient()
	formatUint := strconv.FormatUint(key, 10)
	val, err := client.Get(formatUint).Result()
	if err != nil {
		return contracts.TodoDTO{}, errors.New("unable to find todo in cache")
	}
	todo := contracts.TodoDTO{}
	err = json2.Unmarshal([]byte(val), &todo)
	if err != nil {
		return contracts.TodoDTO{}, errors.New("unable to find todo in cache")
	}
	return todo, nil
}
