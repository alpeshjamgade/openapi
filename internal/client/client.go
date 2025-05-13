package client

import (
	"context"
	"go.elastic.co/apm/module/apmhttp/v2"
	"net/http"
	"open-api-client/config"
	"open-api-client/internal/client/cache"
	"open-api-client/internal/client/db"
	"open-api-client/internal/logger"
	"os"
	"time"
)

var (
	CacheClient *cache.RedisCache = nil
	HttpClient  *http.Client      = nil
	DbClient    db.DB             = nil
)

func GetClients(ctx context.Context) (db.DB, *cache.RedisCache, *http.Client) {
	Logger := logger.CreateLoggerWithCtx(ctx)

	dbUrl := config.DatabaseURL
	DbClient = db.NewPostgresDB(dbUrl)
	err := DbClient.Connect(ctx)
	if err != nil {
		Logger.Panic(err)
		os.Exit(1)
	}

	CacheClient = cache.NewRedisCache(config.RedisHost, config.RedisPort)
	if err := CacheClient.Connect(ctx); err != nil {
		Logger.Panic(err)
		os.Exit(1)
	}

	HttpClient = apmhttp.WrapClient(&http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			MaxIdleConns:        500,
			MaxConnsPerHost:     500,
			MaxIdleConnsPerHost: 500,
			IdleConnTimeout:     20 * time.Second,
		},
	})

	return DbClient, CacheClient, HttpClient
}

func GetCacheClient(ctx context.Context) (*cache.RedisCache, error) {
	Logger := logger.CreateLoggerWithCtx(ctx)

	if CacheClient == nil {
		CacheClient = cache.NewRedisCache(config.RedisHost, config.RedisPort)
		if err := CacheClient.Connect(ctx); err != nil {
			Logger.Panic(err)
			return nil, err
		}
	}

	return CacheClient, nil
}
