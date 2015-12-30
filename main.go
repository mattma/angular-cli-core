package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/mattma/angular-cli-core/cmd"
)

func main() {
	// See details: https://github.com/codegangsta/cli
	app := cli.NewApp()
	app.Name = "angular-cli-core"
	app.Usage = "Angular2 command line utility"

	app.Commands = []cli.Command{
		{
			Name:  "new",
			Usage: "Creates a new Angular2 project in the current folder.",
			Action: func(c *cli.Context) {
				println("added task: ", c.Args().First())
				cmd.New(c.Args().First())
			},
		},
	}

	app.Version = "0.0.1"

	app.Run(os.Args)
}
