package main

import (
	"log"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/mattn/go-slim"
	_ "github.com/mattn/go-sqlite3"
	"github.com/russross/blackfriday"
)

type Page struct {
	Path    string `gorm:"primary_key"`
	Content string `gorm:"not null"`
}

func (p *Page) Title() string {
	title := strings.TrimLeft(p.Path, "/")
	if strings.HasPrefix(p.Content, "#") {
		title = strings.TrimLeft(strings.Split(p.Content, "\n")[0], "#")
	}
	return title
}

func (p *Page) Body() string {
	body := p.Content
	if strings.HasPrefix(p.Content, "#") {
		if pos := strings.Index(p.Content, "\n"); pos >= 0 {
			body = p.Content[pos+1:]
		} else {
			body = ""
		}
	}
	return string(blackfriday.MarkdownBasic([]byte(body)))
}

var (
	tplPage  *slim.Template
	tplEdit  *slim.Template
	tplPages *slim.Template
	db       *gorm.DB
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
	db, err = gorm.Open("sqlite3", "wiki.db")
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&Page{})
}

func edit(c echo.Context) error {
	var page Page
	db.Where("path=?", "/"+c.Param("path")).Find(&page)
	return tplEdit.Execute(c.Response(), map[string]interface{}{
		"title":   page.Title(),
		"path":    "/" + c.Param("path"),
		"content": page.Content,
	})
}

func page(c echo.Context) error {
	var page Page
	db.Where("path=?", "/"+c.Param("path")).Find(&page)
	return tplPage.Execute(c.Response(), map[string]interface{}{
		"title":     page.Title(),
		"path":      c.Param("path"),
		"editpath":  path.Join("/"+c.Param("path"), "edit"),
		"pagespath": "/pages",
		"content":   page.Body(),
	})
}

func pages(c echo.Context) error {
	var pages []Page
	db.Order("path").Find(&pages)
	return tplPages.Execute(c.Response(), map[string]interface{}{
		"pages": pages,
	})
}

func update(c echo.Context) error {
	var page Page
	page.Path = "/" + c.Param("path")
	page.Content = c.FormValue("content")
	db.Save(&page)
	return c.Redirect(http.StatusFound, page.Path)
}

func restricted() []echo.MiddlewareFunc {
	secret := os.Getenv("WIKIGO_AUTH")
	if secret == "" {
		return nil
	}
	return []echo.MiddlewareFunc{
		middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
			if username+":"+password == secret {
				return true, nil
			}
			return false, nil
		}),
	}
}

func main() {
	e := echo.New()
	e.GET("/pages", pages)

	e.GET("/", page)
	e.GET("/:path", page)

	ms := restricted()
	e.GET("/edit", edit, ms...)
	e.POST("/edit", update, ms...)
	e.GET("/:path/edit", edit, ms...)
	e.POST("/:path/edit", update, ms...)

	e.Static("/static", "static")
	e.Start(":8081")
}
