package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	"github.com/yuriadams/payout/controllers"
)

var (
	logOn   *bool
	urlBase string
)

func init() {
	logOn = flag.Bool("l", true, "log on/off")
	flag.Parse()
}

func main() {
	port := map[bool]string{true: os.Getenv("PORT"), false: "8080"}[os.Getenv("PORT") != ""]

	r := httprouter.New()
	r.POST("/api/games/:id/payout", controllers.GamePayoutHandler)

	controllers.Logging("Starting server %s...", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), r))
}
