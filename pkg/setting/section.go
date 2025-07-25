package setting

type Config struct {
	MySQL MySQLSetting `mapstructure:"mysql"`
}

type MySQLSetting struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Dbname   string `mapstructure:"dbname"`
	MaxIdle  int    `mapstructure:"maxIdleConns"`
	MaxOpen  int    `mapstructure:"maxOpenConns"`
	LifeTime int    `mapstructure:"connMaxLifeTime"`
}
