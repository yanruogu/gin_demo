package config

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	*Server
	*Database
	*Redis
	*Log
}

type Server struct {
	Version string `mapstructure:"version"`
	Name    string `mapstructure:"name"`
	Mode    string `mapstructure:"mode"`
	Port    int    `mapstructure:"port"`
}

type Database struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Dbname   string `mapstructure:"dbname"`
}

type Redis struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Db       int    `mapstructure:"db"`
	Password string `mapstructure:"password"`
}

type Log struct {
	Level      string `mapstructure:"level"`
	Path       string `mapstructure:"path"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxBackups int    `mapstructure:"max_backups"`
	MaxAge     int    `mapstructure:"max_age"`
}

func New() *Config {
	return new(Config)
}

func (c *Config) Init(configFile string) {
	// 1. 指定配置文件
	//viper.SetConfigFile("./conf/config.yaml")
	viper.SetConfigFile(configFile)
	// 2. 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	// 3. 映射配置至struct
	if err := viper.Unmarshal(c); err != nil {
		panic(fmt.Errorf("Unmarshal config faild, err: %s \n", err))
	}
	// 4. watch配置
	viper.WatchConfig()
	// 5. 配置重载
	viper.OnConfigChange(
		func(in fsnotify.Event) {
			fmt.Println("配置文件发生变更...")
			if err := viper.Unmarshal(c); err != nil {
				panic(fmt.Errorf("Unmarshal config failed,err: %s \n", err))
			}
		})
}
