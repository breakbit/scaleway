package scaleway

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"time"
)

const (
	libraryVersion        = "0.1"
	defaultAccountBaseURL = "https://account.scaleway.com/"
	defaultComputeBaseURL = "https://api.scaleway.com/"
	userAgent             = "go-scaleway/" + libraryVersion
	contentType           = "application/json"
)

// A Client manages communcation with the Scaleway API.
type Client struct {
	// HTTP client used to communicate with the API.
	client *http.Client
	// Base URL related to account actions.
	AccountBaseURL *url.URL
	// Base URL related to compute actions.
	ComputeBaseURL *url.URL
	// UserAgent used when communicating with the Scaleway API.
	UserAgent string
	// AuthToken used when communication with Scaleway API.
	AuthToken string
	// Services used for talking to Scaleway API.
	Tokens        *TokensService
	Organizations *OrganizationsService
	Users         *UsersService
	Volumes       *VolumesService
}

// timeLayout represents the time layout needed for parsing.
const timeLayout = "2006-01-02T15:04:05.999999Z07:00"

// Ntime represents custom time.
type Ntime time.Time

// UnmarshalJSON decodes JSON custom time.
// See: https://github.com/golang/go/issues/9037
func (t *Ntime) UnmarshalJSON(b []byte) error {
	if string(b) == "null" {
		return nil
	}
	tm, err := time.Parse(timeLayout, string(b[1:len(b)-1]))
	if err != nil {
		return err
	}
	*t = Ntime(tm)
	return nil
}

// Response is a Scaleway API Response. This wrap the standard http.Response.
type Response struct {
	*http.Response
}

// NewClient returns a new Scaleway API Client. If a nil httpClient
// provided, http.DefaultClient will be used.
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	accountBaseURL, _ := url.Parse(defaultAccountBaseURL)
	computeBaseURL, _ := url.Parse(defaultComputeBaseURL)

	c := &Client{client: httpClient,
		AccountBaseURL: accountBaseURL,
		ComputeBaseURL: computeBaseURL,
		UserAgent:      userAgent}
	c.Tokens = &TokensService{client: c}
	c.Organizations = &OrganizationsService{client: c}
	c.Users = &UsersService{client: c}
	c.Volumes = &VolumesService{client: c}
	return c
}

// NewRequestAccount creates an API request. A relative URL can be provided in urlStr,
// in which case it is resolved relative to the BaseURL of the Client.
// Relative URLs should always be specified without a preceding slash.  If
// specified, the value pointed to by body is JSON encoded and included as the
// request body.
func (c *Client) NewRequestAccount(method, urlStr string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := c.AccountBaseURL.ResolveReference(rel)
	return c.newRequest(method, u.String(), body)
}

// NewRequestCompute creates an API request. A relative URL can be provided in urlStr,
// in which case it is resolved relative to the ComputeBaseURL of the Client.
// Relative URLs should always be specified without a preceding slash.  If
// specified, the value pointed to by body is JSON encoded and included as the
// request body.
func (c *Client) NewRequestCompute(method, urlStr string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := c.ComputeBaseURL.ResolveReference(rel)
	return c.newRequest(method, u.String(), body)
}

// newRequest creates http.Request.
func (c *Client) newRequest(method, u string, body interface{}) (*http.Request, error) {
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u, buf)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)
	if c.UserAgent != "" {
		req.Header.Add("User-Agent", c.UserAgent)
	}
	if c.AuthToken != "" {
		req.Header.Add("X-Auth-Token", c.AuthToken)
	}
	return req, nil
}

// Do sends an API request and returns the API response.
// TODO @dmelina Need more documentation
func (c *Client) Do(req *http.Request, v interface{}) (*Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		resp.Body.Close()
	}()

	response := newResponse(resp)

	if v != nil {
		err = json.NewDecoder(resp.Body).Decode(v)
		if err == io.EOF {
			err = nil // ignore EOF errors caused by empty response body
		}
	}
	return response, err
}

// newResponse creates a new Response for the provided http.Response.
func newResponse(r *http.Response) *Response {
	return &Response{Response: r}
}
