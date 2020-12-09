package main

import (
	"github.com/gin-gonic/gin"
	"stusystem/admin"
	stu2 "stusystem/stu"
)

func router(r *gin.Engine){
	stu:=r.Group("/stu")
	ad:=r.Group("/admin")
	stu2.Router(stu)
	admin.Router(ad)
}
