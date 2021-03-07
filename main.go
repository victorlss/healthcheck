package main

import (
	"github.com/joho/godotenv"
	"github.com/victorlss/healthcheck/checker"
	"github.com/victorlss/healthcheck/notifier/telegram"
	"os"
	"strings"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	sitesToCheck := strings.Split(os.Getenv("SITES_TO_CHECK"), ",")
	for _, site := range sitesToCheck {
		if !checker.IsAlive(site) {
			telegram.Notify(site + " is down!")
		}
	}
}
