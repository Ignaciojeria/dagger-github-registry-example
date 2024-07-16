package main

import (
	"example/app/shared/constants"
	_ "example/app/shared/infrastructure/healthcheck"
	_ "example/app/shared/infrastructure/observability"
	"example/app/shared/infrastructure/serverwrapper"
	_ "embed"
	"log"
	"os"

	ioc "github.com/Ignaciojeria/einar-ioc"
)

//go:embed .version
var version string

func main() {
	os.Setenv(constants.Version, version)
	if err := ioc.LoadDependencies(); err != nil {
		log.Fatal(err)
	}
	serverwrapper.Start()
}
