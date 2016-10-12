package main

import (
	"fmt"
	"github.com/lcaballero/kublai/app/settings"
	"github.com/lcaballero/kublai/app/web"
	"github.com/lcaballero/kublai/cli"
	"github.com/lcaballero/kublai/conf"
	"github.com/lcaballero/kublai/shared"
	"os"
	"github.com/lcaballero/kublai/tools/pub"
)

func main() {
	config := cli.ParseArgs(os.Args...)
	shared.ShowJsonOrPanic(config, nil)

	keys, err := conf.LoadKeys(config.DeployedTo)
	shared.ShowJsonOrPanic(keys, err)

	st, err := settings.NewSettings(config, keys)
	if err != nil {
		panic(err)
	}

	ws, err := web.NewWebServer(st)
	if err != nil {
		panic(err)
	}

	if config.Command == "pub" {
		pub.Publish()
	} else {
		fmt.Println("starting web server")
		ws.Start()
	}
}
