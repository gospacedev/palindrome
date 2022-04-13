package palindrome

import "fmt"

//Reverses every character in the string
func Reverse(s string) (result string) {
	for _,v := range s {
	  result = string(v) + result
	}
	return 
    }

//checkPwnd checks if the string is same as it is reversed
func checkPwnd(s string) (b bool) {
	a := Reverse(s)
	if s == a{
		return true
	}else {
		return false
	}
}
