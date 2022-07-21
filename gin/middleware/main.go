package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	router := gin.New()
	// 添加 logger 和 recover 中间件
	// 这样添加的中间件是全局使用的中间件
 	router.Use(gin.Recovery(), gin.Logger())
	router.Use(MyLogger())	// 使用自定义中间件

	// 这样添加的中间件只在接收 "/goods" 路由组的请求才会使用
	//goodsGroup := router.Group("/goods")
	//goodsGroup.Use(AuthRequired)

	router.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"test": "t",
		})
	})

	router.Run(":8085")
}

func AuthRequired(c *gin.Context)  {

}

// 自定义中间件
func MyLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 记录时间
		t := time.Now()
		// 可以向 context 中添加一些信息
		c.Set("token", "123456")
		// 使用 c.Next() 继续执行原有的请求
		c.Next()
		// 记录时间
		end := time.Since(t)
		fmt.Printf("耗时： %v\n", end)
		// 在请求做完后可以做一些其他事情 比如记录状态信息
		status := c.Writer.Status()
		fmt.Println("状态：",status)
	}
}
