package initialize

import (
	"fmt"

	"github.com/quanminhvu/go-ecommerce-backend-api/global"
	"go.uber.org/zap"
)

func Run() {
	LoadConfig()
	m := global.Config.MySQL
	fmt.Println("Load config successfully with MySQL", m.Username)
	InitLogger()
	global.Logger.Info("Config log successfully", zap.String("ok", "success"))
	InitMySQL()
	InitRedis()
	r := InitRouter()

	r.Run(":8002")
}
