package initialize

import (
	"fmt"
	"grpc-myboot/config"
	"grpc-myboot/global"
)

func Timer() {
	if global.Config.Timer.Start {
		for _, detail := range global.Config.Timer.Detail {
			go func(detail config.Detail) {
				global.Timer.AddTaskByFunc("ClearDB", global.Config.Timer.Spec, func() {

					fmt.Println("========> ClearDB")

				})
			}(detail)
		}
	}
}
