package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

// 初始化路由
func Routers() *gin.Engine {
	var Router = gin.Default()
	//Router.Use(middleware.LoadTls())  // 打开就能玩https了
	ApiGroup := Router.Group("")
	InitApiRouter(ApiGroup)
	return Router
}

func RunServer()  {
	Router := Routers()
	Router.Static("/form-generator", "./resource/page")
	address := fmt.Sprintf(":%d", viper.GetString("db.username"))
	s := &http.Server{
		Addr:           address,
		Handler:        Router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	time.Sleep(10 * time.Microsecond)
	fmt.Printf(`欢迎使用 Gin-Vue-Admin
	默认自动化文档地址:http://127.0.0.1%s/swagger/index.html
	默认前端文件运行地址:http://127.0.0.1:8080
`, s.Addr)
	s.ListenAndServe()
}

func Load(g *gin.Engine) *gin.Engine {
	g.Use(gin.Recovery())
	// 404
	g.NoRoute(func (c *gin.Context)  {
		c.String(http.StatusNotFound, "404 not found");
	})

	g.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "asd")
	})

	return g
}