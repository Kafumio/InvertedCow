package config

import "gopkg.in/ini.v1"

// InitSetting
//
//	@Description: 初始化配置
//	@param file 配置文件路径
//	@return error
func InitSetting(file string) (*AppConfig, error) {
	cfg, err := ini.Load(file)
	if err != nil {
		return nil, err
	}
	config := new(AppConfig)
	cfg.MapTo(config)

	config.MySqlConfig = NewMySqlConfig(cfg)
	config.RedisConfig = NewRedisConfig(cfg)
	config.EmailConfig = NewEmailConfig(cfg)
	config.CosConfig = NewCosConfig(cfg)
	return config, nil
}
