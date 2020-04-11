package main

import "fmt" // https://godoc.org/fmt

var y = 42

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Printf("%b\n", y)
	fmt.Printf("%x\n", y)
	fmt.Printf("%#x\n", y)
	y = 503
	fmt.Printf("%#x\n", y)
	fmt.Printf("%#x\t%b\t%x\n", y, y, y)

	str := fmt.Sprintf("%#x\t%b\t%x\n", y, y, y)
	fmt.Printf(str)
}
