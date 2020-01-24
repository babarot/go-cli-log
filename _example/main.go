package main

import (
	"log"

	clilog "github.com/b4b4r07/go-cli-log"
)

func main() {
	clilog.Env = "YOUR_APP_LOG"
	clilog.SetOutput()

	log.Printf("[INFO] run main function")
}
