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
	err = cfg.MapTo(config)
	if err != nil {
		return nil, err
	}

	config.MySqlConfig = newMySqlConfig(cfg)
	config.RedisConfig = newRedisConfig(cfg)
	config.EmailConfig = newEmailConfig(cfg)
	return config, nil
}
