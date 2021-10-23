package core

import (
	"context"
	"encoding/json"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/grpc-ecosystem/go-grpc-middleware/ratelimit"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"grpc-myboot/global"
	middleware_auth_token "grpc-myboot/middleware/auth_token"
	pb "grpc-myboot/proto"
	"log"
	"net"
)

type server struct {
	pb.UnimplementedMyServiceServer
}

func (s *server) RpcExecute(ctx context.Context, in *pb.RpcRequest) (*pb.RpcResponse, error) {

	r := GetInstanceRouter()
	path := in.Route

	if root := r.Trees["POST"]; root != nil {
		if handle, ps, _ := root.getValue(path, r.getParams); handle != nil {
			if ps != nil {
				return handle(in, *ps)
				r.putParams(ps)
			} else {
				return handle(in, nil)
			}
		}
	}

	return &pb.RpcResponse{
		Code:    404,
		Message: "无效的请求",
		Data:    nil,
	}, nil
}

func (s *server) RpcJsonExecute(ctx context.Context, in *pb.RpcJsonRequest) (*pb.RpcJsonResponse, error) {
	log.Printf("Received: %v", in.GetData())

	user := &pb.User{}
	json.Unmarshal([]byte(in.Data), user)

	return &pb.RpcJsonResponse{
		Code:    200,
		Message: user.Username + " " + in.Route,
		Data:    in.Data,
	}, nil
}

// alwaysPassLimiter is an example limiter which implements Limiter interface.
// It does not limit any request because Limit function always returns false.
type alwaysPassLimiter struct{}

func (*alwaysPassLimiter) Limit() bool {
	return false
}

func RunWindowsServer() {
	_port := global.Config.Grpc.Port
	lis, err := net.Listen("tcp", ":"+_port)
	if err != nil {
		global.Logger.Error("failed to listen:", zap.Any("err", err))
	}

	limiter := &alwaysPassLimiter{}

	s := grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_ctxtags.StreamServerInterceptor(),
			grpc_opentracing.StreamServerInterceptor(),
			grpc_zap.StreamServerInterceptor(global.Logger),
			grpc_recovery.StreamServerInterceptor(),
			ratelimit.StreamServerInterceptor(limiter),
			grpc_validator.StreamServerInterceptor(),
			middleware_auth_token.StreamServerInterceptor(),

		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_opentracing.UnaryServerInterceptor(),
			grpc_zap.UnaryServerInterceptor(global.Logger),
			grpc_recovery.UnaryServerInterceptor(),
			ratelimit.UnaryServerInterceptor(limiter),
			grpc_validator.UnaryServerInterceptor(),
			middleware_auth_token.UnaryServerInterceptor(),
		)),
	)
	pb.RegisterMyServiceServer(s, &server{})
	global.Logger.Info("server run success on ", zap.String("port", _port))
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
