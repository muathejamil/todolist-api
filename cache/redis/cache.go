package redis

import (
	"errors"
	"github.com/go-redis/redis/v7"
	"google.golang.org/protobuf/proto"
	"strconv"
	"time"
	"todolist-api/handlers/contracts"
)

const Base int = 10

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
	// get cache client.
	client := cache.getClient()
	// marshaling the value.
	bytes, err := proto.Marshal(&value)
	// handle error if exist when marshaling.
	if err != nil {
		return errors.New("unable to set todo to cache")
	}
	// convert key to string with base.
	formatUint := strconv.FormatUint(key, Base)
	// set value to key, and it will expire in the amount of minutes.
	client.Set(formatUint, bytes, cache.expires*time.Minute)
	return nil
}

// Get gets the todo from the redis cache.
func (cache *Cache) Get(key uint64) (contracts.TodoDTO, error) {
	// get cache client.
	client := cache.getClient()
	// converts key to string.
	formatUint := strconv.FormatUint(key, 10)
	// get the key value from the cache.
	val, err := client.Get(formatUint).Result()
	// handle error if exists.
	if err != nil {
		return contracts.TodoDTO{}, errors.New("unable to find todo in cache")
	}
	// init new todoDto.
	todo := contracts.TodoDTO{}
	// unmarshal value.
	err = proto.Unmarshal([]byte(val), &todo)
	// handle error in un-marshaling if exists.
	if err != nil {
		return contracts.TodoDTO{}, errors.New("unable to find todo in cache")
	}
	return todo, nil
}
