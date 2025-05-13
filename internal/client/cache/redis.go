package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"openapi-client/internal/constants"
	"time"
)

func AccesTokenKey(email string) string {
	return fmt.Sprintf("token:%s", email)
}

func (r *RedisCache) Get(ctx context.Context, key string) (string, error) {

	normalizedKey := normalize(key)
	result, err := r.Client.Get(ctx, normalizedKey).Result()
	if err != nil {
		return "", err
	}

	return result, err
}

func (r *RedisCache) HGet(ctx context.Context, key string, field string) (string, error) {

	normalizedKey := normalize(key)
	result, err := r.Client.HGet(ctx, normalizedKey, field).Result()
	if err != nil {
		return "", err
	}

	return result, err
}

func (r *RedisCache) HGetAll(ctx context.Context, key string) (map[string]string, error) {

	normalizedKey := normalize(key)

	result, err := r.Client.HGetAll(ctx, normalizedKey).Result()
	if err != nil {
		return nil, err
	}

	return result, err
}

func (r *RedisCache) Set(ctx context.Context, key string, value interface{}) error {

	normalizedKey := normalize(key)
	encodedValue, err := json.Marshal(value)
	if err != nil {
		return err
	}

	_, err = r.Client.Set(ctx, normalizedKey, encodedValue, -1*time.Second).Result()
	if err != nil {
		return err
	}

	return nil
}

func (r *RedisCache) SetEx(ctx context.Context, key string, value interface{}, duration time.Duration) error {

	normalizedKey := normalize(key)
	encodedValue, err := json.Marshal(value)
	if err != nil {
		return err
	}
	_, err = r.Client.Set(ctx, normalizedKey, encodedValue, duration).Result()
	if err != nil {
		return err
	}

	return nil
}

func (r *RedisCache) Del(ctx context.Context, key string) error {

	normalizedKey := normalize(key)
	_, err := r.Client.Del(ctx, normalizedKey).Result()
	if err != nil {
		return err
	}

	return nil
}

func (r *RedisCache) HDel(ctx context.Context, key string, field string) error {

	normalizedKey := normalize(key)
	_, err := r.Client.HDel(ctx, normalizedKey, field).Result()
	if err != nil {
		return err
	}

	return nil
}

func normalize(key string) string {
	return fmt.Sprintf("%s:%s", constants.ServiceName, key)
}

func (r *RedisCache) GetAccessToken(ctx context.Context, email string) (string, error) {
	key := AccesTokenKey(email)
	normalizedKey := normalize(key)
	result, err := r.Client.Get(ctx, normalizedKey).Result()
	if err != nil {
		return "", err
	}

	return result, err
}

func (r *RedisCache) SaveAccessToken(ctx context.Context, email string, value interface{}, duration time.Duration) error {

	key := AccesTokenKey(email)
	normalizedKey := normalize(key)
	_, err := r.Client.Set(ctx, normalizedKey, value, duration).Result()
	if err != nil {
		return err
	}

	return nil
}
