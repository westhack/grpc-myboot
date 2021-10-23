package api

import (
	"encoding/json"
	"github.com/golang/protobuf/proto"
	"go.uber.org/zap"
	"grpc-myboot/core"
	"grpc-myboot/global"
	"grpc-myboot/model/system"
	pb "grpc-myboot/proto"
	"strconv"
)

func (a *Api) GetUser(in *pb.RpcRequest, ps core.Params) (*pb.RpcResponse, error) {
	user := &pb.User{}
	err := proto.Unmarshal(in.Data, user)
	if err != nil {
		global.Logger.Error("===>", zap.Any("err", err.Error()))
		return a.ErrorResponse(err.Error())
	}

	global.Logger.Info("===>", zap.Any("data.user", user))
	global.Logger.Info("===>", zap.Any("params.id", ps.ByName("id")))
	global.Logger.Info("===>", zap.Any("params.name", ps.ByName("name")))

	var sysUser system.SysUser
	_ = global.GormDB.Where("id = ?", user.Id).Find(&sysUser).Error

	global.Logger.Info("===>", zap.Any("gorm user", sysUser))

	//b, _ := sysUser.Encode()

	s, _ := strconv.Atoi(ps.ByName("id"))
	user.Id = int32(s)
	user.Username = ps.ByName("name")

	global.Logger.Info("===>", zap.Any("user.Username", user.Username))

	b, err := proto.Marshal(user)

	return a.SuccessDataResponse("success", b)
}

func (a *Api) GetUserJsonString(in *pb.RpcRequest, ps core.Params) (*pb.RpcResponse, error) {
	user := &pb.User{}
	err := proto.Unmarshal(in.Data, user)
	if err != nil {
		return a.ErrorResponse(err.Error())
	}

	global.Logger.Info("===>", zap.Any("err", user))

	var sysUser system.SysUser
	_ = global.GormDB.Where("id = ?", user.Id).Find(&sysUser).Error

	global.Logger.Info("===>", zap.Any("err", sysUser))

	b, _ := json.Marshal(sysUser)
	// b, _ := sysUser.Encode()

	return a.SuccessDataResponse("success", b)
}