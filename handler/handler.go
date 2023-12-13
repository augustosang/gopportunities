package handler

import (
	"github.com/augustosang/gopportunities/configs"
	"gorm.io/gorm"
)

var (
	logger *configs.Logger
	db     *gorm.DB
)

func InitializeHandler() {
	logger = configs.GetLogger("handler")
	db = configs.GetSQLite()
}
