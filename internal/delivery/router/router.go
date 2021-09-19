package router

import (
	"github.com/gin-gonic/gin"
	"prj/pkg/auth"
	"prj/pkg/middlewares"
)

func homeHandler(c *gin.Context) {
	//connect, err := utils.Connect()
	//if err != nil {
	//	return
	//}
	//data := map[string] string{
	//	"host": c.Request.Host,
	//	"method": c.Request.Method,
	//	"uri": c.Request.RequestURI,
	//}
	//err := json.NewEncoder(c.Writer).Encode(data)
	//if err != nil {
	//	return
	//}
}

func New() *gin.Engine {
	var router = gin.New()
	router.Use(middlewares.JsonHeaderMiddleware())
	router.GET("/", homeHandler)
	router.POST("/auth", auth.Authenticate)

	return router
}
