package global

import (
	"github.com/TravisRoad/travis-blog-service/config"
	"gorm.io/gorm"
)

const (
	TEST = "TEST"
	DEV  = "DEV"
	PROD = "PROD"
)

var (
	Config *config.Config
	DB     *gorm.DB
)
