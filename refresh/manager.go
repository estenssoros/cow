package refresh

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
)

type Manager struct {
	*Configuration
	Logger     *Logger
	Restart    chan bool
	gil        *sync.Once
	ID         string
	context    context.Context
	cancelFunc context.CancelFunc
}

func NewWithContext(c *Configuration, ctx context.Context) *Manager {
	ctx, cancelFunc := context.WithCancel(ctx)
	return &Manager{
		Configuration: c,
		Logger:        NewLogger(),
		Restart:       make(chan bool),
		gil:           &sync.Once{},
		ID:            ID(),
		context:       ctx,
		cancelFunc:    cancelFunc,
	}
}

func (r *Manager) Start() error {
	w := NewWatcher(r)
	w.Start()
	go r.build(fsnotify.Event{Name: ":start:"})
	// watch files
	go func() {
		log.Println("watching files...")
		for {
			select {
			case event := <-w.Events:
				if event.Op != fsnotify.Chmod {
					go r.build(event)
				}
				w.Remove(event.Name)
				w.Add(event.Name)
			case <-r.context.Done():
				break
			}
		}
	}()

	go func() {
		for {
			select {
			case err := <-w.Errors:
				r.Logger.Error(err)
			case <-r.context.Done():
				break
			}
		}
	}()
	r.runner()
	return nil
}

func (r *Manager) build(event fsnotify.Event) {
	r.gil.Do(func() {

		defer func() {
			r.gil = &sync.Once{}
		}()

		r.buildTransaction(func() error {
			// time.Sleep(r.BuildDelay * time.Millisecond)

			now := time.Now()
			r.Logger.Print("Rebuild on: %s", event.Name)

			args := []string{"build", "-v"}
			args = append(args, r.BuildFlags...)
			args = append(args, "-o", "../"+r.FullBuildPath(), r.BuildTargetPath)
			cmd := exec.Command("go", args...)
			cmd.Dir = "server"
			if err := r.runAndListen(cmd); err != nil {
				if strings.Contains(err.Error(), "no buildable Go source files") {
					r.cancelFunc()
					log.Fatal(err)
				}
				return err
			}

			tt := time.Since(now)
			r.Logger.Success("Building Completed (PID: %d) (Time: %s)", cmd.Process.Pid, tt)
			r.Restart <- true
			return nil
		})
	})
}

func (r *Manager) buildTransaction(fn func() error) {
	lpath := ErrorLogPath()
	err := fn()
	if err != nil {
		f, _ := os.Create(lpath)
		fmt.Fprint(f, err)
		r.Logger.Error("Error!")
		r.Logger.Error(err)
	} else {
		os.Remove(lpath)
	}
}
