package blogrenderer

import (
	"embed"
	"github.com/ancill/go-cookbook/blogposts"
	"html/template"
	"io"
	"io/fs"
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

func (r *PostRenderer) Render(w io.Writer, fs fs.FS) error {
	p, err := blogposts.NewPostsFromFS(fs)
	if err != nil {
		return err
	}
	if err := r.templ.ExecuteTemplate(w, "blog.html", struct {
		Posts []blogposts.Post
	}{p}); err != nil {
		return err
	}

	return nil
}
