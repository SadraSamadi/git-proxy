package git

import (
	"errors"
	"strings"
)

var Runner CommandRunner = &DefaultRunner{}

func Configure(value string) error {
	_, err := Runner.Run("git", "config", "--global", "http.proxy", value)
	return err
}

func Current() (string, error) {
	out, err := Runner.Run("git", "config", "--global", "http.proxy")
	if err != nil {
		err = errors.New("no proxy configured")
		return "", err
	}
	out = strings.TrimSpace(out)
	return out, nil
}

func Unset() error {
	_, _ = Runner.Run("git", "config", "--global", "--unset", "http.proxy")
	return nil
}
