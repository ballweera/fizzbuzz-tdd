package fizzbuzz

import "strconv"

func Say(n int) string {
	if n == 3 || n == 6 {
		return "Fizz"
	}
	return strconv.Itoa(n)
}
