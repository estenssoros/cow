package build

import (
	"context"
	"os"
	"path/filepath"

	"github.com/estenssoros/cow/app"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type Builder struct {
	App       *app.App
	Root      string
	ctx       context.Context
	steps     []func() error
	cleanups  []func() error
	isWindows bool
}

func New(ctx context.Context, isWindows bool) (*Builder, error) {
	a, err := app.New()
	if err != nil {
		return nil, errors.Wrap(err, "creating new app")
	}

	b := &Builder{
		App:       a,
		ctx:       ctx,
		cleanups:  []func() error{},
		isWindows: isWindows,
	}

	b.steps = []func() error{
		b.prepTarget,
		b.buildYarn,
		b.cleanServerBuild,
		b.moveYarnBuild,
		b.buildAssets,
		b.installBuildDeps,
		b.buildBin,
	}

	formatter := &logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:03",
	}
	logrus.SetFormatter(formatter)
	logrus.SetLevel(logrus.DebugLevel)
	return b, nil
}

func (b *Builder) Run() error {
	for _, s := range b.steps {
		if err := s(); err != nil {
			return errors.WithStack(err)
		}
		os.Chdir(b.Root)
	}
	return nil
}

func (b *Builder) AbsoluteBinaryPath() string {
	return filepath.Join(b.Root, "bin")
}
