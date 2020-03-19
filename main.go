package main

import (
	"echo-example/routers"
)

func main() {
	// create a new echo instance
	r := routers.New()
	// Start server
	r.Logger.Fatal(r.Start(":1323"))
}
