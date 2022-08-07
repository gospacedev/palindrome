package pal

//Reverses every character in the string
func Reverse(s string) (result string){
	for _,v := range s {//Adding the string backwards
	  result = string(v) + result
	}
	return 
    }

//Check if the string is same as it is reversed
func Check(s string) (b bool){// Checks the string when it's the same when reversed
	a := Reverse(s)
	if s == a{
		b = true
	}else {
		b = false
	}
	return
}
