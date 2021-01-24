package api

import (
	"authentication/internal/controller/index"
	"authentication/internal/controller/user"
	"authentication/internal/middleware/identification"
	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine)*gin.Engine{
	r.POST("/register",user.Register)
	r.GET("/login",user.Login)
	r.Use(identification.Identification)
	{
		r.GET("/index",index.Index)
		r.GET("/xxx",index.Index)
	}
	return r
}