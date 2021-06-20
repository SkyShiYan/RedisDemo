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

func (r *greeterRepo) InsertRedis(ctx context.Context, iType int64) error {
	// 连接Redis
	var sg sync.WaitGroup

	index, size := 10000, 20

	var terr error = nil

	for i := 0; i < size; i++ {
		sg.Add(1)
		go func() {
			// 连接并开始使用
			for i := 0; i < index/size; i++ {
				uuid := uuid.New().String()
				err := r.data.rClient.Set(uuid, getString(iType, uuid), 0).Err()
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

func getString(stringType int64, uuid string) string {
	switch stringType {
	case 1:
		// 长度是10
		return string(uuid[0:10])
	case 2:
		// 长度是20
		return string(uuid[0:20])
	case 3:
		// 长度是50
		return appendString(uuid, 50)
	case 4:
		// 长度是100
		return appendString(uuid, 100)
	case 5:
		// 长度是200
		return appendString(uuid, 200)
	case 6:
		// 长度是1024
		return appendString(uuid, 1024)
	case 7:
		// 长度是5120
		return appendString(uuid, 5120)
	default:
		return uuid
	}
}

func appendString(uuid string, length int) string {
	if len(uuid) >= length {
		return uuid
	}

	var newUuid = uuid
	for len(newUuid)*2 < length {
		newUuid += newUuid
	}

	return newUuid + newUuid[0:length-len(newUuid)]
}
