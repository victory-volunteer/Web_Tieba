package routes

import (
	"bluebell/controllers"
	_ "bluebell/docs" // 千万不要忘了导入把你上一步生成的docs
	"bluebell/logger"
	"bluebell/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func Setup(mode string) *gin.Engine {

	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) //gin设置成发布模式
	}

	r := gin.New()
	//r.Use(logger.GinLogger(), logger.GinRecovery(true), middlewares.RateLimitMiddleware(2*time.Second, 1))
	//middlewares.RateLimitMiddleware(2*time.Second, 1)用来给全局添加限流组件，每两秒才可以访问一次
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.LoadHTMLFiles("./templates/index.html")
	r.Static("/static", "./static")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/ping", func(c *gin.Context) {
		//判断是否登录，判断请求头中是否有有效的JWT
		c.String(http.StatusOK, "ok")
	})

	v1 := r.Group("/api/v1")
	//注册
	v1.POST("/signup", controllers.SignUpHandler)

	//登录
	v1.POST("/login", controllers.LoginHandler)

	//根据时间或分数获取帖子列表
	v1.GET("/post/:id", controllers.GetPostDetailHandler)
	v1.GET("/posts", controllers.GetPostListHandler)
	v1.GET("/posts2", controllers.GetPostListHandler2)
	v1.GET("/community", controllers.CommunityHandler)
	v1.GET("/community/:id", controllers.CommunityDetailHandler)

	v1.Use(middlewares.JWTAuthMiddleware())
	{

		v1.POST("/post", controllers.CreatePostHandler)

		// 投票
		v1.POST("/vote", controllers.PostVoteController)
	}

	//pprof.Register(r) //性能优化：注册pprof相关路由

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
