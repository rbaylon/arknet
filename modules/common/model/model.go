package common

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

type Line interface {
	Genline() string
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
	user, password, hasAuth := c.Request.BasicAuth()
	if hasAuth && user == appuser && password == pw {
		log.WithFields(log.Fields{
			"user": user,
		}).Info("User authenticated")
	} else {
		c.Abort()
		c.Writer.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
		c.JSON(http.StatusForbidden, gin.H{
			"error": "Access denied!",
		})
		return
	}
}
