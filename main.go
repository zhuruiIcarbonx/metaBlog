package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/zhuruiIcarbonx/metaBlog/service"
)

func main() {

	fmt.Println("new project")
	router := gin.Default()

	//注册
	router.POST("/blog/v1/register", service.UserRegister)
	//登录
	router.POST("/blog/v1/login", service.UserLogin)

	routerGroup := router.Group("/blog/v1")
	routerGroup.Use(service.JWTAuthMiddleware())
	fmt.Println("---------------------------")

	//创建文章
	routerGroup.POST("/post/create", service.PostCreate)

	//读取文章列表
	routerGroup.GET("/post/list", service.PostList)

	//读取单个文章详细信息
	routerGroup.GET("/post/:userId", service.PostOne)

	//更新文章
	routerGroup.POST("/post/update", service.PostUpdate)

	//删除
	routerGroup.POST("/post/delete", service.PostDelete)

	//创建评论
	routerGroup.POST("/comment/create", service.CommentCreate)

	//获取文章所有评论
	routerGroup.GET("/comment/list", service.CommentList)

	router.Run(":8080")

}
