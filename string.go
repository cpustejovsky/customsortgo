package customsortgo

func ReverseString(word string) string {
	r := []rune(word)
	for i := 0; i < len(r)/2; i++ {
		r[i], r[len(r)-1-i] = r[len(r)-1-i], r[i]
	}
	return string(r)
}

func ReverseStringSlow(word string) string {
	r := []rune(word)
	var rw []rune
	for i := len(r) - 1; i >= 0; i-- {
		rw = append(rw, r[i])
	}
	return string(rw)
}

