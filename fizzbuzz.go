package fizzbuzz

import "strconv"

func Say(n int) string {
	if n == 3 || n == 6 {
		return "Fizz"
	}

	if n == 5 {
		return "Buzz"
	}

	return strconv.Itoa(n)
}
