package cmd

import (
	"context"
	"os"

	"github.com/estenssoros/cow/initializer"
	"github.com/estenssoros/cow/sigtx"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(createCmd)
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "creates a react-typescript application",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("must supply application name")
		}
		appName := args[0]
		if _, err := os.Stat(appName); err == nil {
			return errors.Errorf("%s already exists in directory", appName)
		}
		ctx, cancel := sigtx.WithCancel(context.Background(), os.Interrupt)
		defer cancel()
		i, err := initializer.New(ctx, appName)
		if err != nil {
			return errors.Wrap(err, "new initializer")
		}

		if err := i.Run(); err != nil {
			return err
		}

		return nil
	},
}
