package main

import (
	"bufio"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
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

	rand.Seed(time.Now().UnixNano())
	i := randomInt(1, len(files))
	randomFile := files[i]
	print(randomFile)
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

// Returns an int >= min, < max
func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}
