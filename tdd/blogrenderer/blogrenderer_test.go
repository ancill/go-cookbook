package blogrenderer_test

import (
	"bytes"
	approvals "github.com/approvals/go-approval-tests"
	"testing"
	"testing/fstest"

	"github.com/ancill/go-cookbook/blogrenderer"
)

func TestRender(t *testing.T) {
	const (
		firstBody = `Title: Post 1
Description: Description 1
Tags: tdd, go
---
Hello
World`
		secondBody = `Title: Post 2
Description: Description 2
Tags: rust, borrow-checker
---
A
N
C`
	)
	postRenderer, err := blogrenderer.NewPostRenderer()

	if err != nil {
		t.Fatal(err)
	}

	fs := fstest.MapFS{
		"hello world.md":  {Data: []byte(firstBody)},
		"hello-world2.md": {Data: []byte(secondBody)},
	}

	t.Run("it converts a single post into HTML", func(t *testing.T) {
		buf := bytes.Buffer{}

		if err := postRenderer.Render(&buf, fs); err != nil {
			t.Fatal(err)
		}
		approvals.VerifyString(t, buf.String())
	})
}

//func BenchmarkRender(b *testing.B) {
//	var (
//		aPost = blogposts.Post{
//			Title:       "hello world",
//			Body:        "This is a post",
//			Description: "This is a description",
//			Tags:        []string{"go", "tdd"},
//		}
//	)
//
//	postRenderer, err := blogrenderer.NewPostRenderer()
//
//	if err != nil {
//		b.Fatal(err)
//	}
//
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		err := postRenderer.Render(io.Discard, aPost)
//		if err != nil {
//			log.Fatal(err)
//		}
//	}
//}
