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
				Name:        "date",
				Usage:       "go-build-helper date [FORMAT_STRING]",
				Description: "Get current date with given format, default format is '2006-01-02 15:04:05 -0700 MST'",
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
				Name:        "version",
				Usage:       "go-build-helper version [GIT_REPO_LOCATION]",
				Description: "Get current git veresion, like 'git describe --tags'",
				Action: func(c *R.Context) error {
					repo := c.Args().First()
					output, err := V.Version(repo)
					if err == nil {
						fmt.Println(output)
					} else {
						fmt.Println("n/a")
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
