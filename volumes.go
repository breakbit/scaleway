package scaleway

import "fmt"

// VolumesService handles communication with the volumes related
// methods of the Scaleway API.
//
// Scaleway API docs: https://developer.scaleway.com/#volumes-volumes
type VolumesService struct {
	client *Client
}

// Volume represents a Scaleway volume.
type Volume struct {
	Name         string         `json:"name,omitempty"`
	ID           string         `json:"id,omitempty"`
	ExportURI    string         `json:"export_uri,omitempty"`
	Organization OrganizationID `json:"organization,omitempty"`
	//Server       string         `json:"server,omitempty"`
	Size uint64 `json:"size,omitempty"`
	Type string `json:"volume_type"`
}

// VolumeRequest represents a request to create a volume.
type VolumeRequest struct {
	Name         string         `json:"name"`
	Organization OrganizationID `json:"organization"`
	Type         string         `json:"volume_type"`
	Size         int            `json:"size"`
}

// volumeResponse represents a Scaleway volume creation response.
type volumeResponse struct {
	Volume *Volume `json:"volume"`
}

// volumeListResponse represents a Scaleway volume list response.
type volumeListResponse struct {
	Volumes []*Volume `json:"volumes"`
}

// Create volume corresponding as data storage for your server.
func (s *VolumesService) Create(vr *VolumeRequest) (*Volume, *Response, error) {
	u := fmt.Sprintf("/volumes")
	req, err := s.client.NewRequestCompute("POST", u, vr)
	if err != nil {
		return nil, nil, err
	}

	volume := new(volumeResponse)
	resp, err := s.client.Do(req, volume)
	if err != nil {
		return nil, nil, err
	}
	return volume.Volume, resp, nil
}

// List returns a list of all volumes associate to your account.
func (s *VolumesService) List() ([]*Volume, *Response, error) {
	return s.listVolumes()
}

func (s *VolumesService) listVolumes() ([]*Volume, *Response, error) {
	u := fmt.Sprintf("/volumes")
	req, err := s.client.NewRequestAccount("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	volumes := new(volumeListResponse)
	resp, err := s.client.Do(req, volumes)
	if err != nil {
		return nil, nil, err
	}
	return volumes.Volumes, resp, nil
}

// Get returns info for a specific volume.
func (s *VolumesService) Get(id string) (*Volume, *Response, error) {
	u := fmt.Sprintf("/volumes/%s", id)
	req, err := s.client.NewRequestAccount("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	volume := new(volumeResponse)
	resp, err := s.client.Do(req, volume)
	if err != nil {
		return nil, nil, err
	}
	return volume.Volume, resp, nil
}

// Delete deletes a volume.
func (s *VolumesService) Delete(id string) (*Response, error) {
	u := fmt.Sprintf("/volumes/%s", id)
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
