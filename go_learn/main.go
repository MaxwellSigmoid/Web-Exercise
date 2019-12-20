package main

import (
	"github.com/gin-gonic/gin"
	//"go_learn/controllers"
	/*"go_learn/models"*/
	"go_learn/routers"
	"go_learn/database"
)

func main() {
	r := gin.Default()
	// 加载路由
	database.Init()
	routers.LoadRouters(r)

	r.Static("static","tmp2")
	r.Static("what","tmp")

	/*controllers.GetPosts(r.pool.Get().(*Context))*/

	r.Run(":8082")
}
