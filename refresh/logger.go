package refresh

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/fatih/color"
	homedir "github.com/mitchellh/go-homedir"
)

const lformat = "=== %s ==="

type Logger struct {
	log *log.Logger
}

func NewLogger() *Logger {
	return &Logger{
		log: log.New(os.Stdout, "buff refresh ", log.LstdFlags),
	}
}

func (l *Logger) Error(msg interface{}, args ...interface{}) {
	l.log.Print(color.RedString(fmt.Sprintf(lformat, msg), args...))
}

func (l *Logger) Success(msg interface{}, args ...interface{}) {
	l.log.Print(color.GreenString(fmt.Sprintf(lformat, msg), args...))
}

func (l *Logger) Print(msg interface{}, args ...interface{}) {
	l.log.Printf(fmt.Sprintf(lformat, msg), args...)
}

var LogLocation = func() string {
	dir, _ := homedir.Dir()
	dir, _ = homedir.Expand(dir)
	dir = path.Join(dir, ".refresh")
	os.MkdirAll(dir, 0755)
	return dir
}

var ErrorLogPath = func() string {
	return path.Join(LogLocation(), ID()+".err")
}
