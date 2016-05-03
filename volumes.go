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
	//TODO @dmelina FIX server (string ?)
	Name         string         `json:"name,omitempty"`
	ID           string         `json:"id,omitempty"`
	ExportURI    string         `json:"export_uri,omitempty"`
	Organization OrganizationID `json:"organization,omitempty"`
	//Server       string         `json:"server,omitempty"`
	Size int    `json:"size,omitempty"`
	Type string `json:"volume_type"`
}

// VolumeRequest represents a request to create a volume.
type VolumeRequest struct {
	Name         string         `json:"name"`
	Organization OrganizationID `json:"organization"`
	Type         string         `json:"volume_type"`
	Size         int            `json:"size"`
}

// VolumeResponse represents a Scaleway volume creation response.
type VolumeResponse struct {
	Volume *Volume `json:"volume"`
}

// Create volume corresponding as data storage for your server.
func (s *VolumesService) Create(vr *VolumeRequest) (*Volume, *Response, error) {
	u := fmt.Sprintf("/volumes")
	req, err := s.client.NewRequestCompute("POST", u, vr)
	if err != nil {
		return nil, nil, err
	}

	volume := new(VolumeResponse)
	resp, err := s.client.Do(req, volume)
	if err != nil {
		return nil, nil, err
	}
	return volume.Volume, resp, nil
}
