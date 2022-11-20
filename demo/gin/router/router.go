package router

import "github.com/gin-gonic/gin"

func Routers() *gin.Engine {
	r := gin.Default()
	r.GET("/ads/:provider/:channel", func(c *gin.Context) {
		channel := c.Param("channel")
		provider := c.Param("provider")
		c.String(200, "channel %s, provider %s", channel, provider)
	})
	r.POST("/ads/:provider/:channel", func(c *gin.Context) {
		channel := c.Param("channel")
		provider := c.Param("provider")
		c.String(200, "channel %s, provider %s", channel, provider)
	})
	return r
}
