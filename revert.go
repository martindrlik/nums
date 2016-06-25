package main

func Revert(s string) string {
	r := []rune(s)
	p := len(r) - 1
	q := len(r) / 2
	for i := 0; i < q; i++ {
		r[i], r[p-i] = r[p-i], r[i]
	}
	return string(r)
}
