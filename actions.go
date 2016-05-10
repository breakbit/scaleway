package scaleway

import "fmt"

// ServersService handles communication with the servers related
// methods of the Scaleway API.
//
// Scaleway API docs: https://developer.scaleway.com/#servers-actions
type ActionsService struct {
	client *Client
}

// ActionRequest represents a request to execute action for a server.
type ActionRequest struct {
	Action string `json:"action"`
}

// actionResponse represents a Scaleway action response.
type actionResponse struct {
	Task *Task `json:"task"`
}

// Task represents the return for an action.
type Task struct {
	Description string `json:"description,omitempty"`
	HrefFrom    string `json:"href_from,omitempty"`
	ID          string `json:"id,omitempty"`
	Progress    string `json:"progress,omitempty"`
	Status      string `json:"status,omitempty"`
}

// actionResponse represents a Scaleway actions list response.
type actionListResponse struct {
	Actions []string `json:"actions,omitempty"`
}

// Exec executes action for a specific server.
func (s *ActionsService) Exec(id string, ar *ActionRequest) (*Task, *Response, error) {
	u := fmt.Sprintf("/servers/%s/action", id)
	req, err := s.client.NewRequestAccount("POST", u, ar)
	if err != nil {
		return nil, nil, err
	}

	action := new(actionResponse)
	resp, err := s.client.Do(req, action)
	if err != nil {
		return nil, nil, err
	}
	return action.Task, resp, nil
}

// List returns a list of all actions associate to your account.
func (s *ActionsService) List(id string) ([]string, *Response, error) {
	return s.listActions(id)
}

func (s *ActionsService) listActions(id string) ([]string, *Response, error) {
	u := fmt.Sprintf("/servers/%s/action", id)
	req, err := s.client.NewRequestAccount("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	actions := new(actionListResponse)
	resp, err := s.client.Do(req, actions)
	if err != nil {
		return nil, nil, err
	}
	return actions.Actions, resp, nil
}
