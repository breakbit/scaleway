package scaleway

import "fmt"

// ImagesService handles communication with the images related
// methods of the Scaleway API.
//
// Scaleway API docs: https://developer.scaleway.com/#images
type ImagesService struct {
	client *Client
}

// Images represents a Scaleway images.
type Image struct {
	CreationDate     Ntime    `json:"creation_date,omitempty"`
	ModificationDate Ntime    `json:"modification_date,omitempty"`
	Arch             string   `json:"arch,omitempty"`
	ExtraVolumes     []string `json:"extra_volumes,omitempty"`
	FromImage        string   `json:"from_image,omitempty"`
	FromServer       string   `json:"from_server,omitempty"`
	ID               string   `json:"id,omitempty"`
	MarketplaceKey   string   `json:"markeplace_key,omitempty"`
	Name             string   `json:"name,omitempty"`
	Organization     string   `json:"organization,omitempty"`
	Public           bool     `json:"public,omitempty"`
	RootVolume       *Volume  `json:"root_volume,omitempty"`
}

// ImageRequest represents a request to create a image.
type ImageRequest struct {
	Organization string `json:"organization"`
	Name         string `json:"name"`
	Arch         string `json:"arch"`
	RootVolume   string `json:"root_volume"`
}

// imageResponse represents a Scaleway image creation response.
type imageResponse struct {
	Image *Image `json:"image"`
}

// imageListResponse represents a Scaleway images list response.
type imageListResponse struct {
	Images []*Image `json:"images"`
}

// Create creates images.
func (s *ImagesService) Create(tr *ImageRequest) (*Image, *Response, error) {
	u := fmt.Sprintf("/images")
	req, err := s.client.NewRequestAccount("POST", u, tr)
	if err != nil {
		return nil, nil, err
	}

	image := new(imageResponse)
	resp, err := s.client.Do(req, image)
	if err != nil {
		return nil, nil, err
	}
	return image.Image, resp, nil
}
