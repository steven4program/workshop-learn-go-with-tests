package tinhello

import "testing"

func TestHello(t *testing.T) {
	t.Run("hello world", func(t *testing.T) {
		got := Hello("","")
		want := "hello, world"

		assertCorrectMessage(t, got, want)
	})
	t.Run("hello mm", func(t *testing.T) {
		got := Hello("mm","spanish")
		want := "hola, mm"

		assertCorrectMessage(t, got, want)
	})
	t.Run("hello nn", func(t *testing.T) {
		got := Hello("nn","french")
		want := "bonjour, nn"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string){
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}