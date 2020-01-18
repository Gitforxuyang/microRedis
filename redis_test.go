package microRedis

import (
	"context"
	"fmt"
	"github.com/Gitforxuyang/microBase"
	"github.com/Gitforxuyang/microBase/trace"
	"github.com/Gitforxuyang/microRedis/command"
	"testing"
	"time"
)

func TestGetClient(t *testing.T) {
	tracer, closer, _ := trace.NewTracer("redis", "http://127.0.0.1:14268/api/traces")
	microBase.BaseTracer = tracer
	//client := redis.NewClient(&redis.Options{
	//	Addr:     "192.168.3.3:6379",
	//	Password: "",
	//	DB:       0,
	//	PoolSize: 5,
	//})
	//c := client.Set("demo", 1, time.Second*100)
	//fmt.Println(c.String())
	//fmt.Println(c.Args())
	//fmt.Println(c.Val())
	//fmt.Println(c.Name())

	//
	mr := microRedis{env: "local", clients: make(map[RedisClient]command.MicroRedisClient)}
	client := mr.GetClient(Main)
	err := client.Set(context.TODO(), "micro:demo", 1, time.Duration(-1))
	if err != nil {
		fmt.Printf(err.Error())
	}
	fmt.Printf("123")
	closer.Close()
}
