package config

import (
	//"fmt"
	"mygo_bangforai/api/errors"
	"mygo_bangforai/api/model"
	"mygo_bangforai/pkg/interfacer"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var logger = interfacer.GetLogger()

func InitConfig() error {
	viper.SetConfigName("config") // 配置文件名（不含扩展名）
	viper.SetConfigType("yaml")   // 配置文件类型
	viper.AddConfigPath(".")      // 配置文件路径

	viper.SetEnvPrefix("goai")// 设置环境变量前缀
	viper.AutomaticEnv()// 自动读取环境变量

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		logger.Info("配置文件发生变化", zap.String("filename", in.Name), zap.String("operation", in.Op.String()))
	})
	
	if err := viper.ReadInConfig(); err != nil {// 读取配置文件
		err= errors.WrapError(err, errors.ConfigError, "读取配置文件失败", "pkg/config/viper.ConfigFileUsed()")
		return err
	}
	if err := viper.Unmarshal(&model.AppConfig); err != nil {// 解析配置到 AppConfig 结构体
		err= errors.WrapError(err, errors.ConfigError, "解析配置文件失败", "pkg/config/viper.Unmarshal()")
		return err
	}
	
	return nil
}

func GetServerPort() string {
	if model.AppConfig.Server.Port == "" {
		errors.NewError(errors.ConfigError, "服务器端口未配置", "pkg/config/viper.GetServerPort()")
	}
	return model.AppConfig.Server.Port
}

func GetDatabaseConfig() model.Database {
	if model.AppConfig.Database.Host == "" {
		errors.NewError(errors.ConfigError, "数据库主机未配置", "pkg/config/viper.GetDatabaseConfig()")
	}
	return model.AppConfig.Database
}

func GetLoggerConfig() model.Logger {
	if model.AppConfig.Logger.Level == "" {
		errors.NewError(errors.ConfigError, "日志级别未配置", "pkg/config/viper.GetLoggerConfig()")
	}
	return model.AppConfig.Logger
}

func GetJWTConfig() model.JWT {
	if model.AppConfig.JWT.Secret == "" {
		errors.NewError(errors.ConfigError, "JWT 密钥未配置", "pkg/config/viper.GetJWTConfig()")
	}
	return model.AppConfig.JWT
}

func GetServerConfig() model.Server {
	if model.AppConfig.Server.Name == "" {
		errors.NewError(errors.ConfigError, "服务器名称未配置", "pkg/config/viper.GetServerConfig()")
	}
	return model.AppConfig.Server
}