package templates

import (
	"html/template"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type TemplateCache map[string][]string

func humanDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:04")
}

var functions = template.FuncMap{
	"humanDate": humanDate,
}

// Scans path for all tmpl files. Each template file
func NewTemplateCache(path string) (map[string]*template.Template, error) {
	// get a list of all the tmpl files that don't start with underscore
	pagePaths, err := filepath.Glob(path + "/**/*.page.html")
	cache := make(map[string]*template.Template)

	if err != nil {
		log.Println("Error on read")
		return nil, err
	}

	for _, pagePath := range pagePaths {
		parts := strings.Split(filepath.Dir(pagePath), "/")

		// traverse up the dir chain to find nearest _lahout file
		for i := len(parts); i >= 1; i-- {
			parts = parts[:i]
			layoutPath := strings.Join(parts, "/") + "/_layout.tmpl.html"

			_, err := os.Stat(layoutPath)
			if os.IsNotExist(err) {
				continue
			}

			name := filepath.Base(pagePath)

			// init the template
			ts, err := template.New(name).Funcs(functions).ParseFiles(layoutPath)
			if err != nil {
				log.Println("Error adding layout")
				return nil, err
			}

			// add partials
			matchingFiles, err := getMatchingFiles(path, ".partial.html")
			if err != nil {
				return nil, err
			}
			ts, err = ts.ParseFiles(matchingFiles...)

			if err != nil {
				log.Printf("Error adding partials(%s): %v\n", path+"/**/*.partial.html", err)
				return nil, err
			}

			// add the pages file
			ts, err = ts.ParseFiles(pagePath)
			if err != nil {
				log.Println("Error adding page")
				return nil, err
			}

			cache[name] = ts
		}
	}

	return cache, nil
}

func getMatchingFiles(src, extension string) ([]string, error) {
	var filenames []string

	err := filepath.WalkDir(src, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && strings.HasSuffix(path, extension) {
			filenames = append(filenames, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return filenames, nil
}
