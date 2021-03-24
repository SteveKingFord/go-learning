package global

import (
	"kingford-backend/config"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	DB      *gorm.DB
	Config config.Server
	Log    *zap.Logger
	Viper  *viper.Viper
)
