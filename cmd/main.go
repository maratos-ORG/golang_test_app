package main

import (
	"fmt"
	"golang_test_app/internal/backend"
	"golang_test_app/internal/config"
	log "golang_test_app/internal/logging"
)

var (
	appName, gitTag, gitCommit, gitBranch string
)

func main() {
	conf, err := config.GetBackendParameters()
	if err != nil {
		log.Fatal("unable to generate config (pkg->config)")
	}
	fmt.Println(*conf.Port)
	fmt.Println(*conf.DBURL)
	backend.RunBackend(conf)

}
