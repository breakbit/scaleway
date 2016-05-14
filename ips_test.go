package scaleway

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"
	"reflect"
	"testing"
)

func TestIPsService_Create(t *testing.T) {

	setup()
	defer teardown()

	inBody := &IPRequest{
		Organization: "000a115d-2852-4b0a-9ce8-47f1134ba95a",
	}

	data := testOpenFixture(t, filepath.Join(fixtureDir, "ips_create.json"))

	mux.HandleFunc("/ips", func(w http.ResponseWriter, r *http.Request) {
		v := new(IPRequest)
		json.NewDecoder(r.Body).Decode(v)

		testMethod(t, r, "POST")
		if !reflect.DeepEqual(v, inBody) {
			t.Errorf("Request body = %+v, want %+v", v, inBody)
		}
		w.Header().Add("Content-Type", contentType)

		fmt.Fprint(w, string(data))
	})

	ip, _, err := client.IPs.Create(inBody)
	if err != nil {
		t.Errorf("IPs.Create returned error: %v", err)
	}

	want := &IP{
		Address:      "212.47.226.88",
		ID:           "b50cd740-892d-47d3-8cbf-88510ef626e7",
		Organization: "000a115d-2852-4b0a-9ce8-47f1134ba95a",
	}
	if !reflect.DeepEqual(ip, want) {
		t.Errorf("IPs.Create returned %+v\n, want %+v", ip, want)
	}

}

func TestIPsService_List(t *testing.T) {

	setup()
	defer teardown()

}

func TestIPsService_Get(t *testing.T) {

	setup()
	defer teardown()

}

func TestIPsService_Attach(t *testing.T) {

	setup()
	defer teardown()

}

func TestIPsService_Delete(t *testing.T) {

	setup()
	defer teardown()

}
