package cache

import (
	"github.com/go-redis/redis/v8"
	"github.com/go-xweb/log"
	"github.com/spf13/viper"
)

type RedisCache struct {
	Game *redis.Client
}

type redisConfig struct {
	Addr     string `mapstructure:"addr"`
	Password string `mapstructure:"password"`
	Db       int    `mapstructure:"db"`
}

var RedisDB *RedisCache

func openClient(c *redisConfig) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     c.Addr,
		Password: c.Password, // no password set
		DB:       c.Db,       // use default DB
	})
	return rdb
}

func InitGameClient() *redis.Client {
	var c redisConfig
	if err := viper.UnmarshalKey("redis.game", &c); err != nil {
		log.Errorf("unable to decode database into struct. err:%v", err)
		panic(err)
	}
	return openClient(&c)
}

func (cache *RedisCache) Init() {
	RedisDB = &RedisCache{
		Game: InitGameClient(),
	}
}
