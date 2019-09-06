package main

import(
	"fmt"
	"github.com/urfave/cli"
	"os"
)

var app cli.App

func init() {
	app.Name = "rpc-cli"
	app.Author = "Ge Xin"
	app.Email = "gexin1023@gmail.com"

	app.Commands = subcmds

}

func main() {
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
