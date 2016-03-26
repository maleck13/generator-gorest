package main

import (
	"github.com/codegangsta/cli"
	"<%=basePackage %>/<%=baseName %>/cmd"
	"os"
)

func main() {

	app := cli.NewApp()
	app.Name = "<%=baseName %>"
	commands := []cli.Command{
		cmd.ServeCommand(),
	}
	app.Commands = commands
	app.Run(os.Args)

}

