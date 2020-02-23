package command

import (
	"context"
	"time"
)

func (m *microRedisClient) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) (err error) {
	span := m.handleFunc(ctx, "Set")
	cmd := m.client.Set(key, value, expiration)
	_, err = cmd.Result()
	span.SetCmd(cmd.String())
	span.Finish(err)
	return err
}
