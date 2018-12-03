package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"

	"../t"
)

func copyFile(src, dst string) error {
	inp, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	}
	if err = ioutil.WriteFile(dst, inp, 0644); err != nil {
		return err
	}
	return nil
}

func main() {
	var names []string
	var err error

	build := filepath.Join("..", "build", "t")
	if err = os.MkdirAll(build, 0755); err != nil {
		panic(err)
	}

	if names, err = filepath.Glob(filepath.Join("..", "t", "*.go")); err != nil {
		panic(err)
	}

	for _, n := range names {
		dst := filepath.Join("..", "build", "t", filepath.Base(n))
		if err := copyFile(n, dst); err != nil {
			panic(err)
		}
	}

	tmpl := template.New("generic.go.tpl").Funcs(map[string]interface{}{
		"ToUpper": strings.ToUpper,
	})

	if tmpl, err = tmpl.ParseFiles("../t/generic.go.tpl"); err != nil {
		panic(err)
	}

	for p, s := range t.TypeKeyStrings {
		vars := struct {
			Generic    string
			ValGeneric string
			Types      [13]string
		}{s, strconv.Itoa(p), t.TypeKeyStrings}
		f, err := os.OpenFile(filepath.Join(build, s+".go"), os.O_RDWR|os.O_CREATE, 0644)
		if err != nil {
			panic(err)
		}
		if err := tmpl.Execute(f, vars); err != nil {
			panic(err)
		}
	}
}
