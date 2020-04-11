package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("Not gonna print")
	case (2 == 3):
		fmt.Println("2 is not equal to 3")
	case 3 == 3:
		fmt.Println("3 is equal to 3")
		fallthrough
	case false:
		fmt.Println("You know that is false, but it fellthrough")
	default:
		fmt.Println("This is default")
	}

	n := 503
	switch n {
	case 1:
		fmt.Println("N is equal to one")
	case 500 + 1:
		fmt.Println("N is equal to 500 + 1")
	case 403, 503, 603:
		fmt.Println("N is either 403, 503, or 603")
	default:
		fmt.Println("Ok. I give up!")
	}
}
