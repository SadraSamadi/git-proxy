package git

import (
	"bytes"
	"errors"
	"os/exec"
	"strings"
)

func Current() (string, error) {
	out, err := run("git.exe", "config", "--global", "http.proxy")
	if err != nil {
		return "", err
	}
	out = strings.TrimSpace(out)
	return out, nil
}

func Configure(value string) error {
	_, err := run("git.exe", "config", "--global", "http.proxy", value)
	return err
}

func Unset() error {
	_, err := run("git", "config", "--global", "--unset", "http.proxy")
	return err
}

func run(name string, args ...string) (string, error) {
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
