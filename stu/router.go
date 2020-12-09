package stu

import "github.com/gin-gonic/gin"

func Router(r *gin.RouterGroup){
	r.GET("/login",login)
	r.GET("/regist",regist)
	r.GET("/rregist",rregist)
	r.GET("/rlogin",rlogin)
	r.GET("/index",index)
	r.GET("/edit",edit)
	r.POST("/redit",redit)
	r.GET("/migrate",migrate)
}
