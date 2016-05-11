package scaleway

// Credentials represents a Scaleway login composed by email and password.
type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// NewCredentials returns a new Scaleway Credential need for generating
// the first auth-token.
func NewCredentials(email string, password string) *Credentials {
	return &Credentials{
		Email:    email,
		Password: password,
	}
}
