package main

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// 初始化 session 中间件
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	// 登录页面
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{})
	})

	// 登录接口
	router.POST("/login", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")

		// TODO: 在此处添加根据用户名和密码验证用户的代码

		if username == "admin" && password == "admin" {
			session := sessions.Default(c)
			session.Set("username", username)
			session.Save()

			c.JSON(http.StatusOK, gin.H{"msg": "登录成功"})
		} else {
			c.JSON(http.StatusOK, gin.H{"msg": "用户名或密码错误"})
		}
	})

	// 需要登录才能访问的接口
	router.GET("/profile", func(c *gin.Context) {
		session := sessions.Default(c)
		username := session.Get("username")

		if username == nil {
			c.Redirect(http.StatusFound, "/")
			return
		}

		c.HTML(http.StatusOK, "profile.html", gin.H{"username": username})
	})

	router.Run(":8080")
}
