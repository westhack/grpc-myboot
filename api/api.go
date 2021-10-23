package api

import pb "grpc-myboot/proto"

type Api struct {

}

func NewApi() *Api {
	return new(Api)
}

func (a *Api) RpcResponse(code int32, msg string) (*pb.RpcResponse, error) {
	return &pb.RpcResponse{
		Code: code,
		Message: msg,
		Data: nil,
	}, nil
}

func (a *Api) SuccessResponse(msg string) (*pb.RpcResponse, error) {
	return &pb.RpcResponse{
		Code: 200,
		Message: msg,
		Data: nil,
	}, nil
}

func (a *Api) ErrorResponse(msg string) (*pb.RpcResponse, error) {
	return &pb.RpcResponse{
		Code: 500,
		Message: msg,
		Data: nil,
	}, nil
}

func (a *Api) SuccessDataResponse(msg string, data []byte) (*pb.RpcResponse, error) {
	return &pb.RpcResponse{
		Code: 200,
		Message: msg,
		Data: data,
	}, nil
}

func (a *Api) ErrorDataResponse(msg string, data []byte) (*pb.RpcResponse, error) {
	return &pb.RpcResponse{
		Code: 500,
		Message: msg,
		Data: data,
	}, nil
}