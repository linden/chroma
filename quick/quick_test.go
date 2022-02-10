package quick_test

import (
	"log"
	"os"
	"testing"

	"github.com/linden/chroma/v2/quick"
)

func TestHighlight(t *testing.T) {
	code := `package main

func main() { }
`
	err := quick.Highlight(os.Stdout, code, "go", "html", "monokai")
	if err != nil {
		log.Fatal(err)
	}
}

func TestClarityHighlight(t *testing.T) {
	code := `(define-read-only (example uint)
	(begin
		;; print something
		(print example)
		
		;; example
		(asserts! (is-eq 1 1) (err 1))
		
		;; return something
		(ok some("example"))
	)
)

(example 100)
`

	file, err := os.OpenFile("output.html", os.O_RDWR|os.O_CREATE, 0755)

	if err != nil {
		panic(err)
	}

	err = quick.Highlight(file, code, "clarity", "html", "monokai")

	if err != nil {
		log.Fatal(err)
	}
}
