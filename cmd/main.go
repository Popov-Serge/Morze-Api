package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "", 0)
	webServer := server.GetServer(logger)
	webServer.ListenAndServe()
}
