package hello_test

import (
	"testing"
	hello "github.com/Tech-Book-Community/workshop-learn-go-with-tests/01-hello-world/gson"
)

func checkValue(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
func TestHello(t *testing.T) {
	t.Run("Hello, gson", func(t *testing.T) {
		got := hello.Hello("gson")
		want := "Hello, gson"
		checkValue(t, got, want)
	})
	t.Run("Hello, world", func(t *testing.T) {
		got := hello.Hello("")
		want := "Hello, world"
		checkValue(t, got, want)
	})

}
