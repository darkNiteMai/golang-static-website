// main.go
package main

import (
	"bytes"
	"flag"
	"fmt"
	"html/template"
	"io/fs"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/renderer/html"
)

// Page holds metadata and content
type Page struct {
	Title   string
	Content template.HTML
}

func main() {
	contentDir := flag.String("content", "content", "Content directory")
	tmplDir := flag.String("templates", "templates", "Templates directory")
	staticDir := flag.String("static", "static", "Static assets directory")
	outDir := flag.String("out", "public", "Output directory")
	flag.Parse()

	// Serve the public directory
	fmt.Println("Serving site at http://localhost:8080")
	http.Handle("/", http.FileServer(http.Dir(*outDir)))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}

	// Prepare output folders
	err := os.RemoveAll(*outDir)
	if err != nil {
		panic(err)
	}
	err = os.MkdirAll(*outDir, os.ModePerm)
	if err != nil {
		panic(err)
	}

	// Copy static assets
	copyStatic(*staticDir, *outDir)

	// Parse templates
	tmpl := template.Must(template.ParseGlob(filepath.Join(*tmplDir, "*.html")))

	// Markdown parser
	md := goldmark.New(
		goldmark.WithExtensions(extension.GFM),
		goldmark.WithRendererOptions(html.WithUnsafe()),
	)

	// Walk content
	filepath.WalkDir(*contentDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return err
		}
		if filepath.Ext(path) != ".md" {
			return nil
		}

		// Read markdown file
		input, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		// Render markdown to HTML
		var buf bytes.Buffer
		err = md.Convert(input, &buf)
		if err != nil {
			return err
		}

		// Determine slug and output path
		rel, err := filepath.Rel(*contentDir, path)
		if err != nil {
			return err
		}
		slug := rel[:len(rel)-len(filepath.Ext(rel))] + ".html"
		outPath := filepath.Join(*outDir, slug)

		// Ensure directory
		err = os.MkdirAll(filepath.Dir(outPath), os.ModePerm)
		if err != nil {
			return err
		}

		// Execute template
		page := Page{
			Title:   deriveTitle(rel),
			Content: template.HTML(buf.String()),
		}
		f, err := os.Create(outPath)
		if err != nil {
			return err
		}
		defer f.Close()
		err = tmpl.ExecuteTemplate(f, "base.html", page)
		if err != nil {
			return err
		}

		fmt.Printf("Generated %s\n", outPath)
		return nil
	})
}

// deriveTitle produces a title from filename
func deriveTitle(path string) string {
	name := filepath.Base(path)
	title := name[:len(name)-len(filepath.Ext(name))]
	return title
}

// copyStatic copies static assets recursively
func copyStatic(src, dstBase string) {
	filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		target := filepath.Join(dstBase, rel)
		if info.IsDir() {
			return os.MkdirAll(target, os.ModePerm)
		}
		data, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		return ioutil.WriteFile(target, data, info.Mode())
	})
}
