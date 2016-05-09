package scaleway

import "fmt"

// ServersService handles communication with the servers related
// methods of the Scaleway API.
//
// Scaleway API docs: https://developer.scaleway.com/#servers
type ServersService struct {
	client *Client
}

// Server represents a Scaleway server.
type Server struct {
	ID              string             `json:"id,omitempty"`
	BootScript      string             `json:"bootscript,omitempty"`
	DynamicPublicIP bool               `json:"dynamic_public_ip,omitempty"`
	Image           *Image             `json:"image,omitempty"`
	Name            string             `json:"name,omitempty"`
	Organization    string             `json:"organization,omitempty"`
	PrivateIP       string             `json:"private_ip,omitempty"`
	PublicIP        string             `json:"public_ip,omitempty"`
	State           string             `json:"state,omitempty"`
	Tags            []string           `json:"tags,omitempty"`
	Volumes         map[string]*Volume `json:"volumes,omitempty`
}

// ServerRequest represents a request to create a server.
type ServerRequest struct {
	Organization string   `json:"organization"`
	Name         string   `json:"name"`
	Image        string   `json:"image"`
	Tags         []string `json:"tags"`
}

// serverResponse represents a Scaleway server creation response.
type serverResponse struct {
	Server *Server `json:"server"`
}

// serverListResponse represents a Scaleway servers list response.
type serverListResponse struct {
	Servers []*Server `json:"servers"`
}

// Create creates images.
func (s *ServersService) Create(sr *ServerRequest) (*Server, *Response, error) {
	u := fmt.Sprintf("/servers")
	req, err := s.client.NewRequestAccount("POST", u, sr)
	if err != nil {
		return nil, nil, err
	}

	server := new(serverResponse)
	resp, err := s.client.Do(req, server)
	if err != nil {
		return nil, nil, err
	}
	return server.Server, resp, nil
}

// List returns a list of all servers associate to your account.
func (s *ServersService) List() ([]*Server, *Response, error) {
	return s.listServers()
}

func (s *ServersService) listServers() ([]*Server, *Response, error) {
	u := fmt.Sprintf("/servers")
	req, err := s.client.NewRequestAccount("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	servers := new(serverListResponse)
	resp, err := s.client.Do(req, servers)
	if err != nil {
		return nil, nil, err
	}
	return servers.Servers, resp, nil
}
