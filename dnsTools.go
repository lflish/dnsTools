package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:   "dig",
				Action: Dig,
				Usage:  "dnsTools dig -d github.com [-s 8.8.8.8:53]",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "domain",
						Aliases:  []string{"d"},
						Required: true,
						Usage:    "--domain",
					},
					&cli.StringFlag{
						Name:     "server",
						Aliases:  []string{"s"},
						Required: false,
						Value:    "8.8.8.8:53",
						Usage:    "-s 8.8.8.8:53",
					},
				},
			},
			{
				Name:   "speed",
				Action: Speed,
				Usage:  "dnsTools speed -d github.com [-s 8.8.8.8:53] [-c 10] [-e 100]",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "domain",
						Required: true,
						Aliases:  []string{"d"},
						Usage:    "--domain",
					},
					&cli.StringFlag{
						Name:     "server",
						Aliases:  []string{"s"},
						Required: false,
						Value:    "8.8.8.8:53",
						Usage:    "--server",
					},
					&cli.IntFlag{
						Name:     "count",
						Aliases:  []string{"c"},
						Required: false,
						Value:    10,
						Usage:    "-c 10",
					},
					&cli.IntFlag{
						Name:     "time",
						Aliases:  []string{"t"},
						Required: false,
						Value:    100,
						Usage:    "-t 100",
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
