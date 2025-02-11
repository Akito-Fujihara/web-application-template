package main

import (
	"errors"

	"github.com/Akito-Fujihara/web-application-template/cmd/server/di"
)

func main() {
	app, cleanup, err := di.InitializeServer()
	defer cleanup()
	if err != nil {
		errmsg := errors.New("failed to initialize server")
		panic(errors.Join(err, errmsg))
	}
	app.Start()
}
