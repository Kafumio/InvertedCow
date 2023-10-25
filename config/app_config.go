package config

// AppConfig
// @Description:应用配置
type AppConfig struct {
	Port string `ini:"port"` //端口
	*MySqlConfig
	*RedisConfig
	*EmailConfig
	*CosConfig
}
