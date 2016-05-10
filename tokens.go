package scaleway

import "fmt"

// TokensService handles communication with the tokens related
// methods of the Scaleway API.
//
// Scaleway API docs: https://developer.scaleway.com/#tokens
type TokensService struct {
	client *Client
}

// Token represents a Scaleway auth-token.
type Token struct {
	CreationDate      Ntime    `json:"creation_date,omitempty"`
	Expires           Ntime    `json:"expires,omitempty"`
	ID                string   `json:"id,omitempty"`
	InheritsUserPerms bool     `json:"inherits_user_perms,omitempty"`
	UserID            string   `json:"user_id,omitempty"`
	Permission        []string `json:"permissions"`
}

// TokenRequest represents a request to create a token.
type TokenRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Expires  bool   `json:"expires"`
}

// tokenResponse represents a Scaleway token creation response.
type tokenResponse struct {
	Token *Token `json:"token"`
}

// tokenListResponse represents a Scaleway token list response.
type tokenListResponse struct {
	Tokens []*Token `json:"tokens"`
}

// Create authenticates a user against their username, password, and then
// returns a new Token, which can be used until it expires.
func (s *TokensService) Create(tr *TokenRequest) (*Token, *Response, error) {
	u := fmt.Sprintf("/tokens")
	req, err := s.client.NewRequestAccount("POST", u, tr)
	if err != nil {
		return nil, nil, err
	}

	token := new(tokenResponse)
	resp, err := s.client.Do(req, token)
	if err != nil {
		return nil, nil, err
	}
	return token.Token, resp, nil
}

// List returns a list of all tokens associate to your account.
func (s *TokensService) List() ([]*Token, *Response, error) {
	return s.listTokens()
}

func (s *TokensService) listTokens() ([]*Token, *Response, error) {
	u := fmt.Sprintf("/tokens")
	req, err := s.client.NewRequestAccount("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	tokens := new(tokenListResponse)
	resp, err := s.client.Do(req, tokens)
	if err != nil {
		return nil, nil, err
	}
	return tokens.Tokens, resp, nil
}

// Get returns info for a specific token.
func (s *TokensService) Get(id string) (*Token, *Response, error) {
	u := fmt.Sprintf("/tokens/%s", id)
	req, err := s.client.NewRequestAccount("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	token := new(tokenResponse)
	resp, err := s.client.Do(req, token)
	if err != nil {
		return nil, nil, err
	}
	return token.Token, resp, nil
}

// Update increases token expiration time of 30 minutes.
func (s *TokensService) Update(id string) (*Token, *Response, error) {
	u := fmt.Sprintf("/tokens/%s", id)
	req, err := s.client.NewRequestAccount("PATCH", u, struct{}{})
	if err != nil {
		return nil, nil, err
	}

	token := new(tokenResponse)
	resp, err := s.client.Do(req, token)
	if err != nil {
		return nil, nil, err
	}
	return token.Token, resp, nil
}

// Delete deletes a token.
func (s *TokensService) Delete(id string) (*Response, error) {
	u := fmt.Sprintf("/tokens/%s", id)
	req, err := s.client.NewRequestAccount("DELETE", u, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
