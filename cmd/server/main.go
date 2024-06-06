package main

import (
	"flag"
	"fmt"

	"github.com/KrisjanisP/personal-dashboard/internal/app"
)

func main() {
	var port = ":3000"

	flag.StringVar(&port, "port", port, "port to listen on")
	flag.Parse()

	app := app.NewApp(fmt.Sprintf("localhost%s", port))
	app.ListenAndServe()

}
