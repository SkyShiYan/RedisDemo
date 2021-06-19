package data

import (
	"redis_demo/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo)

// Data .
type Data struct {
	// TODO warpped database client 包含了数据库连接对象以及Redis连接对象
	rClient *redis.Client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	rClient := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       12, // use default DB
	})

	cleanup := func() {
		rClient.Close()
		logger.Log(log.LevelInfo, "closing the data resources")
	}
	return &Data{
		rClient: rClient,
	}, cleanup, nil
}
