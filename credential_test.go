package scaleway

import "testing"

func TestNewCredential(t *testing.T) {
	email, password := "foo@bar.com", "foobar"
	c := NewCredential(email, password)

	if got, want := c.Email, email; got != want {
		t.Errorf("NewCredential Email is %v, want %v", got, want)
	}
	if got, want := c.Password, password; got != want {
		t.Errorf("NewCredential Password is %v, want %v", got, want)
	}
}
