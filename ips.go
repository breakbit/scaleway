package scaleway

import "fmt"

// IPsService handles communication with the servers related
// methods of the Scaleway API.
//
// Scaleway API docs: https://developer.scaleway.com/#ips
type IPsService struct {
	client *Client
}

// IP represents a Scaleway IP.
type IP struct {
	Address      string  `json:"address,omitempty"`
	ID           string  `json:"id,omitempty"`
	Organization string  `json:"organization,omitempty"`
	Server       *Server `json:"server,omitempty"`
	Reverse      string  `json:"reverse,omitempty"`
}

// IPRequest represents a request to create/attach an IP.
type IPRequest struct {
	Organization string `json:"organization"`
	Address      string `json:"adress,omitempty"`
	ID           string `json:"id,omitempty"`
	Server       string `json:"server,omitempty"`
}

// ipResponse represents a Scaleway IP creation response.
type ipResponse struct {
	IP *IP `json:"ip"`
}

// ipListResponse represents a Scaleway IP list response.
type ipListResponse struct {
	IPs []*IP `json:"ips"`
}

// Create a reserved IP.
func (s *IPsService) Create(ir *IPRequest) (*IP, *Response, error) {
	u := fmt.Sprintf("/ips")
	req, err := s.client.NewRequestCompute("POST", u, vr)
	if err != nil {
		return nil, nil, err
	}

	ip := new(ipResponse)
	resp, err := s.client.Do(req, ip)
	if err != nil {
		return nil, nil, err
	}
	return ip.IP, resp, nil
}

// List returns a list of all reserved IPs.
func (s *IPsService) List() ([]*IP, *Response, error) {
	return s.listIPs()
}

func (s *IPsService) listIPs() ([]*IP, *Response, error) {
	u := fmt.Sprintf("/ips")
	req, err := s.client.NewRequestCompute("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	ips := new(ipListResponse)
	resp, err := s.client.Do(req, ips)
	if err != nil {
		return nil, nil, err
	}
	return ips.IPs, resp, nil
}

// Get returns info for a specific reserved IP.
func (s *IPsService) Get(id string) (*IP, *Response, error) {
	u := fmt.Sprintf("/ips/%s", id)
	req, err := s.client.NewRequestCompute("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	ip := new(ipResponse)
	resp, err := s.client.Do(req, ip)
	if err != nil {
		return nil, nil, err
	}
	return ip.IP, resp, nil
}

// Attach allow you to attach an IP to a server.
func (s *IPsService) Attach(ir *IPRequest, id string) (*IP, *Response, error) {
	u := fmt.Sprintf("/ips/%s", id)
	req, err := s.client.NewRequestCompute("PUT", u, nil)
	if err != nil {
		return nil, nil, err
	}

	ip := new(ipResponse)
	resp, err := s.client.Do(req, ip)
	if err != nil {
		return nil, nil, err
	}
	return ip.IP, resp, nil
}

// Delete deletes a reserved IP.
func (s *IPsService) Delete(id string) (*Response, error) {
	u := fmt.Sprintf("/ips/%s", id)
	req, err := s.client.NewRequestCompute("DELETE", u, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
