package main

import (
	"io/ioutil"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"

	"../t"
)

//Max and min values for the last 3 entries are irrelevant (float, double and long double) and won't be used
var maxLimits = [...]uint64{math.MaxInt8, math.MaxUint8, math.MaxInt16, math.MaxUint16, math.MaxInt32,
	math.MaxUint32, math.MaxInt64, math.MaxUint64, math.MaxInt64, math.MaxUint64, 0, 0, 0}

var minLimits = [...]int64{math.MinInt8, 0, math.MinInt16, 0, math.MinInt32,
	0, math.MinInt64, 0, math.MinInt64, 0, 0, 0, 0}

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
			Generic         string
			ValGeneric      string
			Types           [13]string
			MaxGeneric      uint64
			MinGeneric      int64
			PriorityGeneric int
		}{s, strconv.Itoa(p), t.TypeKeyStrings, maxLimits[p], minLimits[p], p}
		f, err := os.OpenFile(filepath.Join(build, s+".go"), os.O_RDWR|os.O_CREATE, 0644)
		if err != nil {
			panic(err)
		}
		if err := tmpl.Execute(f, vars); err != nil {
			panic(err)
		}
	}
}
