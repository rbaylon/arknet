package main

import (
	"arknet/database"
	"arknet/modules/common/model"
	"arknet/modules/os/model"
	"arknet/modules/pf/model"
	"arknet/modules/pf/view"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	var (
		gin_ip   = common.GetEnvVariable("GIN_IP")
		gin_port = common.GetEnvVariable("GIN_PORT")
	)

	log.Printf("Gin Socket: %s:%s\n", gin_ip, gin_port)

	db, err := database.ConnectToSQLite()
	if err != nil {
		log.Fatal(err)
	}

	pfmodel.MigrateDB(db)
	osmodel.MigrateDB(db)
	r := gin.Default()
	pfview.Setroutes(r, db)
	r.Run(gin_ip + ":" + gin_port)
}
