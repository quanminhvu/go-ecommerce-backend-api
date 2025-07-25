package initialize

import (
	"fmt"

	"github.com/quanminhvu/go-ecommerce-backend-api/global"
)

func Run() {
	LoadConfig()
	m := global.Config.MySQL
	fmt.Println("Load config successfully with MySQL", m.Username)
	InitLogger()
	InitMySQL()
	InitRedis()
	r := InitRouter()

	r.Run(":8002")
}
