# palindrome
A palindrome string checker

## Install

    go get -u github.com/gocrazygth/palindrome

## Example

```go

package main

import (
	"fmt"
	"github.com/gocrazygth/palindrome"
)

func main() {
	a := pal.Check("racecar")
	b := pal.Reverse("drawer")
	fmt.Println(a)
	fmt.Println(b)
}

```

Output:

    true
    reward
