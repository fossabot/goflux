package main

import (
	"log"
	"os"

	"github.com/Nerzal/goflux"
	"github.com/urfave/cli/v2"
)

func main() {
	gofluxClient := goflux.New()

	app := &cli.App{
		Name:  "goflux",
		Usage: "Used to automatically generate flux files",
		Commands: []*cli.Command{
			{
				Name:        "init",
				HelpName:    "Initialize",
				Description: "Initializes a new folder structure to work in",
				Usage:       "Initialize new project",
				Flags: []cli.Flag{
					cli.BashCompletionFlag,
					cli.HelpFlag,
					cli.VersionFlag,
					&cli.StringFlag{
						Required: true,
						Name:     "component",
					},
				},
				Action: func(c *cli.Context) error {
					projectName := c.String("component")
					err := gofluxClient.Initialize(projectName)
					if err != nil {
						return err
					}

					return nil
				},
			},
			{
				Name:        "backend",
				HelpName:    "backend",
				Description: "Creates folders and files for a backend service",
				Usage:       "Create files for a backend service",
				Flags: []cli.Flag{
					cli.BashCompletionFlag,
					cli.HelpFlag,
					cli.VersionFlag,
					&cli.StringFlag{
						Required: true,
						Name:     "component",
					},
					&cli.StringFlag{
						Required: true,
						Name:     "namespace",
					},
				},
				Action: func(c *cli.Context) error {
					projectName := c.String("component")
					namespace := c.String("namespace")

					err := gofluxClient.CreateBase(projectName, namespace)
					if err != nil {
						return err
					}

					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
