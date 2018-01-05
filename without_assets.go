// +build !static

package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/mattn/go-slim"
)

func init() {
	exe, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	dir := filepath.Dir(exe)
	tset := []struct {
		p string
		t **slim.Template
	}{
		{"/view/page.slim", &tplPage},
		{"/view/edit.slim", &tplEdit},
		{"/view/pages.slim", &tplPages},
	}
	for _, t := range tset {
		*(t.t), err = slim.ParseFile(filepath.Join(dir, t.p))
		if err != nil {
			log.Fatal(err)
		}
	}
}
