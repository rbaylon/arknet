package core

import (
	"arknet/modules/common/model"
	"github.com/gin-gonic/gin"
	"net/http"
  "github.com/gin-contrib/sessions"
)

func GetIndex() gin.HandlerFunc {
	fn := func(c *gin.Context) {
    c.HTML(
      http.StatusOK,
      "a_dashboard.html",
      gin.H{
        "title": "Welcome to Arknet",
      },
    )
	}
	return gin.HandlerFunc(fn)
}

func GetRules() gin.HandlerFunc {
  fn := func(c *gin.Context) {
    c.HTML(
      http.StatusOK,
      "a_rules.html",
      gin.H{
        "title": "Firewall rules",
      },
    )
  }
  return gin.HandlerFunc(fn)
}

func Login() gin.HandlerFunc {
  fn := func(c *gin.Context) {
    if c.Request.Method == "POST" {
      common.BasicAuth(c)
      c.Redirect(http.StatusMovedPermanently, "/")
    } else {
      c.HTML(
        http.StatusOK,
        "login.html",
        gin.H{
          "title": "Login to Arknet",
          "flash": " ",
        },
      )
    }
  }

  return gin.HandlerFunc(fn)
}

func Unauthorized() gin.HandlerFunc {
  fn := func(c *gin.Context) {
    c.HTML(
      http.StatusOK,
      "a_401.html",
      gin.H{
      },
    )
  }

  return gin.HandlerFunc(fn)
}

func Logout () gin.HandlerFunc {
  fn := func(c *gin.Context) {
    session := sessions.Default(c)
    session.Set("authenticated", 0) // set just to mark the session that it is changed.
    session.Clear()
    session.Options(sessions.Options{Path: "/", MaxAge: -1})
    session.Save()
    c.Redirect(http.StatusTemporaryRedirect, "/") 
  }

  return gin.HandlerFunc(fn)
}

// set routes
func Setroutes(r *gin.Engine) {
	r.GET("/", common.BasicAuth, GetIndex())
  r.GET("/rules", common.BasicAuth, GetRules())
  r.GET("/login", Login())
  r.POST("/login", Login())
  r.GET("/logout", Logout())
  r.GET("/401", Unauthorized())
}
