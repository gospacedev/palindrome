package main

import "fmt"

func Reverse(s string) (result string) {
	for _,v := range s {
	  result = string(v) + result
	}
	return 
    }

func ifPalindrome(s string) (b bool) {
	a := Reverse(s)
	if s == a{
		return true
	}else {
		return false
	}
}

func main(){
	a := ifPalindrome("racecar")
	fmt.Println(a)
}
