package admin

import "github.com/gin-gonic/gin"

func Router(r *gin.RouterGroup){
	r.GET("/index",is_login,index)
	r.GET("/detail",detail)
	r.GET("/auto",auto)
	r.GET("/login",login)
	r.POST("/rlogin",rlogin)
}
