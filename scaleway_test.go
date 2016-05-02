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

func TestNewRequestAccount(t *testing.T) {
	c := NewClient(nil)

	inURL, outURL := "/foo", defaultAccountBaseURL+"foo"
	inBody := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Expires  bool   `json:"expires"`
	}{
		"jsnow@got.com",
		"winteriscoming",
		true,
	}
	outBody := `{"email":"jsnow@got.com","password":"winteriscoming","expires":true}` + "\n"
	req, _ := c.NewRequestAccount("GET", inURL, inBody)

	if got, want := req.URL.String(), outURL; got != want {
		t.Errorf("Newrequest (%q) URL is %v, want %v", inURL, got, want)
	}

	body, _ := ioutil.ReadAll(req.Body)
	if got, want := string(body), outBody; got != want {
		t.Errorf("Newrequest (%v) Body is %v, want %v", inBody, got, want)
	}

	if got, want := req.Header.Get("User-Agent"), c.UserAgent; got != want {
		t.Errorf("Newrequest () User-Agent %v, want %v", got, want)
	}

	if got, want := req.Header.Get("Content-Type"), contentType; got != want {
		t.Errorf("Newrequest () Content-Type %v, want %v", got, want)
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