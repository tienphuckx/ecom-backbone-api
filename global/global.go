package global

import (
	"github.com/tienphuckx/ecom-backbone-api.git/pkg/logger"
	"github.com/tienphuckx/ecom-backbone-api.git/pkg/setting"
)

var (
	SysConfig         setting.SysConfig
	Logger            *logger.LoggerZap
	ApplicationConfig setting.ApplicationConfig
)
