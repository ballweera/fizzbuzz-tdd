package fizzbuzz_test

import "testing"
import "github.com/ballweera/fizzbuzz-tdd"

func TestFizzBuzzShouldSayOne(t *testing.T) {
	got := fizzbuzz.Say(1)
	want := "1"

	if want != got {
		t.Errorf("it should say %s but got %s", want, got)
	}
}

func TestFizzBuzzShouldSayTwo(t *testing.T) {
	got := fizzbuzz.Say(2)
	want := "2"

	if want != got {
		t.Errorf("it should say %s but got %s", want, got)
	}
}

func TestFizzBuzzShouldSayFizz(t *testing.T) {
	got := fizzbuzz.Say(3)
	want := "Fizz"

	if want != got {
		t.Errorf("it should say %s but got %s", want, got)
	}
}
