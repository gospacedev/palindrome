# puede
Checks if a word is a palindrome

## Install

    go get -u github.com/gocrazygh/puede

## Example

```go

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

```

Output:

    true
    reward
