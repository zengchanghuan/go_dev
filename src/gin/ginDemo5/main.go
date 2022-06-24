package main

import (
	"encoding/xml"
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type Article struct {
	Title   string `json:"title" xml:"title"`
	Content string `json:"content" xml:"content"`
}

// UnixToTime 时间戳转换成日期
func UnixToTime(timestamp int) string {
	fmt.Println(timestamp)
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-02 15:04:05")
}

func main() {
	// 创建一个默认的路由引擎
	r := gin.Default()
	//自定义模板函数  注意要把这个函数放在加载模板前
	r.SetFuncMap(template.FuncMap{
		"UnixToTime": UnixToTime,
	})
	//加载模板 放在配置路由前面
	r.LoadHTMLGlob("templates/**/*")
	//配置静态web目录   第一个参数表示路由, 第二个参数表示映射的目录
	r.Static("/static", "./static")

	r.GET("/", func(context *gin.Context) {
		username := context.Query("username")
		age := context.Query("age")
		page := context.DefaultQuery("page", "1")

		context.JSON(http.StatusOK, gin.H{
			"username": username,
			"age":      age,
			"page":     page,
		})
	})

	r.GET("/article", func(context *gin.Context) {
		id := context.DefaultQuery("id", "1")
		context.JSON(http.StatusOK, gin.H{
			"msg": "新闻详情",
			"id":  id,
		})
	})

	r.GET("/user", func(context *gin.Context) {
		context.HTML(http.StatusOK, "default/user.html", gin.H{})

	})

	r.POST("/doAddUser1", func(context *gin.Context) {
		username := context.PostForm("username")
		password := context.PostForm("password")
		age := context.DefaultPostForm("age", "20")

		context.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
			"age":      age,
		})
	})

	//获取 GET POST 传递的数据绑定到结构体
	r.GET("/getUser", func(context *gin.Context) {
		user := &UserInfo{}
		if err := context.ShouldBind(&user); err == nil {
			fmt.Printf("%#v", user)
			context.JSON(http.StatusOK, user)
		} else {
			context.JSON(http.StatusOK, gin.H{
				"error": err.Error(),
			})
		}
	})

	r.POST("/doAddUser2", func(context *gin.Context) {
		user := &UserInfo{}
		if err := context.ShouldBind(&user); err == nil {
			context.JSON(http.StatusOK, user)
		} else {
			context.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
		}
	})

	//获取 Post Xml 数据
	r.POST("/xml", func(c *gin.Context) {
		article := &Article{}

		xmlSliceData, _ := c.GetRawData() //获取 c.Request.Body 读取请求数据

		fmt.Println(xmlSliceData)

		if err := xml.Unmarshal(xmlSliceData, &article); err == nil {
			c.JSON(http.StatusOK, article)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
		}

	})

	//动态路由传值
	r.GET("/list/:cid", func(context *gin.Context) {
		cid := context.Param("cid")
		context.String(200, "%v", cid)
	})

	r.GET("/admin", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin/index.html", gin.H{
			"title": "后台首页",
		})
	})

	r.Run("8080")
}
