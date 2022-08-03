package main

import (
	"testing"
	"strings"
)

func TestConfigFile(t *testing.T) {
	conf := ConfigFile{"./test.conf"}

	got := conf.Exists()
	want := false

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

	conf.Create()

	got = conf.Exists()
	want = true

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

	conf.Remove()

	got = conf.Exists()
	want = false

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

	conf = ConfigFile{""}.FromHome("./test.conf")

	got = conf.path
	if strings.Contains(got, "home") {
		t.Errorf("wanted home in %v", got)
	}
}

func TestConfigDir(t *testing.T) {
	conf := ConfigDir{"./conf"}

	got := conf.Exists()
	want := false

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

	conf.Create()

	got = conf.Exists()
	want = true

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

	conf.Remove()

	got = conf.Exists()
	want = false

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
