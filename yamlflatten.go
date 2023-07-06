package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/M"
	"github.com/kokizzu/gotro/S"
	"gopkg.in/yaml.v3"
)

var Q = func(s string) string {
	return fmt.Sprintf("%q", s)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: yamlflatten <file>")
		os.Exit(1)
	}
	ifile := os.Args[1]

	// read file
	f, err := os.Open(ifile)
	if L.IsError(err, `os.Open: `+ifile) {
		return
	}

	byt, err := io.ReadAll(f)
	if L.IsError(err, `io.ReadAll`) {
		return
	}

	m := M.SX{}

	err = yaml.Unmarshal(byt, m)
	if L.IsError(err, `yaml.Unmarshal`) {
		return
	}

	recurseMap(m, ``)
}

func recurseMap(m M.SX, s string) {
	for k, v := range m {
		nextParent := S.IfElse(s == ``, k, s+`.`+k)
		printOne(v, nextParent)
	}
}

func printOne(v interface{}, nextParent string) {
	if m2, isMap := v.(M.SX); isMap {
		recurseMap(m2, nextParent)
	} else if m2, isMap := v.(map[string]any); isMap {
		recurseMap(m2, nextParent)
	} else if b, isBool := v.(bool); isBool {
		fmt.Println(nextParent + ` = ` + fmt.Sprint(b))
	} else if s, isString := v.(string); isString {
		fmt.Println(nextParent + ` = ` + Q(s))
	} else if arr, isArr := v.([]any); isArr {
		if len(arr) == 0 {
			fmt.Println(nextParent + ` = ` + fmt.Sprint(v))
			return
		}
		recuseArray(arr, nextParent)
	} else if i, isInt := v.(int); isInt {
		fmt.Println(nextParent + ` = ` + fmt.Sprint(i))
	} else if i, isInt := v.(int64); isInt {
		fmt.Println(nextParent + ` = ` + fmt.Sprint(i))
	} else if f, isFloat := v.(float64); isFloat {
		fmt.Println(nextParent + ` = ` + fmt.Sprint(f))
	} else if v == nil {
		fmt.Println(nextParent + ` = null`)
	} else {
		log.Fatalf(`unknown type: %v: %T`, nextParent, v)
	}
}

func recuseArray(arr []any, oldParent string) {
	for i, v := range arr {
		nextParent := oldParent + `.` + fmt.Sprint(i)
		printOne(v, nextParent)
	}
}
