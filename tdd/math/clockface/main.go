// Writes an SVG clockface of the current time to Stdout.
package main

import (
	"os"
	"time"

	svg "github.com/ancill/go-cookbook/math/svg"
)

func main() {
	t := time.Now()
	svg.Write(os.Stdout, t)
}
