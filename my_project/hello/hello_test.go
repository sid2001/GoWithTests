package hello

import "testing"

func TestHello(t *testing.T) {
	//sub tests t.Run
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("sisa")
		want := "Hello, sisa"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to 'world'", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	//tells that this is a helper function and any error will point to the line which called this function
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
