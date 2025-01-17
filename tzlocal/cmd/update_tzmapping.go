package main

import (
	"os"
	"path/filepath"

	"github.com/davidgiffin/go-timezone-local/tzdata"
)

func main() {
	path, _ := filepath.Abs("./tzmapping.go")
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, err = f.WriteString("// Code generated by tzlocal/update_tzmapping.go DO NOT EDIT.\n")
	if err != nil {
		panic(err)
	}
	_, err = f.WriteString("package tzlocal\n\n")
	if err != nil {
		panic(err)
	}
	tzdata.UpdateWindowsTZMapping(f)
}
