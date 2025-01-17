package main

import (
	"log"
	"os"

	"github.com/ancill/go-cookbook/blogposts"
)

func main() {
	posts, err := blogposts.NewPostsFromFS(os.DirFS("posts"))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(posts)
}
