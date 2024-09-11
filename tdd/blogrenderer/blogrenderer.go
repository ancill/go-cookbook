package blogrenderer

import (
	"fmt"
	"io"

	blogposts "github.com/ancill/go-cookbook/blogrenderer/blogpost"
)

func Render(w io.Writer, p blogposts.Post) error {

	_, err := fmt.Fprintf(w, "<h1>%s</h1>", p.Title)
	return err
}
