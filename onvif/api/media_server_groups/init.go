package mediaservergroups

import (
	"github.com/gin-gonic/gin"
)

// Init 初始化路由
func Init(router gin.IRouter) {
	router = router.Group("/media_server_groups")
	//
	router.GET("/", list)
	router.GET("/:id", get)
	router.POST("/", post)
	router.PATCH("/", patch)
	router.DELETE("/:id", delete)
	router.DELETE("/", batchDelete)
}
