package cmd

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "initializes a react-typescript application",
	RunE: func(cmd *cobra.Command, arg []string) error {
		return errors.New("init not implemented")
	},
}
