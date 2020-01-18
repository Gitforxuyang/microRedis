package microRedis

const (
	Main RedisClient = "main"
	Node RedisClient = "node"
)

type redisConfig struct {
	Addr     string
	Password string
	DB       int
	PoolSize int
}

type microRedisConfig struct {
	Clients map[RedisClient]redisConfig
}

func getConfig(env string, client RedisClient) redisConfig {
	switch env {
	case "local":
		return initLocalConfig().Clients[client]
	default:
		panic("env error")
	}
}

//初始化本地配置
func initLocalConfig() microRedisConfig {
	cfg := microRedisConfig{Clients: make(map[RedisClient]redisConfig)}
	cfg.Clients[Main] = redisConfig{Addr: "127.0.0.1:6379", Password: "", DB: 0, PoolSize: 5}
	cfg.Clients[Node] = redisConfig{Addr: "127.0.0.1:6379", Password: "", DB: 0, PoolSize: 5}
	return cfg
}
