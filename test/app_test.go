package test

import (
	"fmt"
	"testing"

	"github.com/SadraSamadi/git-proxy/internal/app"
	"github.com/SadraSamadi/git-proxy/internal/store"
)

var sample = fmt.Sprintf(`{"proxies":{"k":"%s"}}`, proxy)

func TestValidation(t *testing.T) {
	proxies := map[string]bool{
		proxy:                      true,
		"socks5://127.0.0.1:1080":  true,
		"socks5h://127.0.0.1:1080": true,
		"invalid":                  false,
	}
	for proxy, expected := range proxies {
		if store.Validate(proxy) != expected {
			t.Errorf("unexpected proxy %s", proxy)
		}
	}
}

func TestSave(t *testing.T) {
	mockAdaptor(t, `{"proxies":{}}`, nil)
	err := app.Save("k", proxy)
	if err != nil {
		t.Fatal(err)
	}
	data, err := store.Read()
	if err != nil {
		t.Fatal(err)
	}
	if v, ok := data.Proxies["k"]; !ok || v != proxy {
		t.Errorf("want k=%s, got k=%s", proxy, v)
	}
}

func TestList(t *testing.T) {
	mockAdaptor(t, sample, nil)
	proxies, err := app.List()
	if err != nil {
		t.Fatal(err)
	}
	if v, ok := proxies["k"]; !ok || v != proxy {
		t.Errorf("want k=%s, got k=%s", proxy, v)
	}
}

func TestRemove(t *testing.T) {
	mockAdaptor(t, sample, nil)
	err := app.Remove("k")
	if err != nil {
		t.Fatal(err)
	}
	data, err := store.Read()
	if err != nil {
		t.Fatal(err)
	}
	if v, ok := data.Proxies["k"]; ok || v != "" {
		t.Errorf("want k=<empty>, got k=%s", v)
	}
}

func TestUse(t *testing.T) {
	mockAdaptor(t, sample, nil)
	mockRunner(t, "", nil)
	err := app.Use("k")
	if err != nil {
		t.Fatal(err)
	}
}

func TestCurrent(t *testing.T) {
	mockRunner(t, proxy, nil)
	value, err := app.Current()
	if err != nil {
		t.Fatal(err)
	}
	if value != proxy {
		t.Errorf("want %s, got %s", proxy, value)
	}
}

func TestUnset(t *testing.T) {
	mockRunner(t, "", nil)
	err := app.Unset()
	if err != nil {
		t.Fatal(err)
	}
}
