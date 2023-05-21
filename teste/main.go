package main

import (
	"embed"
	"io/fs"
	"log"
	"strings"
)

//go:embed templates
var templateFS embed.FS

func listEmbedFiles(efs embed.FS, pattern string) ([]string, error) {
	lista := make([]string, 0)
	err := fs.WalkDir(efs, ".", func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}
		if strings.Contains(path, pattern) {
			lista = append(lista, path)
		}
		return nil
	})
	return lista, err
}

func main() {
	lista, err := listEmbedFiles(templateFS, "signup.page.tmpl")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(lista)
}
