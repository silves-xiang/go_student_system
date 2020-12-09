package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main(){
	r:=gin.Default()
	r.StaticFS("/static",http.Dir("./static"))
	r.LoadHTMLGlob("./*/*/*")
	store :=cookie.NewStore([]byte("hello"))
	r.Use(sessions.Sessions("gin_session",store))
	router(r)
	r.Run(":8080")
}
