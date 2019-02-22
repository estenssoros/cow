package initializer

import (
	"context"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type Initializer struct {
	appName string
	ctx     context.Context
	steps   []func() error
}

func New(ctx context.Context, appName string) (*Initializer, error) {
	i := &Initializer{
		appName: appName,
		ctx:     ctx,
	}

	i.steps = []func() error{
		i.createReactApp,
		i.prepSubDirectories,
		i.writeGitignore,
		i.prepServerFolder,
		i.prepSrcFolder,
		i.writeBinData,
	}

	formatter := &logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:03",
	}
	logrus.SetFormatter(formatter)
	logrus.SetLevel(logrus.DebugLevel)
	return i, nil
}

func (i *Initializer) exec(name string, args ...string) error {
	cmd := exec.CommandContext(i.ctx, name, args...)
	logrus.Debugf("running %s", strings.Join(cmd.Args, " "))
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (i *Initializer) createReactApp() error {
	logrus.Infof("creating new app at %s", i.appName)
	args := []string{
		"create-react-app",
		i.appName,
		"--typescript",
	}
	if err := i.exec("npx", args...); err != nil {
		return err
	}
	logrus.Info("created new react application")
	return nil
}

func (i *Initializer) prepSubDirectories() error {
	for _, dir := range []string{"server", "bin", "tmp"} {
		dir = filepath.Join(i.appName, dir)
		logrus.Infof("creating folder: %s", dir)
		if err := os.Mkdir(dir, 0700); err != nil {
			return errors.Wrapf(err, "making subdirectory %s", dir)
		}
	}
	return nil
}

func (i *Initializer) prepServerFolder() error {
	for _, dir := range []string{"api", "data", "models", "seed", "service"} {
		dir = filepath.Join("server", dir)
		logrus.Infof("creating folder: %s", dir)
		if err := os.Mkdir(filepath.Join(i.appName, dir), 0700); err != nil {
			return errors.Wrapf(err, "making subdirectory %s", dir)
		}
	}
	return nil
}

func (i *Initializer) prepSrcFolder() error {
	for _, dir := range []string{"actions", "assets", "components", "constants", "modules", "reducers", "store", "views"} {
		dir = filepath.Join("src", dir)
		logrus.Infof("creating folder: %s", dir)
		if err := os.Mkdir(filepath.Join(i.appName, dir), 0700); err != nil {
			return errors.Wrapf(err, "making subdirectory %s", dir)
		}
	}
	return nil
}

func (i *Initializer) writeGitignore() error {
	logrus.Info("writing .gitignore")
	f, err := os.Create(filepath.Join(i.appName, ".gitignore"))
	if err != nil {
		return err
	}
	ignore := []string{
		"/node_modules",
		"/bin",
		"/public",
		"/tmp",
		"/vendor",
		"yarn.lock",
		"Gopkg.lock",
		"yarn-error.log",
		".DS_Store",
		"/server/data",
		"/server/bindata.go",
	}
	for _, line := range ignore {
		if _, err := f.Write([]byte(line + "\n")); err != nil {
			return err
		}
	}
	return nil
}

func (i *Initializer) Run() error {
	for _, step := range i.steps {
		if err := step(); err != nil {
			return err
		}
	}
	return nil
}

func (i *Initializer) writeBinData() error {
	for _, assetName := range AssetNames() {
		data, err := Asset(assetName)
		if err != nil {
			return err
		}
		assetName = filepath.Join(i.appName, assetName)
		logrus.Info(assetName)
		f, err := os.Create(assetName)
		if err != nil {
			return err
		}
		if _, err := f.Write(data); err != nil {
			return err
		}
	}
	return nil
}
