package customsortgo

func ReverseString(word string) string {
	r := []rune(word)
	var rw []rune
	for i := len(r) - 1; i >= 0; i-- {
		rw = append(rw, r[i])
	}
	return string(rw)
}
