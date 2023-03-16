package main

import (
	"os"

	"github.com/dimus/backme/cmd"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Logger = log.Output(
		zerolog.ConsoleWriter{
			Out:        os.Stderr,
			TimeFormat: "15:04:05",
		},
	)
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	cmd.Execute()
}
