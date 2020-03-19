package main

import (
	"echo-example/routers"
)

func main() {
	// create a new echo instance
	e := routers.New()
	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
