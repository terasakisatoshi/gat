package main

import (
	"os"

	"C"
	"github.com/koki-develop/gat/internal/gat"
	"github.com/koki-develop/gat/cmd"
)

// go build -buildmode=c-shared -o export.so

//export cmd_gat
func cmd_gat(file *C.char) *C.char {
	g, err := gat.New(&gat.Config{
		Language:       "",
		Format:         "terminal256",
		Theme:          "monokai",
		RenderMarkdown: false,
		ForceBinary:    false,
	})
	if err != nil {
		return nil
	}
	filename := C.GoString(file)
	f, err := os.Open(filename)
	if err != nil {
		return nil
	}
	defer f.Close()
	if err := g.Print(os.Stdout, f, gat.WithPretty(false), gat.WithFilename(filename)); err != nil {
		return nil
	}

	return nil
}

func main() {
	cmd.Execute()
}
