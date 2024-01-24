package main

import (
	"github.com/gin-gonic/gin"
	"github.com/uptrace/uptrace-go/uptrace"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"net/http"
	"strings"
	v1 "uptrace-example/api/v1"
	"uptrace-example/biz"
	"uptrace-example/global"
	"uptrace-example/initializer"
	"uptrace-example/store"
)

func init() {
	// https://uptrace.dev/get/opentelemetry-go.html#uptrace-go
	// 设置 UPTRACE_DSN 地址环境变量，在 uptrace 控制台获取
	// https://app.uptrace.dev/
	// UPTRACE_DSN=https://xxxxxxxxxxxx@api.uptrace.dev?grpc=4317
	uptrace.ConfigureOpentelemetry(
		uptrace.WithServiceName("red-book-backend"),
		uptrace.WithServiceVersion("v1.0.0"),
		uptrace.WithDeploymentEnvironment("production"),
	)
}

// main
// go get github.com/uptrace/uptrace-go
// go get -u github.com/gin-gonic/gin
// go get go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin
// https://uptrace.dev/get/instrument/opentelemetry-gin.html#gin-instrumentation
// go get -u gorm.io/gorm
// go get github.com/glebarez/sqlite
// go get github.com/uptrace/opentelemetry-go-extra/otelgorm
// https://uptrace.dev/get/instrument/opentelemetry-gorm.html#what-is-gorm
func main() {
	global.DB = initializer.InitGorm()

	router := gin.Default()
	userStore := store.NewUserStore(global.DB)
	ctl := v1.NewUserController(biz.NewUserService(userStore))

	router.Use(otelgin.Middleware("read-book", otelgin.WithFilter(func(r *http.Request) bool {
		// 过滤掉 favicon.ico，该请求不需要记录
		if strings.LastIndex(r.URL.Path, "/favicon.ico") == 0 {
			return false
		}
		return true
	})))
	router.GET("/api/v1/user/:id", ctl.GetUserById)

	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
