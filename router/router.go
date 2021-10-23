package router

import (
	"grpc-myboot/api"
	"grpc-myboot/core"
)

func InitRouter(r *core.Router) {

	apis := api.NewApi()

	r.AddRoute("/getUser/:id/:name", apis.GetUser)
	r.AddRoute("/getUserJsonString", apis.GetUserJsonString)
}
