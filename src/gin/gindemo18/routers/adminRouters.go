package routers

import (
	"gindemo18/controllers/admin"
	"gindemo18/middlewares"

	"github.com/gin-gonic/gin"
)

func AdminRoutersInit(r *gin.Engine) {
	//middlewares.InitMiddleware中间件
	adminRouters := r.Group("/admin", middlewares.InitMiddleware)
	{
		adminRouters.GET("/", admin.IndexController{}.Index)

		adminRouters.GET("/user", admin.UserController{}.Index)
		adminRouters.GET("/user/add", admin.UserController{}.Add)
		adminRouters.GET("/user/edit", admin.UserController{}.Edit)
		adminRouters.GET("/user/delete", admin.UserController{}.Delete)

		adminRouters.GET("/article", admin.ArticleController{}.Index)
		adminRouters.GET("/article/add", admin.ArticleController{}.Add)
		adminRouters.GET("/article/edit", admin.ArticleController{}.Edit)

		adminRouters.GET("/nav", admin.NavController{}.Index)

		adminRouters.GET("/student", admin.StudentController{}.Index)
		adminRouters.GET("/bank", admin.BankController{}.Index)
	}
}
