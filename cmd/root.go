package cmd

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use: "buff",
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		color.Red("ERROR: " + err.Error())
	}
}
