package customsortgo


//ReverseString reverses a string by replace first and last letter with each other
//Starts outward and moves inward. only goes for half the link of the rune slice so it doesn't keep going in a circle
//Credit to Cyrus Javan (https://github.com/CyrusJavan) for this code recommendation
func ReverseString(word string) string {
	r := []rune(word)
	for i := 0; i < len(r)/2; i++ {
		r[i], r[len(r)-1-i] = r[len(r)-1-i], r[i]
	}
	return string(r)
}

//ReverseStringSlow is a slightly slower way to reverse a string by creating a second rune slice 
func ReverseStringSlow(word string) string {
	r := []rune(word)
	var rw []rune
	for i := len(r) - 1; i >= 0; i-- {
		rw = append(rw, r[i])
	}
	return string(rw)
}

