package scaleway

import "fmt"

// OrganizationsService handles communication with the account API.
type OrganizationsService struct {
	client *Client
}

// Organization represents a Scaleway organization.
type Organization struct {
	ID    string  `json:"id,omitempty"`
	Name  string  `json:"name,omitempty"`
	Users []*User `json:"users,omitempty"`
}

// organizationListResponse represents a Scaleway organization list response.
type organizationListResponse struct {
	Organizations []*Organization `json:"organizations"`
}

// List returns a list of all organisation associate to your account.
func (s *OrganizationsService) List() ([]*Organization, *Response, error) {
	return s.listOrganizations()
}

func (s *OrganizationsService) listOrganizations() ([]*Organization, *Response, error) {
	u := fmt.Sprintf("/organizations")
	req, err := s.client.NewRequestAccount("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	organizations := new(organizationListResponse)
	resp, err := s.client.Do(req, organizations)
	if err != nil {
		return nil, nil, err
	}

	return organizations.Organizations, resp, nil
}
