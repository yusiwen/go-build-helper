package main

import (
	"fmt"
	R "github.com/urfave/cli/v2"
	"github.com/yusiwen/go-build-helper/constant"
	D "github.com/yusiwen/go-build-helper/date"
	V "github.com/yusiwen/go-build-helper/version"
	"log"
	"os"
	"strings"
)

func main() {
	app := &R.App{
		Name:    "Go build helper",
		Usage:   "My build helper for go projects",
		Version: strings.Join([]string{constant.Version, " (", constant.BuildTime, ")"}, ""),
		Commands: []*R.Command{
			{
				Name:  "date",
				Usage: "Get current date",
				Action: func(c *R.Context) error {
					format := c.Args().First()
					output, err := D.Date(format)
					if err == nil {
						fmt.Println(output)
					}
					return err
				},
			},
			{
				Name:  "version",
				Usage: "Get git version",
				Action: func(c *R.Context) error {
					output, err := V.Version()
					if err == nil {
						fmt.Println(output)
						fmt.Println(output)
					}
					return err
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
