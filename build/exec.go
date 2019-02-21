package build

import (
	"os"
	"os/exec"
	"strings"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func (b *Builder) exec(name string, args ...string) error {
	cmd := exec.CommandContext(b.ctx, name, args...)
	logrus.Debugf("running %s", strings.Join(cmd.Args, " "))
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}
