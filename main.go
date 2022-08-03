package main

import (
	"os"	
	"log"
)

type ConfigFile struct {
	path string
}

func (f ConfigFile) Exists() bool {
	if _, err := os.Stat(f.path); err == nil {
		return true
	}

	return false
}

func (f ConfigFile) Create() {
	if !f.Exists() {
		_, err := os.OpenFile(f.path, os.O_CREATE|os.O_WRONLY, 0644)
    	if err != nil {
        	log.Fatal(err)
    	}
	}
}

func (f ConfigFile) Remove() {
	e := os.Remove(f.path)
    if e != nil {
        log.Fatal(e)
    }
}

func (f ConfigFile) FromHome(path string) ConfigFile {
	dirname, err := os.UserHomeDir()
    if err != nil {
        log.Fatal( err )
    }

	return ConfigFile { dirname + path }
}

type ConfigDir struct {
	path string
}

func (f ConfigDir) Exists() bool {
	if _, err := os.Stat(f.path); err == nil {
		return true
	}

	return false
}

func (f ConfigDir) Create() {
	if !f.Exists() {
		if err := os.Mkdir(f.path, os.ModePerm); err != nil {
        	log.Fatal(err)
    	}
	}
}

func (f ConfigDir) Remove() {
	e := os.Remove(f.path)
    if e != nil {
        log.Fatal(e)
    }
}

func (f ConfigDir) FromHome(path string) ConfigDir {
	dirname, err := os.UserHomeDir()
    if err != nil {
        log.Fatal( err )
    }

	return ConfigDir { dirname + path }
}
