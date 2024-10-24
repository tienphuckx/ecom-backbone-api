package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/tienphuckx/ecom-backbone-api.git/pkg/logger"
	"github.com/tienphuckx/ecom-backbone-api.git/pkg/setting"
	"gorm.io/gorm"
)

var (
	SysConfig         setting.SysConfig
	Logger            *logger.LoggerZap
	ApplicationConfig setting.ApplicationConfig
	Mdb               *gorm.DB
	Rdb               *redis.Client
)
