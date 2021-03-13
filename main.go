package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	files, err := ioutil.ReadDir("files")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		newName := strings.Replace(f.Name(), "tayfunguler.org - ", "", 1)
		src := "files/"+f.Name()
		fmt.Println(src)
		dst := "files/"+newName
		os.Rename(src, dst)
	}
}
