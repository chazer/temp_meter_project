package main

import (
	"log"
	"net/http"
	"os"
	"tmeter/app"
	"tmeter/app/consts"
	"tmeter/app/modules/auth"
	"tmeter/app/modules/devices"
	"tmeter/app/modules/users"
	"tmeter/lib/debug"
	"tmeter/lib/env"
	"tmeter/lib/router"
)

func main() {
	debug.SetPrefix("{{func}}> ")
	debug.SetOutput(os.Stderr)

	listenInterface := env.GetEnvOrDefault(consts.EnvKeyListenHost, "0.0.0.0")
	listenPort := env.GetEnvOrDefault(consts.EnvKeyListenPort, "8080")

	factory := app.NewAppFactory()

	r := router.NewRouter()

	// TODO: add global json content type guard

	auth.Init(r, factory)
	devices.Init(r, factory)
	users.Init(r, factory)

	s := &http.Server{
		Addr:    listenInterface + ":" + listenPort,
		Handler: r,
	}

	debug.Printf("Start listening on %s:%s", listenInterface, listenPort)
	log.Fatal(s.ListenAndServe())
}
