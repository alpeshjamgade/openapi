package cache

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	apmgoredis "go.elastic.co/apm/module/apmgoredisv8/v2"
	"open-api/config"
	"open-api/internal/logger"
	"time"
)

type ICache interface {
	Get(ctx context.Context, key string) (string, error)
	HGet(ctx context.Context, key string, field string) (string, error)
	HGetAll(ctx context.Context, key string) (map[string]string, error)
	Set(ctx context.Context, key string, value interface{}) error
	HSet(ctx context.Context, key string, field, value interface{}) error
	Del(ctx context.Context, key string) error
	HDel(ctx context.Context, key string, field string) error
}

type RedisCache struct {
	Host   string
	Port   string
	Client *redis.Client
}

func NewRedisCache(host string, port string) *RedisCache {
	return &RedisCache{Host: host, Port: port}
}

func (r *RedisCache) Connect(ctx context.Context) error {
	var err error
	var count int8

	Logger := logger.CreateLoggerWithCtx(ctx)

	for {
		r.Client = redis.NewClient(&redis.Options{
			Addr:               fmt.Sprintf("%s:%s", r.Host, r.Port),
			Password:           "",
			DB:                 0,
			MinIdleConns:       10,
			PoolSize:           config.RedisPoolSize,
			IdleTimeout:        time.Second * 10,
			IdleCheckFrequency: time.Second * 5,
		})

		r.Client = r.Client.WithContext(ctx)
		_, err = r.Client.Ping(ctx).Result()
		if err != nil {
			Logger.Errorf("unable to connect to redis %v", err)
			count++
		} else {
			Logger.Infof("connceted to redis at Host: %s, Port: %s", r.Host, r.Port)
			r.Client.AddHook(apmgoredis.NewHook())
			break
		}

		if count > 5 {
			Logger.Errorf(err.Error())
			return err
		}
		Logger.Warnf("Retrying in 5 seconds...")
		time.Sleep(5 * time.Second)
	}

	return nil
}
