package scaleway

// Credential represents a Scaleway login composed by email and password.
type Credential struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// NewCredential returns a new Scaleway Credential need for generating
// the first auth-token.
func NewCredential(email string, password string) *Credential {
	return &Credential{
		Email:    email,
		Password: password,
	}
}
