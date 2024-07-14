package grpc

import (
	"context"
)

// HelloServiceImpl 实现服务接口
type HelloServiceImpl struct {
	UnimplementedHelloServiceServer
}

func (p *HelloServiceImpl) Hello(
	ctx context.Context, args *String,
) (*String, error) {
	reply := &String{Value: "hello:" + args.GetValue()}
	return reply, nil
}
