package lock

import (
	"log"
	"os"

	"github.com/go-redis/redis"
)

var logger = log.New(os.Stderr, "[redis]", log.LstdFlags)

type Redis struct {
	Client    *redis.Client
	RedisUser RedisUser
	PoolSize  int
}

type RedisUser struct {
	Address  string
	Password string
}

func (r *Redis) Init() {
	r.Client = redis.NewClient(&redis.Options{
		Addr:     r.RedisUser.Address,
		Password: r.RedisUser.Password,
		PoolSize: r.PoolSize,
	})
	if err := r.Client.Ping().Err(); err != nil {
		panic(err)
	}
}

func NewRedis(user RedisUser) *Redis {
	r := &Redis{RedisUser: user, PoolSize: 10}
	r.Init()
	logger.Println("redis connect successfully")
	return r
}
