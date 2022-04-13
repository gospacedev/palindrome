package pnde

func ifPalindrome(s string) (b bool) {
	a := ""

    Reverse := func(s string) (result string) {
	for _,v := range s {
	  result = string(v) + result
	}
	return 
    }

	a = Reverse(s)
	if s == a{
		return true
	}else {
		return false
	}
}
