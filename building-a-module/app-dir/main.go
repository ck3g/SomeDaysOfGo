package main

import "github.com/ck3g/SomeDaysOfGo/building-a-module/toolkit"

func main() {
	var tools toolkit.Tools

	tools.CreateDirIfNotExist("./test-dir")
}
