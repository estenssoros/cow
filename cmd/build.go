package cmd

import (
	"context"
	"os"
	"time"

	"github.com/estenssoros/cow/build"
	"github.com/estenssoros/cow/sigtx"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var isWindows bool

func init() {
	RootCmd.AddCommand(buildCmd)
	buildCmd.Flags().BoolVarP(&isWindows, "windows", "w", false, "compile for windows")
}

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "builds applications into single binary",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, cancel := sigtx.WithCancel(context.Background(), os.Interrupt)
		defer cancel()
		b, err := build.New(ctx, isWindows)
		if err != nil {
			return errors.Wrap(err, "new build")
		}
		start := time.Now()
		if err := b.Run(); err != nil {
			return errors.WithStack(err)
		}

		logrus.Infof("Your application was successfully built at %s", b.AbsoluteBinaryPath())
		logrus.Infof("build took %v", time.Since(start))
		return nil
	},
}
