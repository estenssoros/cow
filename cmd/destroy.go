package cmd

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(destroyCmd)
}

func captureUserInput(message string) string {
	fmt.Print(message)
	var input string
	fmt.Scanln(&input)
	return input
}

var destroyCmd = &cobra.Command{
	Use:   "destroy",
	Short: "destroys a cow app",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("must supply application name")
		}
		confirm := captureUserInput(fmt.Sprintf("destroy %s? [y/n]: ", args[0]))
		if confirm != "y" {
			logrus.Warning("bailing out")
			return nil
		}
		if err := os.RemoveAll(args[0]); err != nil {
			return err
		}
		return nil
	},
}
