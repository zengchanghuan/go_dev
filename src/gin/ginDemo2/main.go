package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Article struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

func main() {
	r := gin.Default()
	//配置模板的文件 如果模板在多级目录里面的话需要这样配置 r.LoadHTMLGlob("templates/**/**/*") /** 表示目录
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.String(200, "值：%v", "首页")
	})
	r.GET("/json1", func(c *gin.Context) {
		c.JSON(200, map[string]interface{}{
			"success": true,
			"msg":     "你好gin",
		})
	})

	r.GET("/json2", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"success": true,
			"msg":     "你好gin--22",
		})
	})

	r.GET("/json3", func(c *gin.Context) {

		a := &Article{
			Title:   "我是一个标题",
			Desc:    "描述",
			Content: "测试内容",
		}
		c.JSON(200, a)
	})

	//响应Jsonp请求
	// http://localhost:8080/jsonp?callback=xxxx
	// xxxx({"title":"我是一个标题-jsonp","desc":"描述","content":"测试内容"});
	r.GET("/jsonp", func(c *gin.Context) {

		a := &Article{
			Title:   "我是一个标题-jsonp",
			Desc:    "描述",
			Content: "测试内容",
		}
		c.JSONP(200, a)
	})

	r.GET("/xml", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{
			"success": true,
			"msg":     "你好gin 我是一个xml",
		})
	})

	r.GET("/news", func(c *gin.Context) {
		//r.LoadHTMLGlob("templates/*")
		c.HTML(http.StatusOK, "news.html", gin.H{
			"title": "我是后台的数据",
		})
	})

	r.GET("/goods", func(c *gin.Context) {
		c.HTML(http.StatusOK, "goods.html", gin.H{
			"title": "我是商品页面",
			"price": 20,
		})
	})

	//	冒号不可少
	r.Run(":8000")
}
