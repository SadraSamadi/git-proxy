package test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/SadraSamadi/git-proxy/internal/store"
)

const proxy = "http://127.0.0.1:1080"

type MockAdapter struct {
	bytes []byte
	err   error
}

var _ store.ConfigAdapter = (*MockAdapter)(nil)

func (a *MockAdapter) Read() ([]byte, error) {
	return a.bytes, a.err
}

func (a *MockAdapter) Write(bytes []byte) error {
	a.bytes = bytes
	return a.err
}

func TestReadSuccess(t *testing.T) {
	json := fmt.Sprintf(`{"proxies":{"k":"%s"}}`, proxy)
	mockAdaptor(t, json, nil)
	config, err := store.Read()
	if err != nil {
		t.Error(err)
	}
	if v, ok := config.Proxies["k"]; !ok || v != proxy {
		t.Errorf("want k=%s, got k=%s", proxy, v)
	}
}

func TestReadFailure(t *testing.T) {
	oops := errors.New("oops")
	mockAdaptor(t, "", oops)
	config, err := store.Read()
	if err == nil {
		t.Error("expected error")
	}
	if config != nil {
		t.Error("config should be nil")
	}
}

func TestWriteSuccess(t *testing.T) {
	mockAdaptor(t, "", nil)
	config := store.DefaultConfig()
	err := store.Write(config)
	if err != nil {
		t.Error(err)
	}
}

func TestWriteFailure(t *testing.T) {
	oops := errors.New("oops")
	mockAdaptor(t, "", oops)
	err := store.Write(nil)
	if err == nil {
		t.Error("expected error")
	}
}

func mockAdaptor(t *testing.T, json string, err error) {
	original := store.Adapter
	t.Cleanup(func() { store.Adapter = original })
	store.Adapter = &MockAdapter{[]byte(json), err}
}
