package main

import (
	"github.com/prs-io/openobserve-go/cmd"
)

func main() {
	//cobra.CheckErr(cmd.NewCLI().ExecuteContext(context.Background()))
	cmd.Execute()
}
