package setup

import (
	"os"
	"path/filepath"
	"runtime"

	"log/slog"

	"github.com/TravisRoad/travis-blog-service/global"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
)

const (
	DEV_CONFIG_FILE  = "config.dev.yaml"
	PROD_CONFIG_FILE = "config.prod.yaml"
)

func InitViper() {
	mode, ok := os.LookupEnv("MODE")
	if !ok {
		mode = "dev"
	}

	v := viper.New()
	configFile := DEV_CONFIG_FILE

	switch mode {
	case global.DEV:
		gin.SetMode(gin.DebugMode)
		configFile = DEV_CONFIG_FILE
	case global.TEST:
		gin.SetMode(gin.TestMode)
		_, b, _, _ := runtime.Caller(0)
		path := filepath.Dir(filepath.Dir(b))
		configFile = filepath.Join(path, DEV_CONFIG_FILE)
	case global.PROD:
		gin.SetMode(gin.ReleaseMode)
		configFile = PROD_CONFIG_FILE
	}
	v.SetConfigFile(configFile)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		slog.Error(err.Error())
		panic(err)
	}
	if err := v.Unmarshal(&global.Config); err != nil {
		slog.Error(err.Error())
		panic(err)
	}
	RegisterEnv()
}

func RegisterEnv() {
	if token, ok := os.LookupEnv("TOKEN"); ok {
		global.Config.Token = token
	}
	if salt, ok := os.LookupEnv("SALT"); ok {
		global.Config.Salt = salt
	}
	if port, ok := os.LookupEnv("PORT"); ok {
		global.Config.Port = cast.ToInt(port)
	}
}
