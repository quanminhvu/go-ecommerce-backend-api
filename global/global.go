package global

import (
	"github.com/quanminhvu/go-ecommerce-backend-api/pkg/logger"
	"github.com/quanminhvu/go-ecommerce-backend-api/pkg/setting"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Mdb    *gorm.DB
)

// Config, Redis, MySQL,...
