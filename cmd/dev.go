package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/estenssoros/cow/app"
	"github.com/estenssoros/cow/refresh"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"
)

var projectName string

func init() {
	RootCmd.AddCommand(devCMD)
	wd, _ := os.Getwd()
	projectName = filepath.Base(wd)
}

var devCMD = &cobra.Command{
	Use:   "dev",
	Short: "runs the development server",
	RunE: func(c *cobra.Command, args []string) error {
		defer func() {
			if r := recover(); r != nil {
				log.Println(r)
			}
		}()
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		wg, ctx := errgroup.WithContext(ctx)
		wg.Go(func() error {
			return errors.Wrap(startDevServer(ctx), "start dev server")
		})
		wg.Go(func() error {
			return errors.Wrap(startWebpack(ctx), "start webpack")
		})
		if err := wg.Wait(); err != nil {
			return errors.WithStack(err)
		}
		return nil
	},
}

func startDevServer(ctx context.Context) error {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("ERROR: ", r)
		}
	}()
	c := &refresh.Configuration{
		AppRoot: ".",
		IgnoredFolders: []string{
			"node_modules",
			"public",
			"src",
			"vendor",
		},
		IncludedExtensions: []string{
			".go",
		},
		BuildPath:    "tmp",
		BuildDelay:   time.Duration(200),
		BinaryName:   fmt.Sprintf("%s-build", projectName),
		CommandFlags: []string{},
		EnableColors: true,
		LogName:      "buff",
		Debug:        false,
	}
	r := refresh.NewWithContext(c, ctx)
	return r.Start()
}

func startWebpack(ctx context.Context) error {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("ERROR: ", r)
		}
	}()
	app, err := app.New()
	if err != nil {
		return errors.Wrap(err, "creating new app")
	}
	tool := "yarn"
	if _, err := os.Stat(filepath.Join(app.Root, "node_modules")); err != nil {
		if _, err := exec.LookPath(tool); err != nil {
			return errors.Errorf("no node_modules directory found, and couldn't find %s to install it with", tool)
		}
		cmd := exec.CommandContext(ctx, tool, "install")
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		if err := cmd.Run(); err != nil {
			return errors.WithStack(err)
		}
	}
	cmd := exec.CommandContext(ctx, tool, "start")
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

func craftBuildPath(path string) string {
	wd, _ := os.Getwd()
	return filepath.Join(filepath.Clean(wd), path)
}
