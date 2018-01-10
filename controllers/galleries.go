package controllers

import (
	"gallery/views"
	"net/http"
)

type Galleries struct {
	NewView *views.View
}

type GalleryForm struct {
	Name   string `schema:"name"`
	Author string `schema:"author"`
}

func NewGalleries() *Galleries {
	return &Galleries{
		NewView: views.NewView("layout", "galleries/new"),
	}
}

func (g *Galleries) New(w http.ResponseWriter, r *http.Request) {
	if err := g.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}
