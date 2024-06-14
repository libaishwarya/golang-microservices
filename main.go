package main

import (
	"log"

	"github.com/libaishwarya/myapp/server"
	"github.com/libaishwarya/myapp/store/mysql"
)

func main() {
	db, err := mysql.NewDB()
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	defer db.Close()

	userStore := mysql.NewUserStore(db)
	r := server.SetupRouter(userStore)
	r.Run(":8080")
}
