package delivery

import (
	athletcomp "github.com/ArdentK/db-cp-final/pkg/athletComp"
	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.RouterGroup, usecase athletcomp.ACUseCase) {
	h := newHandler(usecase)

	router.POST("/competitions-new", h.NewRow)
	router.POST("/competitions-del", h.DelRow)
	router.GET("/test", h.Test)
}
