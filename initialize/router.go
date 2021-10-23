package initialize

import (
	"grpc-myboot/core"
	"grpc-myboot/router"
)

func Routers() {
	//r := core.NewRoute()
	r := core.GetInstanceRouter()
	router.InitRouter(r)
}

