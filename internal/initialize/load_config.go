package initialize

import (
	"fmt"

	"github.com/quanminhvu/go-ecommerce-backend-api/global"
	"github.com/spf13/viper"
)

func LoadConfig() {
	viper := viper.New()
	viper.AddConfigPath("./config/")
	viper.SetConfigName("local")
	viper.SetConfigType("yml")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
	fmt.Println("Server running on port:", viper.Get("server.port"))

	err = viper.Unmarshal(&global.Config)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

}
