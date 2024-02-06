package router

import (
	swaggerFiles "github.com/swaggo/files"
	"net/http"
	"novel/endpoint"
	"novel/middleware"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "novel/docs"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func prometheusF(r *gin.Engine) {
	opsProcessed := promauto.NewCounter(
		prometheus.CounterOpts{
			Name: "my_api_requests_total",
			Help: "The total number of requests to the my_api service.",
		},
	)
	r.Use(func(ctx *gin.Context) {
		if ctx.Request.RequestURI != "/metrics" {
			opsProcessed.Inc()
			ctx.Next()
		}
	})
}

func HttpRouter() *gin.Engine {
	r := gin.Default()
	prometheusF(r)
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	r.Use(middleware.Cors())
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/book/")
	})
	r.Static("/book/", "./resource/book/")

	user(r)
	novel(r)
	charge(r)

	admin(r)
	captcha(r)
	return r
}

func captcha(r *gin.Engine) {
	g := r.Group("captcha")
	g.GET("/generate", endpoint.CaptchaGenerate)
	g.GET("/image", endpoint.CaptchaImage)
}

func user(r *gin.Engine) {
	uSrv := endpoint.NewUserSrv()
	r.POST("/login", uSrv.Login)
	r.GET("/logout", uSrv.Logout, middleware.Auth())
	r.POST("/register", uSrv.Register)
	r.GET("/find/:phone", uSrv.Find)
	auth := r.Group("/user")
	auth.Use(middleware.Auth())
	auth.GET("/info", uSrv.UserInfo)
}

func novel(r *gin.Engine) {
	nSrv := endpoint.NewNovelSrv()
	novel := r.Group("/novel")
	novel.GET("/cates", nSrv.Cates)
	novel.GET("/list", nSrv.Novels)
	novel.POST("/index", nSrv.Index)
	novel.GET("/search", nSrv.SearchNovels)
	novel.GET("/novel", nSrv.Novel)
	novel.GET("/chapters", nSrv.Chapters)
	novel.Use(middleware.Auth())
	novel.GET("/chapter", nSrv.Chapter)
	novel.GET("/notes", nSrv.Notes)
	novel.GET("/note/del/:novel_id", nSrv.DelNote)
	novel.GET("/join-book", nSrv.JoinBook)
	novel.GET("/buy_logs", nSrv.BuyLogs)
	novel.GET("/buy_chapter", nSrv.BuyChapter)
}

func charge(r *gin.Engine) {
	chargeSrv := endpoint.NewChargeSrv()
	charge := r.Group("/charge", middleware.Auth())
	charge.POST("/create", chargeSrv.CreateOrder)
	charge.GET("/order", chargeSrv.QueryOrder)
	r.POST("/charge/callback", chargeSrv.Callback)
	r.GET("/charge/callback/USD", chargeSrv.USDCallback)
}
