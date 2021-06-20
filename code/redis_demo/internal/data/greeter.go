package data

import (
	"context"
	"redis_demo/internal/biz"
	"sync"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
)

type greeterRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewGreeterRepo(data *Data, logger log.Logger) biz.GreeterRepo {
	return &greeterRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *greeterRepo) CreateGreeter(ctx context.Context, g *biz.Greeter) error {
	return nil
}

func (r *greeterRepo) UpdateGreeter(ctx context.Context, g *biz.Greeter) error {
	return nil
}

func (r *greeterRepo) InsertRedis(ctx context.Context) error {
	// 连接Redis
	var sg sync.WaitGroup

	index, size := 10000, 5

	var terr error = nil

	for i := 0; i < size; i++ {
		sg.Add(1)
		go func() {
			// 连接并开始使用
			for i := 0; i < index/size; i++ {
				uuid := uuid.New().String()
				err := r.data.rClient.Set(uuid, uuid, 0).Err()
				if err != nil {
					r.log.Warn("插入失败-", err)
					terr = err
				}
			}

			sg.Done()
		}()
	}

	sg.Wait()

	return terr
}
