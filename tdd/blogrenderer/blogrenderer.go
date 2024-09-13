package blogrenderer

import (
	"embed"
	blogposts "github.com/ancill/go-cookbook/blogrenderer/blogpost"
	"html/template"
	"io"
)

var (
	//go:embed "templates/*"
	postTemplates embed.FS
)

type PostRenderer struct {
	templ *template.Template
}

func NewPostRenderer() (*PostRenderer, error) {
	templ, err := template.ParseFS(postTemplates, "templates/*.html")
	if err != nil {
		return nil, err
	}

	return &PostRenderer{templ: templ}, nil
}

func (r *PostRenderer) Render(w io.Writer, p blogposts.Post) error {
	if err := r.templ.ExecuteTemplate(w, "blog.html", p); err != nil {
		return err
	}

	return nil
}
