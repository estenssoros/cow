package app

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

type App struct {
	Pwd         string `json:"pwd"`
	Root        string `json:"root"`
	GoPath      string `json:"go_path"`
	Name        string `json:"name"`
	Bin         string `json:"bin"`
	PackagePkg  string `json:"package_path"`
	ViewsPkg    string `json:"views_path"`
	ModelsPkg   string `json:"models_path"`
	WithDep     bool   `json:"with_dep"`
	WithWebpack bool   `json:"with_webpack"`
	WithYarn    bool   `json:"with_yarn"`
	WithDocker  bool   `json:"with_docker"`
}

func (a App) String() string {
	ju, _ := json.MarshalIndent(a, "", "\t")
	return string(ju)
}

func New() (*App, error) {
	pwd, _ := os.Getwd()
	name := filepath.Base(pwd)
	if f, err := os.Stat("server"); err != nil {
		return nil, errors.Wrap(err, "missing server folder in directory")
	} else {
		if !f.IsDir() {
			return nil, errors.New("server is not a directory")
		}
	}
	goPath := os.Getenv("GOPATH")
	app := &App{
		Pwd:        pwd,
		Root:       ".",
		GoPath:     goPath,
		Name:       name,
		Bin:        filepath.Join(pwd, "bin"),
		PackagePkg: "server",
		ModelsPkg:  "server/models",
		ViewsPkg:   "server/views",
	}

	if _, err := os.Stat(filepath.Join(pwd, "Gopkg.toml")); err == nil {
		app.WithDep = true
	}
	if _, err := os.Stat(filepath.Join(pwd, "webpack.config.js")); err == nil {
		app.WithWebpack = true
	}
	if _, err := os.Stat(filepath.Join(pwd, "yarn.lock")); err == nil {
		app.WithYarn = true
	}
	return app, nil
}
