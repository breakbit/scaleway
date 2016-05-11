package scaleway

import "testing"

func TestNewCredentials(t *testing.T) {
	email, password := "foo@bar.com", "foobar"
	c := NewCredentials(email, password)

	if got, want := c.Email, email; got != want {
		t.Errorf("NewCredentials Email is %v, want %v", got, want)
	}
	if got, want := c.Password, password; got != want {
		t.Errorf("NewCredentials Password is %v, want %v", got, want)
	}
}
