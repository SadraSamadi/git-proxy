package git

import (
	"bytes"
	"errors"
	"os/exec"
	"strings"
)

type CommandRunner interface {
	Run(name string, args ...string) (string, error)
}

type DefaultRunner struct{}

var _ CommandRunner = (*DefaultRunner)(nil)

func (r *DefaultRunner) Run(name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		msg := stderr.String()
		msg = strings.TrimSpace(msg)
		if msg != "" {
			err = errors.New(msg)
			return "", err
		}
		return "", err
	}
	out := stdout.String()
	return out, nil
}
