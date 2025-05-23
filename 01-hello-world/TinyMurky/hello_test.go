package tinymurky

import "testing"

func Test_Hello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")

		want := "Hello, Chris"

		assertCorrectMessage(t, got, want)

	})

	t.Run("saying 'Hello, World' when an empty string is supplied to people", func(t *testing.T) {
		got := Hello("", "")

		want := "Hello, World"

		assertCorrectMessage(t, got, want)

	})

	t.Run("in spanish", func(t *testing.T) {
		got := Hello("Murky", "Spanish")

		want := "Hola, Murky"

		assertCorrectMessage(t, got, want)

	})

	t.Run("in french", func(t *testing.T) {
		got := Hello("Murky", "French")

		want := "Bonjour, Murky"

		assertCorrectMessage(t, got, want)

	})

}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("expect %q, got %q", want, got)
	}
}