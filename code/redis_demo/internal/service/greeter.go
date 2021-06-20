package service

import (
	"context"

	v1 "redis_demo/api/helloworld/v1"
	"redis_demo/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

// GreeterService is a greeter service.
type GreeterService struct {
	v1.UnimplementedGreeterServer

	uc  *biz.GreeterUsecase
	log *log.Helper
}

// NewGreeterService new a greeter service.
func NewGreeterService(uc *biz.GreeterUsecase, logger log.Logger) *GreeterService {
	return &GreeterService{uc: uc, log: log.NewHelper(logger)}
}

// SayHello implements helloworld.GreeterServer
func (s *GreeterService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	s.log.WithContext(ctx).Infof("SayHello Received: %v", in.GetName())

	if in.GetName() == "error" {
		// return nil, v1.ErrorUserNotFound("user not found: %s", in.GetName())
	}
	return &v1.HelloReply{Message: "Hello " + in.GetName()}, nil
}

// InsertRandomRedisData(context.Context, *v1.RRequest) (*v1.RReply, error)
// SayHello implements helloworld.GreeterServer
func (s *GreeterService) InsertRandomRedisData(ctx context.Context, in *v1.RRequest) (*v1.RReply, error) {
	s.log.WithContext(ctx).Infof("InsertRandomRedisData Received:")
	err := s.uc.InsertRedis(ctx)
	if err != nil {
		return nil, err
	}
	return &v1.RReply{Message: "插入Redis成功"}, nil
}
