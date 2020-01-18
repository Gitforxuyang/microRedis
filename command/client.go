package command

import (
	"context"
	"github.com/Gitforxuyang/microRedis/wrapper"
	"github.com/go-redis/redis/v7"
	"time"
)

type MicroRedisClient interface {
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) (err error)
}

func NewClient(client *redis.Client, handleFunc wrapper.RedisTraceHandleFunc) MicroRedisClient {
	return &microRedisClient{client: client, handleFunc: handleFunc}
}

type microRedisClient struct {
	client     *redis.Client
	handleFunc wrapper.RedisTraceHandleFunc
}
