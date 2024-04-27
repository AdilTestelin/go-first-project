package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying Hello to people", func(t *testing.T) {
		got := Hello("Adil", "English")
		want := "Hello Adil !"

		assertCorrectMessage(t, got, want)
	})
	t.Run("say, 'Hello World !' when empty string is supplied", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello World !"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello in spanish", func(t *testing.T) {
		got := Hello("Adil", "Spanish")
		want := "Hola Adil !"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello in french", func(t *testing.T) {
		got := Hello("Adil", "French")
		want := "Bonjour Adil !"

		assertCorrectMessage(t, got, want)
	})

}

func assertCorrectMessage(t testing.TB, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
