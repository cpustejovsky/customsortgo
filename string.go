package customsortgo

func ReverseString(word string) string {
	var r []rune
	for _, v := range word {
		r = append(r, v)
	}
	var rw []rune
	for i := len(r) - 1; i >= 0; i-- {
		rw = append(rw, r[i])
	}
	return string(rw)
}
