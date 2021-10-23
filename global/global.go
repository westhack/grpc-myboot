package global

import (
	"grpc-myboot/utils/timer"

	"golang.org/x/sync/singleflight"

	"go.uber.org/zap"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"grpc-myboot/config"
)

var (
	GormDB  *gorm.DB
	Redis   *redis.Client
	Config  config.Server
	Viper   *viper.Viper
	Logger  *zap.Logger
	Timer              = timer.NewTimerTask()
	ConcurrencyControl = &singleflight.Group{}
)
