package test

import (
	"errors"
	"testing"

	"github.com/SadraSamadi/git-proxy/internal/git"
)

type MockRunner struct {
	out string
	err error
}

var _ git.CommandRunner = (*MockRunner)(nil)

func (r *MockRunner) Run(_ string, _ ...string) (string, error) {
	return r.out, r.err
}

func TestConfigureSuccess(t *testing.T) {
	mockRunner(t, "", nil)
	err := git.Configure(proxy)
	if err != nil {
		t.Error(err)
	}
}

func TestConfigureFailure(t *testing.T) {
	oops := errors.New("oops")
	mockRunner(t, "", oops)
	err := git.Configure(proxy)
	if err == nil {
		t.Error("expected error")
	}
}

func TestCurrentSuccess(t *testing.T) {
	mockRunner(t, proxy, nil)
	out, err := git.Current()
	if err != nil {
		t.Error(err)
	}
	if out != proxy {
		t.Errorf("want %s, got %s", proxy, out)
	}
}

func TestCurrentFailure(t *testing.T) {
	oops := errors.New("oops")
	mockRunner(t, "", oops)
	out, err := git.Current()
	if err == nil {
		t.Error("expected error")
	}
	if out != "" {
		t.Errorf("want <empty>, got %s", out)
	}
}

func TestUnsetSuccess(t *testing.T) {
	mockRunner(t, "", nil)
	err := git.Unset()
	if err != nil {
		t.Error(err)
	}
}

func TestUnsetFailure(t *testing.T) {
	oops := errors.New("oops")
	mockRunner(t, "", oops)
	err := git.Unset()
	if err != nil {
		t.Error(err)
	}
}

func mockRunner(t *testing.T, out string, err error) {
	original := git.Runner
	t.Cleanup(func() { git.Runner = original })
	git.Runner = &MockRunner{out, err}
}
