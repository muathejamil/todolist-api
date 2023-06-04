package redis

import (
	"errors"
	"github.com/go-redis/redis/v7"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
	"strconv"
	"time"
	"todolist-api/contracts/generated"
	"todolist-api/models"
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
func (cache *Cache) Set(key uint64, value models.Todo) error {
	// get cache client.
	client := cache.getClient()
	// marshaling the value.
	generatedDto := generated.TodoDTO{
		Id:          uint32(value.ID),
		Title:       value.Title,
		Description: value.Description,
		DueDay:      nil,
	}
	bytes, err := proto.Marshal(&generatedDto)
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
func (cache *Cache) Get(key uint64) (models.Todo, error) {
	// get cache client.
	client := cache.getClient()
	// converts key to string.
	formatUint := strconv.FormatUint(key, 10)
	// get the key value from the cache.
	val, err := client.Get(formatUint).Result()
	// handle error if exists.
	if err != nil {
		return models.Todo{}, errors.New("unable to find todo in cache")
	}
	// init new todoDto.
	generatedDto := generated.TodoDTO{}
	// unmarshal value.
	err = proto.Unmarshal([]byte(val), &generatedDto)
	// handle error in un-marshaling if exists.
	if err != nil {
		return models.Todo{}, errors.New("unable to find todo in cache")
	}
	result := models.NewTodo(uint(generatedDto.Id), generatedDto.Title, generatedDto.Description, ConvertTimestampToTime(generatedDto.DueDay))
	return result, nil
}

// ConvertTimestampToTime converts *timestamppb.Timestamp to *time.Time
func ConvertTimestampToTime(timestamp *timestamppb.Timestamp) *time.Time {
	if timestamp == nil {
		return nil
	}
	t := timestamp.AsTime()
	return &t
}
