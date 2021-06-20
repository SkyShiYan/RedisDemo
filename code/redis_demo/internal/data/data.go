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
	// warpped database client
	rClient *redis.Client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	// 创建Redis客户端
	rClient := redis.NewClient(&redis.Options{
		Addr:     "192.168.1.235:6379",
		Password: "iam59!z$", // no password set
		DB:       12,         // use default DB
	})

	cleanup := func() {
		rClient.Close()
		logger.Log(log.LevelInfo, "closing the data resources")
	}
	return &Data{
		rClient: rClient,
	}, cleanup, nil
}
