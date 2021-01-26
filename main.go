package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/tecnologer/sudoku/clients/api/router"
)

var (
	port     = flag.Int("port", 8088, "Starting port server")
	verbouse = flag.Bool("v", false, "enanble verbouse log")
)

func main() {
	checkVersion()

	if *verbouse {
		logrus.SetLevel(logrus.DebugLevel)
	}

	host := fmt.Sprintf(":%d", *port)

	r := router.Router()
	logrus.Infof("Starting server on %s...", host)
	logrus.Fatal(http.ListenAndServe(host, r))
}
