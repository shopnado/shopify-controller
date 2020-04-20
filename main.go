//go:generate go run controller/codegen/cleanup/main.go
//go:generate rm -rf controller/generated
//go:generate go run controller/codegen/main.go

package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/shopnado/shopify-controller/api"
	"github.com/shopnado/shopify-controller/controller"
	"github.com/shopnado/shopify-controller/shopnado"
	"github.com/shopnado/shopify-controller/ui"
	"github.com/sirupsen/logrus"
	cli "github.com/urfave/cli/v2"
)

func run(c *cli.Context) error {
	logrus.Info("[shopnado] starting services")
	server := shopnado.NewServer()
	server.Register(
		api.NewService(c),
		controller.NewService(c),
		ui.NewService(c))
	server.Run()

	return nil
}

func main() {
	app := &cli.App{
		Name:   "new-command",
		Usage:  "",
		Action: run,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "kubeconfig",
				Aliases: []string{"k"},
				Usage:   "Filepath to the location of a kubeconfig file for a kubernetes server",
			},
			&cli.IntFlag{
				Name:    "threads",
				Aliases: []string{"t"},
				Usage:   "Controller threads to run",
				Value:   2,
			},
			&cli.StringFlag{
				Name:    "masterurl",
				Aliases: []string{"m"},
				Usage:   "Masterurl for kubernetes",
			},
			&cli.StringFlag{
				Name:    "namespace",
				Aliases: []string{"n"},
				Usage:   "Namespace to run the controller, \"default\" is the default namespace",
				Value:   "default",
			},
			&cli.BoolFlag{
				Name:    "debug",
				Aliases: []string{"d"},
				Usage:   "Turn on verbose debug logging",
			},
			&cli.BoolFlag{
				Name:    "quiet",
				Aliases: []string{"q"},
				Usage:   "Turn on off all logging",
			},
		},
		Before: func(ctx *cli.Context) error {
			if ctx.Bool("debug") {
				logrus.SetLevel(logrus.DebugLevel)
			}
			if ctx.Bool("quiet") {
				logrus.SetOutput(ioutil.Discard)
			}

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
