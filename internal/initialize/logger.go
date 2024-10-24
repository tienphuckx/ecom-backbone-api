package initialize

import (
	"github.com/tienphuckx/ecom-backbone-api.git/global"
	"github.com/tienphuckx/ecom-backbone-api.git/pkg/logger"
)

func InitLogger() {
	// Assuming the log settings are part of SysConfig
	logConfig := global.SysConfig.LogConfig

	// Pass the LogConfig to the logger initialization
	global.Logger = logger.NewLogger(logConfig)
}
