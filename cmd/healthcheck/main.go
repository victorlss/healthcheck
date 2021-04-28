package main

import (
	"github.com/joho/godotenv"
	checker2 "github.com/victorlss/healthcheck/pkg/checker"
	telegram2 "github.com/victorlss/healthcheck/pkg/notifier/telegram"
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
		if !checker2.IsAlive(site) {
			telegram2.Notify(site + " is down!")
		}
	}
}
