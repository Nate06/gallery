package controllers

import (
	"gallery/views"
)

type Static struct {
	Home     *views.View
	Contact  *views.View
	NotFound *views.View
}

func NewStatic() *Static {
	return &Static{
		Home:     views.NewView("layout", "static/home"),
		Contact:  views.NewView("layout", "static/contact"),
		NotFound: views.NewView("layout", "404"),
	}
}
