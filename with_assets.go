// +build static

package main

import (
	"log"

	"github.com/mattn/go-slim"
)

func init() {
	tset := []struct {
		p string
		t **slim.Template
	}{
		{"/view/page.slim", &tplPage},
		{"/view/edit.slim", &tplEdit},
		{"/view/pages.slim", &tplPages},
	}
	for _, t := range tset {
		f, err := Assets.Open(t.p)
		if err != nil {
			log.Fatal(err)
		}
		*(t.t), err = slim.Parse(f)
		if err != nil {
			log.Fatal(err)
		}
		f.Close()
	}
}
