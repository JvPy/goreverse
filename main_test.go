package goreverser

import (
	"testing"
)

func TestReverse(t *testing.T) {

	ErrorMessage := func(t *testing.T, got, want string, input string) {
		t.Helper()
		if got != want {
			t.Errorf("[FAILED]\ngot => %s\nwant => %s\ngiven => %v", got, want, input)
		}
	}

	t.Run("Should return the reversed string", func(t *testing.T) {
		inputString := "1 and 2 and 3"

		got := Reverse(inputString)
		want := "3 and 2 and 1"

		ErrorMessage(t, got, want, inputString)
	})

	t.Run("Should return the original string", func(t *testing.T) {
		inputString := "unreversable"

		got := Reverse(inputString)
		want := "unreversable"

		ErrorMessage(t, got, want, inputString)
	})
}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Reverse("3 and 2 and 1")
	}
}
