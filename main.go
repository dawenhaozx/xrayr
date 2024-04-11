package main

import (
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/wyx2685/XrayR/cmd"
)

func main() {
	time.Sleep(time.Duration(15-time.Now().Unix()%15) * time.Second)
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
