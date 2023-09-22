package string

// ReverseStringSlow reverses a string by replace first and last letter with each other
// Starts outward and moves inward. only goes for half the link of the rune slice so it doesn't keep going in a circle
// Credit to Cyrus Javan (https://github.com/CyrusJavan) for this code recommendation
func ReverseStringSlow(word string) string {
	r := []rune(word)
	length := len(r)
	for i := 0; i < length/2; i++ {
		r[i], r[length-1-i] = r[length-1-i], r[i]
	}
	return string(r)
}

// ReverseString is a slightly faster way to reverse a string
// It creates a rune array of the length
func ReverseString(word string) string {
	wordLength := len(word)
	rw := make([]rune, wordLength)
	for i, l := range word {
		rw[wordLength-1-i] = l
	}
	return string(rw)
}
