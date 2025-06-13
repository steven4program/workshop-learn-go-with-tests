package murky

import (
	"errors"
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(10)

		got := wallet.Balance()
		want := Bitcoin(10)

		assertBalance(t, got, want)
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{Bitcoin(20)}

		err := wallet.Withdraw(10)
		assertNoError(t, err)
		got := wallet.Balance()
		want := Bitcoin(10)

		assertBalance(t, got, want)
	})

	t.Run("Withdraw too much", func(t *testing.T) {
		wallet := Wallet{Bitcoin(10)}

		got := wallet.Withdraw(20)
		want := ErrInsufficientBalance

		assertError(t, got, want)
	})
}

func assertBalance(t testing.TB, got, want Bitcoin) {
	t.Helper()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if !errors.Is(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()

	if got != nil {
		t.Error("I said no err!")
	}
} 