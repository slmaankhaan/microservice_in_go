// Main entry point to the broker service
package main

import (
	"fmt"
	"log"
	"net/http"
)

const WebPort = "80"

type Config struct {
}

func main() {
	app := Config{}

	log.Printf("Starting broker service on port %s\n", WebPort)

	// define http server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", WebPort),
		Handler: app.routes(),
	}

	// Start the server
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}

}
