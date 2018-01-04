// +build !static

package main

import (
	"log"

	"github.com/mattn/go-slim"
)

func init() {
	var err error
	tplPage, err = slim.ParseFile("view/page.slim")
	if err != nil {
		log.Fatal(err)
	}
	tplEdit, err = slim.ParseFile("view/edit.slim")
	if err != nil {
		log.Fatal(err)
	}
	tplPages, err = slim.ParseFile("view/pages.slim")
	if err != nil {
		log.Fatal(err)
	}
}
