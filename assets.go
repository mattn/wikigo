// +build static

package main

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assets9d466d949d40b190143fcf9cd4a17650f4abcc7c = "doctype 5\nhtml\n  head\n    meta charset=\"UTF-8\"\n    link rel=stylesheet href=/static/style.css\n    title history\n  body\n    div#content\n      - for page in pages\n        a href=\"#{page.Path}\" = page.Title()\n        br\n"
var _Assetsd6140ea6f0aaf72598310b1bdb1231a8bc0282dd = "body {\n\tfont-family: Meiryo\n}\nh1.title {\n  text-decoration: underline;\n}\ntextarea {\n  width: 99%;\n  height: 200px;\n}\n"
var _Assets074224e84e52a9655f3d36b70c9f796af41c09e8 = "doctype 5\nhtml\n  head\n    meta charset=\"UTF-8\"\n    link rel=stylesheet href=/static/style.css\n    title = title\n  body\n    div#content\n      h1.title = title\n      form method=post\n        textarea name=content = content\n        br\n        input type=submit\n"
var _Assetscdcf4ddc85bd38d29da2616856db88c64884a6ed = "doctype 5\nhtml\n  head\n    meta charset=\"UTF-8\"\n    link rel=stylesheet href=/static/style.css\n    title = title\n  body\n    div#content\n      h1.title = title\n      div.content = content\n      a href=\"#{editpath}\" edit\n      a href=\"#{pagespath}\" pages\n"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/view": []string{"edit.slim", "page.slim", "pages.slim"}, "/": []string{"static", "view"}, "/static": []string{"style.css"}}, map[string]*assets.File{
	"/view/pages.slim": &assets.File{
		Path:     "/view/pages.slim",
		FileMode: 0x1b6,
		Mtime:    time.Unix(1515069673, 1515069673476859600),
		Data:     []byte(_Assets9d466d949d40b190143fcf9cd4a17650f4abcc7c),
	}, "/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ff,
		Mtime:    time.Unix(1515070902, 1515070902653061200),
		Data:     nil,
	}, "/static": &assets.File{
		Path:     "/static",
		FileMode: 0x800001ff,
		Mtime:    time.Unix(1514901119, 1514901119925036400),
		Data:     nil,
	}, "/static/style.css": &assets.File{
		Path:     "/static/style.css",
		FileMode: 0x1b6,
		Mtime:    time.Unix(1514899811, 1514899811034606500),
		Data:     []byte(_Assetsd6140ea6f0aaf72598310b1bdb1231a8bc0282dd),
	}, "/view": &assets.File{
		Path:     "/view",
		FileMode: 0x800001ff,
		Mtime:    time.Unix(1515069673, 1515069673477836500),
		Data:     nil,
	}, "/view/edit.slim": &assets.File{
		Path:     "/view/edit.slim",
		FileMode: 0x1b6,
		Mtime:    time.Unix(1514897675, 1514897675408947800),
		Data:     []byte(_Assets074224e84e52a9655f3d36b70c9f796af41c09e8),
	}, "/view/page.slim": &assets.File{
		Path:     "/view/page.slim",
		FileMode: 0x1b6,
		Mtime:    time.Unix(1515069595, 1515069595748703000),
		Data:     []byte(_Assetscdcf4ddc85bd38d29da2616856db88c64884a6ed),
	}}, "")
