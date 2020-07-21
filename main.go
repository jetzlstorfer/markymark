package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func main() {

	argsWithProg := os.Args
	filename := argsWithProg[1]
	data, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	//fmt.Print(string(data))

	r, _ := regexp.Compile("{{ snippets/([_a-zA-Z0-9./-]*) }}")

	placeholders := r.FindAllString(string(data), -1)

	newContent := string(data)
	for _, placeholder := range placeholders {
		filePath := placeholder
		filePath = strings.Replace(filePath, "{{ ", "", 1)
		filePath = strings.Replace(filePath, " }}", "", 1)
		fmt.Println(filePath)
		filecontent, err := ioutil.ReadFile(filePath)
		if err != nil {
			panic(err)
		}
		newContent = strings.ReplaceAll(newContent, placeholder, string(filecontent))
	}

	newFilename := strings.Replace(filename, ".md", "_gen.md", 1)
	ioutil.WriteFile(newFilename, []byte(newContent), 0644)
	//fmt.Println(newContent)
}
