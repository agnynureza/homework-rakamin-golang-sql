package main

import (
	"os"

	"github.com/agnynureza/homework-rakamin-golang-sql/app"
	"github.com/agnynureza/homework-rakamin-golang-sql/cli"
)

func main() {
	c := cli.NewCli(os.Args)
	c.Run(app.Init())
}
