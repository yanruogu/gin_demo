package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yanruogu/gin_demo/yap"
)

func RegisterRoute() {
	api := yap.App.Engine.Group("/api")
	{
		api.GET("/version", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": yap.App.Config.Version,
			})
		})
	}
}
