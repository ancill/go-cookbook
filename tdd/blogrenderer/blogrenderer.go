package blogrenderer

import (
	"embed"
	"html/template"
	"io"

	blogposts "github.com/ancill/go-cookbook/blogrenderer/blogpost"
)

var (
	//go:embed "templates/*"
	postTemplates embed.FS
)

// templ, err := template.New("blog").Parse(postTemplate)
func Render(w io.Writer, p blogposts.Post) error {
	templ, err := template.ParseFS(postTemplates, "templates/*.html")
	if err != nil {
		return err
	}

	if err := templ.Execute(w, p); err != nil {
		return err
	}

	return nil
}
