package initializer

import (
	"github.com/glebarez/sqlite"
	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	"gorm.io/gorm"
	"uptrace-example/global"
	"uptrace-example/store/model"
)

func InitGorm() {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if err = db.Use(otelgorm.NewPlugin()); err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&model.User{})
	if err != nil {
		panic(err)
	}

	// 初始化数据
	db.Create(&model.User{
		Name: "张三",
		Age:  18,
	})

	global.DB = db
}
