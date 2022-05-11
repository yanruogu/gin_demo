package yap

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yanruogu/gin_demo/pkg/config"
	"github.com/yanruogu/gin_demo/pkg/logger"
	"github.com/yanruogu/gin_demo/pkg/mysql"
	"github.com/yanruogu/gin_demo/pkg/redis"

	ginzap "github.com/gin-contrib/zap"
)

var App *yap

const (
	DebugMode   = "dev"
	ReleaseMode = "release"
)

func init() {
	App = New()
}

type yap struct {
	Logger *logger.Logger
	DB     *mysql.DB
	Redis  *redis.RedisClient
	Engine *gin.Engine
	Config *config.Config
}

func New() *yap {
	// 生成配置
	//conf := config.New()
	//conf.Init(configFile)
	return &yap{
		Engine: gin.New(),
	}
}

func (y *yap) Init(configFile string) {

	// 初始化配置文件
	y.initConfig(configFile)
	// 注册日志组件
	if err := y.registerLogger(); err != nil {
		panic(fmt.Errorf("register logger failed,err: %s", err))
	}
	y.Logger.LoggerHandler.Info("logger init success")
	// 注册redis组件
	if err := y.registerRedis(); err != nil {
		panic(fmt.Errorf("register redis failed,err: %s", err))
	}
	y.Logger.LoggerHandler.Info("redis init success")
	// 注册orm
	if err := y.registerOrm(); err != nil {
		panic(fmt.Errorf("register orm failed,err: %s", err))
	}
	y.Logger.LoggerHandler.Info("orm init success")

	y.Engine.Use(ginzap.Ginzap(y.Logger.LoggerHandler, "", true))
	y.Engine.Use(ginzap.RecoveryWithZap(y.Logger.LoggerHandler, true))

}

func (y *yap) Start() {
	if y.Config.Server.Mode == ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	addr := fmt.Sprintf(":%s", strconv.Itoa(y.Config.Server.Port))
	if err := y.Engine.Run(addr); err != nil {
		panic(err)
	}
}

func (y *yap) registerLogger() error {
	y.Logger = logger.New(y.Config.Log)
	return y.Logger.Init()
}

func (y *yap) registerRedis() error {
	y.Redis = redis.New(y.Config.Redis)
	return y.Redis.Init()
}

func (y *yap) registerOrm() error {
	y.DB = mysql.New(y.Config.Database)
	return y.DB.Init()
}

func (y *yap) initConfig(configFile string) {
	y.Config = config.New()
	y.Config.Init(configFile)
}
