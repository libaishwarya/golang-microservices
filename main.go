package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/libaishwarya/myapp/catservice/realtimecatfact"
	"github.com/libaishwarya/myapp/server"
	"github.com/libaishwarya/myapp/store/mysql"
	thirdparty "github.com/libaishwarya/myapp/userservice/third_party"
)

func main() {
	db, err := mysql.NewDB() //Connects to a MySQL database using a function mysql.NewDB().
	if err != nil {          //If connection fails, it logs the error and exits.
		log.Fatalf("Could not connect to the database: %v", err)
	}
	defer db.Close()

	userStore := mysql.NewUserStore(db) //Creates a user store (mysql.NewUserStore(db)) to interact with the user-related data in the database.
	userservice := &thirdparty.ThirdParty{}
	cs := &realtimecatfact.RealTimeCatFact{}
	r := server.SetupRouter(userStore, userservice, cs) //Sets up an HTTP server using server.SetupRouter(userStore), likely with routes that use userStore.

	r.Run(":8080") //Starts the server on port 8080 (r.Run(":8080")).

}
