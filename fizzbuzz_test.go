package fizzbuzz_test

import "testing"
import "github.com/ballweera/fizzbuzz-tdd"

func TestFizzBuzzShouldSayOne(t *testing.T) {
	got := fizzbuzz.Say(1)
	want := "1"
	assertError(want, got, t)
}

func TestFizzBuzzShouldSayTwo(t *testing.T) {
	got := fizzbuzz.Say(2)
	want := "2"
	assertError(want, got, t)
}

func TestFizzBuzzShouldSayFizzWhenGetThree(t *testing.T) {
	got := fizzbuzz.Say(3)
	want := "Fizz"
	assertError(want, got, t)
}

func TestFizzBuzzShouldSayFizzWhenGetSix(t *testing.T) {
	got := fizzbuzz.Say(6)
	want := "Fizz"
	assertError(want, got, t)
}

func TestFizzBuzzShouldSayBuzzWhenGetFive(t *testing.T) {
	got := fizzbuzz.Say(5)
	want := "Buzz"
	assertError(want, got, t)
}

func assertError(want, got string, t *testing.T) {
	if want != got {
		t.Errorf("it should say %s but got %s", want, got)
	}
}
