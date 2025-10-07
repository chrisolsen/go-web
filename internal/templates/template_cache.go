package templates

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// type TemplateCache struct {
//
// }

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
		fmt.Println("Error on read")
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
				fmt.Println("Error adding layout")
				return nil, err
			}

			// add partials
			// GO ParseGlob ** does't search all subfolders :(
			// SEE: https://www.google.com/search?q=golang+template.Glob++wildcard+all+subfolders&client=firefox-b-d&sca_esv=475e326674d7cb46&sxsrf=AE3TifM5HNoCFlA-YIViWA3HWXZMDnHJEQ%3A1759729053379&ei=nVXjaO31Fo6D0PEPo7Gz8Q0&ved=0ahUKEwituILJ7Y6QAxWOATQIHaPYLN4Q4dUDCBI&uact=5&oq=golang+template.Glob++wildcard+all+subfolders&gs_lp=Egxnd3Mtd2l6LXNlcnAiLWdvbGFuZyB0ZW1wbGF0ZS5HbG9iICB3aWxkY2FyZCBhbGwgc3ViZm9sZGVyczIFEAAY7wUyBRAAGO8FMgUQABjvBTIFEAAY7wUyBRAAGO8FSJ17UNATWKV0cAZ4AZABAJgBlgGgAesNqgEDOS45uAEDyAEA-AEBmAIUoAK1DMICChAAGLADGNYEGEfCAggQABiiBBiJBcICCBAAGIAEGKIEmAMAiAYBkAYIkgcEMTQuNqAHrUCyBwM4Lja4B8oLwgcGMi0yLjE4yAfKAQ&sclient=gws-wiz-serp
			/**
			t := template.New("t")
			var filenames []string

			// Walk the "templates" directory and its subdirectories
			err := filepath.WalkDir("templates", func(path string, d os.DirEntry, err error) error {
				if err != nil {
					return err
				}
				// If it's a regular file and ends with .html, add it to our list
				if !d.IsDir() && filepath.Ext(path) == ".html" {
					filenames = append(filenames, path)
				}
				return nil
			})

			if err != nil {
				log.Fatalf("walking directory failed: %v", err)
			}

			// Parse the files found into the template. Note the `...` to pass the slice.
			t, err = t.ParseFiles(filenames...)
			if err != nil {
				log.Fatalf("parsing templates failed: %v", err)
			}

			// Now you can execute your templates
			// ...
			*/
			matchingFiles, err := getMatchingFiles(path, ".partial.html")
			if err != nil {
				return nil, err
			}
			fmt.Printf("FOOOOO: %s => %v\n", path, matchingFiles)
			ts, err = ts.ParseFiles(matchingFiles...)

			if err != nil {
				fmt.Printf("Error adding partials(%s): %v\n", path+"/**/*.partial.html", err)
				return nil, err
			}

			// add the pages file
			ts, err = ts.ParseFiles(pagePath)
			if err != nil {
				fmt.Println("Error adding page")
				return nil, err
			}

			fmt.Printf("cache[%s] = %v", name, ts)
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
