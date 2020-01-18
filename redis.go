package microRedis

import (
	"fmt"
	"github.com/Gitforxuyang/microBase"
	"github.com/Gitforxuyang/microRedis/command"
	"github.com/Gitforxuyang/microRedis/wrapper"
	"github.com/go-redis/redis/v7"
	"sync"
)

type RedisClient string

type MicroRedis interface {
	GetClient(client RedisClient) command.MicroRedisClient
}

type microRedis struct {
	sync.Mutex
	env     string
	clients map[RedisClient]command.MicroRedisClient
}

func MicroRedisInit() MicroRedis {
	return &microRedis{env: microBase.BaseConfig.ServerConfig.Env, clients: make(map[RedisClient]command.MicroRedisClient)}
}

func (m *microRedis) GetClient(redisClient RedisClient) command.MicroRedisClient {
	m.Lock()
	c := m.clients[redisClient]
	if c == nil {
		config := getConfig(m.env, redisClient)
		client := redis.NewClient(&redis.Options{
			Addr:     config.Addr,
			Password: config.Password,
			DB:       config.DB,
			PoolSize: config.PoolSize,
		})
		handleFunc := wrapper.RedisTraceWrapper(microBase.BaseTracer, fmt.Sprintf("%s:%d", redisClient, config.DB))
		c = command.NewClient(client, handleFunc)
	}
	fmt.Println(m.clients[redisClient])
	fmt.Println(c)
	fmt.Println("1111")
	m.clients[redisClient] = c
	m.Unlock()
	return c
}
