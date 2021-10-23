/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a server for Greeter service.
package main

import (
	"grpc-myboot/core"
	"grpc-myboot/global"
	"grpc-myboot/initialize"
)


//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

func main() {

	global.Viper = core.Viper()      // 初始化Viper
	global.Logger = core.Zap()       // 初始化zap日志库
	global.GormDB = initialize.Gorm() // gorm连接数据库
	initialize.Timer()
	initialize.Redis()
	initialize.Routers()

	if global.GormDB != nil {
		initialize.MysqlTables(global.GormDB) // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.GormDB.DB()
		defer db.Close()
	}

	core.RunWindowsServer()
}
