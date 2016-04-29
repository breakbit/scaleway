package scaleway

// UsersService handles communication with the account API.
type UsersService struct {
	client *Client
}

// User represents a Scaleway user.
// @TODO @dmelina add role, organization
type User struct {
	Email      string   `json:"email,omitempty"`
	Firstname  string   `json:"firstname,omitempty"`
	Lastname   string   `json:"lastname,omitempty"`
	Fullname   string   `json:"fullname,omitempty"`
	ID         string   `json:"id,omitempty"`
	SSHPubKeys []string `json:"ssh_public_keys,omitempty"`
}
