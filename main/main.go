package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"myweb/controller"
	"myweb/db"
	"myweb/po"
)

func main() {
	//初始化数据库连接
	db.Init()
	defer db.Close()
	//初始化entity
	po.SyncArticleTable()

	// 1.创建路由
	r := gin.Default()
	// 2.绑定路由规则，执行的函数
	// gin.Context，封装了request和response
	r.POST("/api/article/list", controller.ListArticle)
	r.POST("/api/article/add", controller.SaveArticle)
	r.Run("127.0.0.1:8080")
}
