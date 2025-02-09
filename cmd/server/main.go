package main

import (
	"errors"

	"github.com/Akito-Fujihara/web-application-template/cmd/server/di"
)

func main() {
	app, err := di.InitializeServer()
	if err != nil {
		errmsg := errors.New("failed to initialize server")
		panic(errors.Join(err, errmsg))
	}
	app.Start()
}
