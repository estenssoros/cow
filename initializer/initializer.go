package initializer

import (
	"context"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"

	uuid "github.com/satori/go.uuid"

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
		i.prepSubDirectory,
		i.prepSubDirectories,
		i.writeGitignore,
		i.prepServerFolder,
		i.writeBinData,
		i.runYarnInstall,
		i.removeNodeModules,
		i.runYarnInstall,
		i.cowBuild,
		i.gitInit,
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

func (i *Initializer) execDirectory(dir, name string, args ...string) error {
	cmd := exec.CommandContext(i.ctx, name, args...)
	logrus.Debugf("running %s", strings.Join(cmd.Args, " "))
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	wd, _ := os.Getwd()
	cmd.Dir = filepath.Join(wd, dir)
	err := cmd.Run()
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (i *Initializer) prepSubDirectory() error {
	if _, err := os.Stat(i.appName); err == nil {
		return errors.Errorf("prep sub directory: %s allready exists", i.appName)
	}
	if err := os.Mkdir(i.appName, 0700); err != nil {
		return err
	}
	return nil
}

func (i *Initializer) prepSubDirectories() error {
	for _, dir := range []string{
		"bin",
		"server",
		"src",
		"tmp",
	} {
		dir = filepath.Join(i.appName, dir)
		logrus.Infof("creating folder: %s", dir)
		if err := os.Mkdir(dir, 0700); err != nil {
			return errors.Wrapf(err, "prep sub directories: making subdirectory %s", dir)
		}
	}
	return nil
}

func (i *Initializer) prepServerFolder() error {
	dirs := []string{
		"api",
		"data",
		"models",
		"seed",
		"service",
	}
	for _, dir := range dirs {
		dir = filepath.Join("server", dir)
		logrus.Infof("creating folder: %s", dir)
		if err := os.Mkdir(filepath.Join(i.appName, dir), 0700); err != nil {
			return errors.Wrapf(err, "prep server folder: making subdirectory %s", dir)
		}
	}
	return nil
}

func (i *Initializer) gitInit() error {
	logrus.Info("creating new git repository")
	if err := i.exec("git", "init"); err != nil {
		return err
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

func ensureDirExists(fileName string) error {
	parDir := filepath.Dir(fileName)
	if _, err := os.Stat(parDir); err == nil {
		return nil
	}
	if err := os.MkdirAll(parDir, 0700); err != nil {
		return err
	}
	return nil
}

type TemplateData struct {
	ServerAPIPath string
	AppName       string
	SecretKey     string
}

func (i *Initializer) newTemplate() (*TemplateData, error) {
	gopath := os.Getenv("GOPATH") // /Users/estenssoros/go
	wd, err := os.Getwd()         // /Users/estenssoros/go/src/github.com/estenssoros/asdf
	if err != nil {
		return nil, err
	}
	if !strings.Contains(wd, gopath) {
		return nil, errors.New("must be installed in GOPATH")
	}
	wd = strings.Replace(wd, gopath, "", 1)[5:] // github.com/estenssoros/asdf
	wd = filepath.Join(wd, i.appName, "server", "api")
	return &TemplateData{
		ServerAPIPath: wd,
		AppName:       i.appName,
		SecretKey:     uuid.Must(uuid.NewV4()).String(),
	}, nil
}

func (i *Initializer) shouldTemplate(assetName string) bool {
	switch assetName {
	case "server/app.go", "package.json", "public/index.html", "server/api/routes.go", "server/service/jwt.go":
		return true
	}
	return false
}

func (i *Initializer) writeBinData() error {
	logrus.Info("writing bin data")
	tmplData, err := i.newTemplate()
	if err != nil {
		return err
	}
	for _, assetName := range AssetNames() {
		logrus.Info(assetName)
		data, err := Asset(assetName)
		if err != nil {
			return err
		}
		dst := filepath.Join(i.appName, assetName)
		if err := ensureDirExists(dst); err != nil {
			return err
		}
		f, err := os.Create(dst)
		if err != nil {
			return err
		}
		defer f.Close()
		if i.shouldTemplate(assetName) {

			tmpl, err := template.New("").Parse(string(data))
			if err != nil {
				return err
			}
			if err := tmpl.Execute(f, tmplData); err != nil {
				return err
			}
		} else {
			if _, err := f.Write(data); err != nil {
				return err
			}
		}

	}
	return nil
}

func (i *Initializer) runYarnInstall() error {
	if err := i.execDirectory(i.appName, "yarn", "install"); err != nil {
		return err
	}
	return nil
}

func (i *Initializer) removeNodeModules() error {
	if err := os.RemoveAll(filepath.Join(i.appName, "node_modules")); err != nil {
		return err
	}
	return nil
}

func (i *Initializer) cowBuild() error {
	if err := i.execDirectory(i.appName, "cow", "build"); err != nil {
		return err
	}
	return nil
}
