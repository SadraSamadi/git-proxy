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

var _ store.DataAdapter = (*MockAdapter)(nil)

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
	data, err := store.Read()
	if err != nil {
		t.Error(err)
	}
	if v, ok := data.Proxies["k"]; !ok || v != proxy {
		t.Errorf("want k=%s, got k=%s", proxy, v)
	}
}

func TestReadFailure(t *testing.T) {
	oops := errors.New("oops")
	mockAdaptor(t, "", oops)
	data, err := store.Read()
	if err == nil {
		t.Error("expected error")
	}
	if data != nil {
		t.Error("data should be nil")
	}
}

func TestWriteSuccess(t *testing.T) {
	mockAdaptor(t, "", nil)
	data := store.DefaultData()
	err := store.Write(data)
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
	original := store.Adaptor
	t.Cleanup(func() { store.Adaptor = original })
	store.Adaptor = &MockAdapter{[]byte(json), err}
}
