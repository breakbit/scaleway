package scaleway

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"reflect"
	"testing"
)

var (
	// mux is the HTTP request multiplexer used with the test server.
	mux *http.ServeMux
	// client is the GitHub client being tested.
	client *Client
	// server is a test HTTP server used to provide mock API responses.
	server *httptest.Server
)

// setup sets up a test HTTP server along with a scaleway.Client that is
// configured to talk to that test server.  Tests should register handlers on
// mux which provide mock responses for the API method being tested.
func setup() {
	// test server
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	// scaleway client configured to use test server
	client = NewClient(nil)
	url, _ := url.Parse(server.URL)
	client.AccountBaseURL = url
	client.ComputeBaseURL = url
}

// teardown closes the test HTTP server.
func teardown() {
	server.Close()
}

// testMethod check if the method request if what you want.
func testMethod(t *testing.T, r *http.Request, want string) {
	if got := r.Method; got != want {
		t.Errorf("Request method: %v, want %v", got, want)
	}
}

// This is the directory where our test fixtures are
const fixtureDir = "./test-fixtures"

// Test Helpers
func testOpenFixture(t *testing.T, name string) []byte {
	f, err := os.Open(name)
	defer f.Close()
	if err != nil {
		t.Fatal(err)
	}
	data, err := ioutil.ReadAll(f)
	if err != nil {
		t.Fatal(err)
	}
	return data
}

func TestNewClient(t *testing.T) {
	c := NewClient(nil)

	if got, want := c.AccountBaseURL.String(), defaultAccountBaseURL; got != want {
		t.Errorf("NewClient AccountBaseURL is %v, want %v", got, want)
	}
	if got, want := c.ComputeBaseURL.String(), defaultComputeBaseURL; got != want {
		t.Errorf("NewClient ComputeBaseURL is %v, want %v", got, want)
	}
	if got, want := c.UserAgent, userAgent; got != want {
		t.Errorf("NewClient AccountBaseURL is %v, want %v", got, want)
	}
}

type inBody struct {
	A interface{}
}

type requestAccountTest struct {
	inURL   string
	outURL  string
	inBody  inBody
	outBody string
	wantErr bool
}

var requestAccountTests = []requestAccountTest{
	{"/foo", defaultAccountBaseURL + "foo", inBody{A: "a"}, `{"A":"a"}` + "\n", false},
	{"%zzzzzz", defaultAccountBaseURL + "foo", inBody{A: "a"}, ``, true},
	{"/foo", defaultAccountBaseURL + "foo", inBody{A: func() {}}, ``, true},
}

func TestNewRequestAccount(t *testing.T) {
	for _, tt := range requestAccountTests {
		c := NewClient(nil)

		req, err := c.NewRequestAccount("GET", tt.inURL, tt.inBody)
		if (err != nil) != tt.wantErr {
			t.Errorf("NewRequest error: %v", err)
			continue
		}

		if req == nil {
			continue
		}

		if got, want := req.URL.String(), tt.outURL; got != want {
			t.Errorf("Newrequest (%q) URL is %v, want %v", tt.inURL, got, want)
		}

		body, _ := ioutil.ReadAll(req.Body)
		if got, want := string(body), tt.outBody; got != want {
			t.Errorf("Newrequest (%v) Body is %#v, want %#v", tt.inBody, got, want)
		}

		if got, want := req.Header.Get("User-Agent"), c.UserAgent; got != want {
			t.Errorf("Newrequest () User-Agent %v, want %v", got, want)
		}

		if got, want := req.Header.Get("Content-Type"), contentType; got != want {
			t.Errorf("Newrequest () Content-Type %v, want %v", got, want)
		}
	}
}

type requestComputeTest struct {
	inURL   string
	outURL  string
	inBody  inBody
	outBody string
	wantErr bool
}

var requestComputeTests = []requestComputeTest{
	{"/foo", defaultComputeBaseURL + "foo", inBody{A: "a"}, `{"A":"a"}` + "\n", false},
	{"%zzzzzz", defaultComputeBaseURL + "foo", inBody{A: "a"}, ``, true},
	{"/foo", defaultComputeBaseURL + "foo", inBody{A: func() {}}, ``, true},
}

var requestTests = []requestComputeTest{
	{"/foo", "/foo", inBody{A: "a"}, `{"A":"a"}` + "\n", false},
	{"%zzzzzz", "", inBody{A: "a"}, ``, true},
	{"/foo", "/foo", inBody{A: func() {}}, ``, true},
}

func TestNewRequest(t *testing.T) {
	for _, tt := range requestTests {
		c := NewClient(nil)

		req, err := c.newRequest("GET", tt.inURL, tt.inBody)
		if (err != nil) != tt.wantErr {
			t.Errorf("NewRequest error: %v", err)
			continue
		}

		if req == nil {
			continue
		}

		if got, want := req.URL.String(), tt.outURL; got != want {
			t.Errorf("Newrequest (%q) URL is %v, want %v", tt.inURL, got, want)
		}

		body, _ := ioutil.ReadAll(req.Body)
		if got, want := string(body), tt.outBody; got != want {
			t.Errorf("Newrequest (%v) Body is %#v, want %#v", tt.inBody, got, want)
		}

		if got, want := req.Header.Get("User-Agent"), c.UserAgent; got != want {
			t.Errorf("Newrequest () User-Agent %v, want %v", got, want)
		}

		if got, want := req.Header.Get("Content-Type"), contentType; got != want {
			t.Errorf("Newrequest () Content-Type %v, want %v", got, want)
		}
	}
}

func TestNewRequestCompute(t *testing.T) {
	for _, tt := range requestComputeTests {
		c := NewClient(nil)

		req, err := c.NewRequestCompute("GET", tt.inURL, tt.inBody)
		if (err != nil) != tt.wantErr {
			t.Errorf("NewRequest error: %v", err)
			continue
		}

		if req == nil {
			continue
		}

		if got, want := req.URL.String(), tt.outURL; got != want {
			t.Errorf("Newrequest (%q) URL is %v, want %v", tt.inURL, got, want)
		}

		body, _ := ioutil.ReadAll(req.Body)
		if got, want := string(body), tt.outBody; got != want {
			t.Errorf("Newrequest (%v) Body is %#v, want %#v", tt.inBody, got, want)
		}

		if got, want := req.Header.Get("User-Agent"), c.UserAgent; got != want {
			t.Errorf("Newrequest () User-Agent %v, want %v", got, want)
		}

		if got, want := req.Header.Get("Content-Type"), contentType; got != want {
			t.Errorf("Newrequest () Content-Type %v, want %v", got, want)
		}
	}
}

func TestDo(t *testing.T) {
	setup()
	defer teardown()

	type foo struct {
		A string
	}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"A":"a"}`)
	})

	req, _ := client.NewRequestAccount("GET", "/", nil)
	body := new(foo)
	client.Do(req, body)

	want := &foo{"a"}
	if !reflect.DeepEqual(body, want) {
		t.Errorf("Response body = %v, want %v", body, want)
	}
}
