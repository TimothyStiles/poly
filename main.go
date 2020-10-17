package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

/******************************************************************************
Oct, 15, 2020

This file is special because it is the entry point for our command line utility.
It also acts as a general template that outlines everything available to the user.

Initial argparsing and app definition is done entirely through
"github.com/urfave/cli/v2" for which you can find the docs here:

https://github.com/urfave/cli/blob/master/docs/v2/manual.md

Essentially poly's app is defined via the &cli.App{} struct which you initialize
with data needed to run your app. In our case we're providing it Name, Usage, Flags,
and Commands at the top level. Commands can also be nested to provide n-level sub commands.

When naming new flags please make sure they don't collide with already existent
flags and try to follow these naming conventions:

http://www.catb.org/~esr/writings/taoup/html/ch10s05.html

Happy hacking,
Tim

******************************************************************************/

func main() {
	app := &cli.App{
		Name:  "poly",
		Usage: "A command line utility for engineering organisms.",

		// This is where you define global flags. Each sub command can also have its own flags that overide globals
		Flags: []cli.Flag{

			&cli.BoolFlag{
				Name:  "y",
				Usage: "Answers yes for all confirmations before doing something possibly destructive.",
			},

			&cli.StringFlag{
				Name:  "i",
				Usage: "Specify file input type or input path.",
			},

			&cli.StringFlag{
				Name:  "o",
				Usage: "Specify file output type or output path.",
			},

			&cli.BoolFlag{
				Name:  "-log",
				Value: false,
				Usage: "Forces output to stdout.",
			},
		},

		// This is where you start defining subcommands there's a lot of spacing to enhance readability since these nested brackets can be a little much.
		Commands: []*cli.Command{

			// defining the kind of *cli.Context our subcommand will use.
			{
				Name:    "convert",
				Aliases: []string{"c"},
				Usage:   "Convert a single file or set of files from one type to another. Genbank to Json, Json to Gff, etc.",

				// defining flags for this specific sub command
				Flags: []cli.Flag{

					&cli.StringFlag{
						Name:  "o",
						Value: "json",
						Usage: "Specify file output type. Options are Gff, gbk/gb, and json. Defaults to json.",
					},

					&cli.StringFlag{
						Name:  "i",
						Value: "",
						Usage: "Specify file input type. Options are Gff, gbk/gb, and json. Defaults to none.",
					},
				},
				// where we provide the actual function that is called by the subcommand.
				Action: func(c *cli.Context) error {
					convert(c)
					return nil
				},
			},

			{
				Name:    "hash",
				Aliases: []string{"ha"},
				Usage:   "Hash a sequence while accounting for circularity.",

				Flags: []cli.Flag{

					&cli.StringFlag{
						Name:  "f",
						Value: "blake3",
						Usage: "Specify hash function type. Has many options. Blake3 is probably fastest.",
					},

					&cli.StringFlag{
						Name:  "i",
						Value: "json",
						Usage: "Specify file input type. Options are Gff, gbk/gb, and json.",
					},

					&cli.StringFlag{
						Name:  "o",
						Value: "string",
						Usage: "Specify output type. Options are string and json. Defaults to string.",
					},

					&cli.BoolFlag{
						Name:  "stdout",
						Value: false,
						Usage: "Will write to standard out whenever applicable. Defaults to false.",
					},
				},
				Action: func(c *cli.Context) error {
					hash(c)
					return nil
				},
			},
		}, // subcommands list ends here
	} // app definition ends here

	err := app.Run(os.Args) // run app and log errors
	if err != nil {
		log.Fatal(err)
	}
}
