package refresh

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"runtime"
	"strings"
	"time"

	yaml "gopkg.in/yaml.v2"
)

type Configuration struct {
	AppRoot            string        `yaml:"app_root" json:"app_root"`
	AbsPath            string        `json:"abs_path"`
	IgnoredFolders     []string      `yaml:"ignored_folders" json:"ignored_folders"`
	IncludedExtensions []string      `yaml:"included_extensions" json:"included_extensions"`
	BuildTargetPath    string        `yaml:"build_target_path" json:"build_target_path"`
	BuildPath          string        `yaml:"build_path" json:"build_path"`
	BuildFlags         []string      `yaml:"build_flags" json:"build_flags"`
	BuildDelay         time.Duration `yaml:"build_delay" json:"build_delay"`
	BinaryName         string        `yaml:"binary_name" json:"binary_name"`
	CommandFlags       []string      `yaml:"command_flags" json:"command_flags"`
	CommandEnv         []string      `yaml:"command_env" json:"command_env"`
	EnableColors       bool          `yaml:"enable_colors" json:"enable_colors"`
	LogName            string        `yaml:"log_name" json:"log_name"`
	Debug              bool          `yaml:"-" json:"debug"`
}

func (c Configuration) String() string {
	ju, _ := json.MarshalIndent(c, "", "\t")
	return string(ju)
}

func (c *Configuration) FullBuildPath() string {
	buildPath := path.Join(c.BuildPath, c.BinaryName)
	if runtime.GOOS == "windows" {
		if !strings.HasSuffix(strings.ToLower(buildPath), ".exe") {
			buildPath += ".exe"
		}
	}
	return buildPath
}

func (c *Configuration) Load(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(data, c)
}

func (c *Configuration) Dump(path string) error {
	data, err := yaml.Marshal(c)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path, data, 0666)
}

func ID() string {
	d, _ := os.Getwd()
	return fmt.Sprintf("%x", md5.Sum([]byte(d)))
}
