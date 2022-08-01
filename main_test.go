package main

import "testing"

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
