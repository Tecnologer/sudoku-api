package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/tecnologer/go-secrets"
	"github.com/tecnologer/go-secrets/config"
	"github.com/tecnologer/sudoku/clients/sudoku-api/router"
)

var (
	port     = flag.Int("port", 8088, "Starting port server")
	verbouse = flag.Bool("v", false, "enanble verbouse log")
)

func main() {

	secrets.InitWithConfig(&config.Config{})

	checkVersion()

	if *verbouse {
		logrus.SetLevel(logrus.DebugLevel)
	}
	logrus.Debug("debugging")

	host := fmt.Sprintf(":%d", *port)

	r := router.Router()
	logrus.Infof("Starting server on %s...", host)
	logrus.Fatal(http.ListenAndServe(host, r))
}
