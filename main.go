package main

import (
	"echo-example/models"
	"echo-example/routers"
)

func main() {
	// connect db
	db := models.ConnectDB()
	defer db.Close()
	// create a new echo instance
	r := routers.New()
	// Start server
	r.Logger.Fatal(r.Start(":1323"))
}
