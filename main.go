package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"path"
)

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	router.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname")

		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})
	router.GET("/user/groups", func(c *gin.Context) {
		c.String(http.StatusOK, "The available groups are [...]")
	})
	router.POST("/post", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		name := c.PostForm("name")
		message := c.PostForm("message")

		fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
	})
	router.POST("/post1", func(c *gin.Context) {
		ids := c.QueryMap("ids")
		names := c.PostFormMap("names")
		fmt.Printf("ids: %v ;names: %v", ids, names)
	})
	router.MaxMultipartMemory = 8 << 20
	router.POST("upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		//打印日志
		log.Println(file.Filename)
		// 设置文件需要保存的指定位置并设置保存的文件名字
		dst := path.Join("./upload", file.Filename)
		if err := c.SaveUploadedFile(file, dst); err != nil {
			fmt.Println(err)
		}
		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})
	router.LoadHTMLGlob("templates/**/*")
	//router.LoadHTMLGlob("templates/users/index.html","templates/posts/index.html")
	router.GET("/posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/index.html", gin.H{
			"title": "posts/index",
		})
	})
	router.GET("/users/index.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.html", gin.H{
			"title": "users/index",
		})
	})
	router.Run(":8080")
}
