package common

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
  "github.com/gin-contrib/sessions"
)

type Line interface {
	Genline() string
}

type Useraccount struct {
  Username string `form:"username"`
  Password string `form:"password"`
}

func GetEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
	return os.Getenv(key)
}

func BasicAuth(c *gin.Context) {
	var (
		appuser = GetEnvVariable("APP_USER")
		pw      = GetEnvVariable("APP_USER_PASSWORD")
	)
  session := sessions.Default(c)
  authenticated := session.Get("authenticated")
  if authenticated != nil {
    log.Println("Session ok for", authenticated)
    return
  } else {
    log.Println("Session initiated.")
    var u Useraccount
    c.Bind(&u)
    if u.Username == appuser && u.Password == pw {
      log.WithFields(log.Fields{
        "user": u.Username,
      }).Info("User authenticated")
      session.Set("authenticated", u.Username)
      session.Save()
      return
    } else {
      c.Abort()
      c.Redirect(http.StatusTemporaryRedirect, "/login")
      return
    }
  }
}
