package admin

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
	"stusystem/stu"
)
func is_login(c *gin.Context){
	session:=sessions.Default(c)
	id:=session.Get("adminid")
	//fmt.Println(id)
	if id==nil{
		c.Redirect(http.StatusTemporaryRedirect,"/admin/login")
	}
}
func cmysql() *gorm.DB{
	db,_:=gorm.Open("mysql","root:@tcp(:3308)/stu?charset=utf8")
	return db
}
func index(c *gin.Context){
	var students []stu.Studesc
	db:=cmysql()
	db.Find(&students)
	c.HTML(http.StatusOK,"adminindex",students)
}
func detail(c *gin.Context){
	session :=sessions.Default(c)
	id:=c.Query("id")
	var Student stu.Studesc
	db:=cmysql()
	db.Find(&Student,id)
	fmt.Println(session.Get("adminid"))
	c.HTML(http.StatusOK,"detail",gin.H{
		"Student":Student,
		"Adminid":session.Get("adminid"),
	})
}
func auto(c *gin.Context){
	var student stu.Studesc
	var  admin stu.Admin
	id:=c.Query("id")//获取传参
	db:=cmysql()//链接数据库
	db.Find(&student,id)//找到student
	student.Isauto="yes"//更改
	adminid:=c.Query("adminid")
	db.Find(&admin,adminid)//找到admin
	fmt.Println(admin)
	student.Autoname=admin.Name
	db.Model(&stu.Studesc{}).Update(&student)

}
func login(c *gin.Context){
	c.HTML(http.StatusOK,"adminlogin",nil)
}
func rlogin(c *gin.Context){
	var admin stu.Admin
	name:=c.PostForm("name")
	pwd:=c.PostForm("password")
	db:=cmysql()
	db.Find(&admin,"name=?",name)
	//fmt.Println(admin)
	if admin.Name!="" && admin.Password==pwd{
		session:=sessions.Default(c)
		session.Set("adminid",admin.Id)
		session.Save()
		c.Redirect(http.StatusMovedPermanently,"/admin/index")
	}else{
		c.Redirect(http.StatusMovedPermanently,"/admin/login")
	}
}