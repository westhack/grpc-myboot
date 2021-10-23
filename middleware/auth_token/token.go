package auth_token

import (
	"context"
	"go.uber.org/zap"
	"grpc-myboot/global"
	pb "grpc-myboot/proto"
	"reflect"

	"google.golang.org/grpc"
)

func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		global.Logger.Info("====>", zap.Any("req", req))

		if r, ok := reflect.ValueOf(req).Interface().(*pb.RpcRequest); ok {
			global.Logger.Info("====>", zap.Any("token", r.Token))

			// 认证 token
			if r.Token == "myBoot" {
				return handler(ctx, req)
			}

			// 认证失败 直接返回
			return &pb.RpcResponse{
				Code: 401,
				Message: "未登录",
				Data: nil,
			}, nil
		}

		return handler(ctx, req)
	}
}

func StreamServerInterceptor() grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		global.Logger.Info("====>", zap.Any("srv", srv))
		return handler(srv, stream)
	}
}
