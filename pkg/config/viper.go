package config

import (
	"fmt"

	"mygo_bangforai/api/model"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func InitConfig() error {
	viper.SetConfigName("config") // 配置文件名（不含扩展名）
	viper.SetConfigType("yaml")   // 配置文件类型
	viper.AddConfigPath(".")      // 配置文件路径

	viper.SetEnvPrefix("goai")// 设置环境变量前缀
	viper.AutomaticEnv()// 自动读取环境变量

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件发生变化", in.Name, in.Op)
	})
	
	if err := viper.ReadInConfig(); err != nil {// 读取配置文件
		return err
	}
	if err := viper.Unmarshal(&model.AppConfig); err != nil {// 解析配置到 AppConfig 结构体
		return err
	}
	
	return nil
}

func GetServerPort() string {
	return model.AppConfig.Server.Port
}