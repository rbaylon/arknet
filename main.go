package main

import (
	"arknet/database"
	"arknet/modules/common/model"
	"arknet/modules/os/model"
  "arknet/modules/os/view"
	"arknet/modules/pf/model"
	"arknet/modules/pf/view"
	"github.com/gin-gonic/gin"
  "arknet/view"
	"log"
  "github.com/gin-contrib/sessions"
  "github.com/gin-contrib/sessions/cookie"
)

func main() {
	var (
		gin_ip   = common.GetEnvVariable("GIN_IP")
		gin_port = common.GetEnvVariable("GIN_PORT")
    gin_secret = common.GetEnvVariable("GIN_SECRET")
	)

	log.Printf("Gin Socket: %s:%s\n", gin_ip, gin_port)

	db, err := database.ConnectToSQLite()
	if err != nil {
		log.Fatal(err)
	}

	pfmodel.MigrateDB(db)
	osmodel.MigrateDB(db)
	r := gin.Default()
  store := cookie.NewStore([]byte(gin_secret))
  r.Use(sessions.Sessions("arknet", store))
  r.LoadHTMLGlob("templates/*")
  r.Static("/assets", "./public/core/assets")
  r.Static("/js", "./public/core/js")
  r.Static("/css", "./public/core/css")
  core.Setroutes(r)
	pfapi.Setroutes(r, db)
  osapi.Setroutes(r, db)
	r.Run(gin_ip + ":" + gin_port)
}
