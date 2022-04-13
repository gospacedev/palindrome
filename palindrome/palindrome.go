package palindrome

import "fmt"

func Reverse(s string) (result string) {
	for _,v := range s {
	  result = string(v) + result
	}
	return 
    }

func checkPwnd(s string) (b bool) {
	a := Reverse(s)
	if s == a{
		return true
	}else {
		return false
	}
}
