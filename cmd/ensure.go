package cmd

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var ensureCmd = &cobra.Command{
	Use:   "ensure",
	Short: "ensures all dependencies are met",
	RunE: func(cmd *cobra.Command, args []string) error {
		return errors.New("ensure not implemented")
	},
}
