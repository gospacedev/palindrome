package main

import (
	"fmt"
	"github.com/gocrazygh/puede"
)

func main() {
	a := puede.Check("racecar")
	b := puede.Reverse("drawer")
	fmt.Println(a)
	fmt.Println(b)
}
