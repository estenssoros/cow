package build

import (
	"os"
	"os/exec"
	"path/filepath"

	"github.com/estenssoros/cow/app"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func (b *Builder) buildYarn() error {
	logrus.Info("build yarn")
	tool := "yarn"
	if _, err := os.Stat(filepath.Join(b.App.Pwd, "node_modules")); err != nil {
		logrus.Info("creating node modules")
		if _, err := exec.LookPath(tool); err != nil {
			return errors.Errorf("no node_modules directory found, and couldn't find %s to install it with", tool)
		}
		if err := b.exec(tool, "install"); err != nil {
			return errors.WithStack(err)
		}
	}
	return b.exec(tool, "build")
}

func (b *Builder) cleanServerBuild() error {
	args := []string{
		"-rf",
		filepath.Join(b.App.Pwd, "server", "data", "build"),
	}
	return b.exec("rm", args...)
}

func (b *Builder) moveYarnBuild() error {
	args := []string{
		filepath.Join(b.App.Pwd, "build"),
		filepath.Join(b.App.Pwd, "server", "data"),
	}
	return b.exec("mv", args...)
}

// creates the target dir
func (b *Builder) prepTarget() error {
	logrus.Info("prep target")
	// Create output directory if not exists
	outputDir := filepath.Join(b.App.Pwd, "bin")
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		os.MkdirAll(outputDir, 0776)
		logrus.Infof("creating target dir %s", outputDir)
	}
	return nil
}

func (b *Builder) buildAssets() error {
	logrus.Info("baking assets into binary")
	if _, err := exec.LookPath("go-bindata"); err != nil {
		return errors.New("missing go-bindata on path")
	}
	tool := "go-bindata"
	args := []string{
		"-o",
		filepath.Join("server", "bindata.go"),
		filepath.Join("server", "data", "build"),
		filepath.Join("server", "data", "build", "static"),
		filepath.Join("server", "data", "build", "static", "css"),
		filepath.Join("server", "data", "build", "static", "js"),
		filepath.Join("server", "data", "build", "static", "media"),
	}
	return b.exec(tool, args...)
}

func (b *Builder) installBuildDeps() error {
	logrus.Info("install build deps")
	app, err := app.New()
	if err != nil {
		return errors.Wrap(err, "creating new app")
	}
	if !app.WithDep {
		logrus.Info("skipping dep")
		return nil
	}
	tool := "dep"
	args := []string{
		"ensure",
		"-v",
	}
	return b.exec(tool, args...)
}

func (b *Builder) cleanBin() error {
	binaryPath := filepath.Join(b.App.Pwd, "bin", b.App.Name+"-build")
	if _, err := os.Stat(binaryPath); err != nil {
		return nil
	}
	return b.exec("rm", binaryPath)
}

func (b *Builder) buildBin() error {
	args := []string{
		"build",
		"-v",
		"-o",
	}
	if b.isWindows {
		args = append(args, filepath.Join(b.App.Pwd, "bin", b.App.Name+"-build.exe"))
	} else {
		args = append(args, filepath.Join(b.App.Pwd, "bin", b.App.Name+"-build"))
	}
	cmd := exec.CommandContext(b.ctx, "go", args...)
	if b.isWindows {
		cmd.Env = os.Environ()
		cmd.Env = append(cmd.Env, "GOOS=windows", "GOARCH=386")
		logrus.Info("building windows executable")
	} else {
		logrus.Info("building darwin executable")
	}
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Dir = "server"
	if err := cmd.Run(); err != nil {
		return errors.WithStack(err)
	}
	return nil
}
