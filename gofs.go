package main

import (
	"github.com/urfave/cli"
	"github.com/uz2020/gofs/internal/cmd"
	"os"
	log "unknwon.dev/clog/v2"
)

func init() {
	err := log.NewFile(100,
		log.FileConfig{
			Level:    log.LevelInfo,
			Filename: "logs/is.log",
			FileRotationConfig: log.FileRotationConfig{
				Rotate: true,
				Daily:  true,
			},
		},
	)
	if err != nil {
		panic("unable to create new logger: " + err.Error())
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "gofs"
	app.Usage = "a distributed file system written in Go"
	app.Commands = []cli.Command{
		cmd.Web,
		cmd.Store,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal("Failed to start application: %v", err)
	}

	log.Stop()
}
