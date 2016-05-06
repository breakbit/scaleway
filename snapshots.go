package scaleway

import "fmt"

// SnapshotsService handles communication with the tokens related
// methods of the Scaleway API.
//
// Scaleway API docs: https://developer.scaleway.com/#snapshots
type SnapshotsService struct {
	client *Client
}

// Snapshot represents a Scaleway snapshot.
type Snapshot struct {
	ID               string  `json:"id,omitempty"`
	Name             string  `json:"name,omitempty"`
	CreationDate     Ntime   `json:"creation_date,omitempty"`
	ModificationDate Ntime   `json:"modification_date,omitempty"`
	Organization     string  `json:"organization,omitempty"`
	Size             uint64  `json:"size,omitempty"`
	State            string  `json:"state,omitempty"`
	Type             string  `json:"volume_type,omitempty"`
	BaseVolume       *Volume `json:"base_volume,omitempty"`
}

// SnapshotRequest represents a request to create a snapshot.
type SnapshotRequest struct {
	Name         string `json:"name,omitempty"`
	Organization string `json:"organization,omitempty"`
	Volume       string `json:"volume_id,omitempty"`
}

// snapshotResponse represents a Scaleway snapshot creation response.
type snapshotResponse struct {
	Snapshot *Snapshot `json:"snapshot"`
}

// snapshotListResponse represents a Scaleway snapshot list response.
type snapshotListResponse struct {
	Snapshots []*Snapshot `json:"snapshots"`
}

// Create create snapshot of a volume.
func (s *SnapshotsService) Create(sr *SnapshotRequest) (*Snapshot, *Response, error) {
	u := fmt.Sprintf("/snapshots")
	req, err := s.client.NewRequestAccount("POST", u, sr)
	if err != nil {
		return nil, nil, err
	}

	snapshot := new(snapshotResponse)
	resp, err := s.client.Do(req, snapshot)
	if err != nil {
		return nil, nil, err
	}
	return snapshot.Snapshot, resp, nil
}

// List returns a list of all snapshots associate to your account.
func (s *SnapshotsService) List() ([]*Snapshot, *Response, error) {
	return s.listSnapshots()
}

func (s *SnapshotsService) listSnapshots() ([]*Snapshot, *Response, error) {
	u := fmt.Sprintf("/snapshots")
	req, err := s.client.NewRequestAccount("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	snapshots := new(snapshotListResponse)
	resp, err := s.client.Do(req, snapshots)
	if err != nil {
		return nil, nil, err
	}
	return snapshots.Snapshots, resp, nil
}

// Get returns info for a specific snapshot.
func (s *SnapshotsService) Get(id string) (*Snapshot, *Response, error) {
	u := fmt.Sprintf("/snapshots/%s", id)
	req, err := s.client.NewRequestAccount("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	snapshot := new(snapshotResponse)
	resp, err := s.client.Do(req, snapshot)
	if err != nil {
		return nil, nil, err
	}
	return snapshot.Snapshot, resp, nil
}

// Update updates the details about snapshot.
func (s *SnapshotsService) Update(id string, sr *SnapshotRequest) (*Snapshot, *Response, error) {
	u := fmt.Sprintf("/snapshots/%s", id)
	req, err := s.client.NewRequestAccount("PUT", u, sr)
	if err != nil {
		return nil, nil, err
	}

	snapshot := new(snapshotResponse)
	resp, err := s.client.Do(req, snapshot)
	if err != nil {
		return nil, nil, err
	}
	return snapshot.Snapshot, resp, nil
}

// Delete deletes a snapshot.
func (s *SnapshotsService) Delete(id string) (*Response, error) {
	u := fmt.Sprintf("/snapshots/%s", id)
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
