package main

import (
	"bufio"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var files []string

func main() {
	fortuneCommand := exec.Command("fortune", "-f")
	pipe, err := fortuneCommand.StderrPipe()
	if err != nil {
		panic(err)
	}

	fortuneCommand.Start()
	outputStream := bufio.NewScanner(pipe)
	outputStream.Scan()
	line := outputStream.Text()
	root := line[strings.Index(line, "/"):]

	err = filepath.Walk(root, visit)
	if err != nil {
		panic(err)
	}

	println(len(files))
}

func visit(path string, f os.FileInfo, err error) error {
	if strings.Contains(path, "/off/") {
		return nil
	}

	if filepath.Ext(path) == ".dat" {
		return nil
	}

	if f.IsDir() {
		return nil
	}

	files = append(files, path)
	return nil
}
