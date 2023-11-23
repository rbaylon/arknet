package main

import (
	"github.com/rbaylon/arknet/database"
  "github.com/rbaylon/arknet/modules/pf/model"
  "github.com/rbaylon/arknet/modules/os/model"
	"log"
)

func main() {
	var (
		gin_ip   = database.GetEnvVariable("GIN_IP")
		gin_port = database.GetEnvVariable("GIN_PORT")
	)

  log.Printf("Gin Socket: %s:%s\n", gin_ip, gin_port)

	db, err := database.ConnectToSQLite()
	if err != nil {
		log.Fatal(err)
	}

	pfmodel.MigrateDB(db)
  osmodel.MigrateDB(db)
}
