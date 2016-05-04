package scaleway

import "fmt"

// UsersService handles communication with the account API.
type UsersService struct {
	client *Client
}

// User represents a Scaleway user.
// TODO @dmelina add role ? and organization ?
type User struct {
	Email      string   `json:"email,omitempty"`
	Firstname  string   `json:"firstname,omitempty"`
	Lastname   string   `json:"lastname,omitempty"`
	Fullname   string   `json:"fullname,omitempty"`
	ID         string   `json:"id,omitempty"`
	SSHPubKeys []string `json:"ssh_public_keys,omitempty"`
}

// userResponse represents a Scaleway token creation response.
type userResponse struct {
	User *User `json:"user"`
}

// Get returns info for a specific user.
func (s *UsersService) Get(id string) (*User, *Response, error) {
	u := fmt.Sprintf("/users/%s", id)
	req, err := s.client.NewRequestAccount("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	user := new(userResponse)
	resp, err := s.client.Do(req, user)
	if err != nil {
		return nil, nil, err
	}
	return user.User, resp, nil
}
