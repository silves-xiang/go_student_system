package stu

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)
func cmysql() *gorm.DB{
	db,_:=gorm.Open("mysql","root:@tcp(:3308)/stu?charset=utf8")
	return db
}
func login(c *gin.Context){
	c.HTML(http.StatusOK,"login",nil)
}
func regist(c *gin.Context){
	c.HTML(http.StatusOK,"regist",nil)
}
func rregist(c *gin.Context){
	var student student
	student.Name=c.Query("name")
	student.Password=c.Query("password")
	db:=cmysql()
	db.Create(&student)
	var descstu Studesc
	descstu.Name=student.Name
	descstu.Stuid=student.Id
	descstu.Isauto="no"
	db.Create(&descstu)
	c.Redirect(http.StatusTemporaryRedirect,"/stu/login")
}
func rlogin(c *gin.Context){
	var student,r student
	student.Name=c.Query("name")
	student.Password=c.Query("password")
	db:=cmysql()
	db.Find(&r,"name=?",student.Name)

	if r.Name==""{
		c.Redirect(http.StatusTemporaryRedirect,"/stu/login")
	}else{
		if r.Password==student.Password{
			session:=sessions.Default(c)
			session.Set("id",r.Id)
			session.Save()
			c.Redirect(http.StatusTemporaryRedirect,"index")
		}else{
			c.Redirect(http.StatusTemporaryRedirect,"/stu/login")
		}
	}
}
func index(c *gin.Context){
	db:=cmysql()
	session:=sessions.Default(c)
	id:=session.Get("id")
	var student Studesc
	db.Find(&student,"stuid=?",id)
	c.HTML(http.StatusOK,"index",student)
}
func edit(c *gin.Context){
	db:=cmysql()
	var  student Studesc
	session:=sessions.Default(c)
	id :=session.Get("id")
	db.Find(&student,"stuid=?",id)
	c.HTML(http.StatusOK,"edit",student)
}
func redit(c *gin.Context){
	var student Studesc
	db:=cmysql()
	session:=sessions.Default(c)
	id:=session.Get("id")
	db.Find(&student,"stuid=?",id)
	student.Name=c.PostForm("name")
	student.Hobby=c.PostForm("hobby")
	student.Address=c.PostForm("address")
	student.Sex=c.PostForm("sex")
	student.School=c.PostForm("school")
	student.Isauto="no"
	student.Autoname="default name"
	fmt.Println(&student)
	db.Model(&Studesc{}).Update(&student)
	c.Redirect(http.StatusMovedPermanently,"/stu/index")
}
func migrate(c *gin.Context){
	db:=cmysql()
	db.AutoMigrate(&Studesc{})
}