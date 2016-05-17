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

	client.AuthToken = "654c95b0-2cf5-41a3-b3cc-733ffba4b4b7"

	data := testOpenFixture(t, filepath.Join(fixtureDir, "ips_list.json"))

	mux.HandleFunc("/ips", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Header().Add("Content-Type", contentType)
		fmt.Fprint(w, string(data))
	})

	ips, _, err := client.IPs.List()
	if err != nil {
		t.Errorf("IPs.List returned error: %v", err)
	}
	want := []*IP{
		{
			Address:      "212.47.226.88",
			ID:           "b50cd740-892d-47d3-8cbf-88510ef626e7",
			Organization: "000a115d-2852-4b0a-9ce8-47f1134ba95a",
		},
	}
	if !reflect.DeepEqual(ips, want) {
		t.Errorf("IPs.List returned %+v\n, want %+v", ips, want)
	}

}

func TestIPsService_Get(t *testing.T) {

	setup()
	defer teardown()

	client.AuthToken = "654c95b0-2cf5-41a3-b3cc-733ffba4b4b7"

	data := testOpenFixture(t, filepath.Join(fixtureDir, "ips_get.json"))

	want := &IP{
		Address:      "212.47.226.88",
		ID:           "b50cd740-892d-47d3-8cbf-88510ef626e7",
		Organization: "000a115d-2852-4b0a-9ce8-47f1134ba95a",
	}

	mux.HandleFunc(fmt.Sprintf("/ips/%s", want.ID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Header().Add("Content-Type", contentType)
		fmt.Fprint(w, string(data))
	})

	ip, _, err := client.IPs.Get(want.ID)
	if err != nil {
		t.Errorf("IPs.Get returned error: %v", err)
	}
	if !reflect.DeepEqual(ip, want) {
		t.Errorf("IPS.Get returned %+v\n, want %+v", ip, want)
	}

}

func TestIPsService_Attach(t *testing.T) {

	setup()
	defer teardown()

	inBody := &IPRequest{
		Address:      "212.47.226.88",
		ID:           "b50cd740-892d-47d3-8cbf-88510ef626e7",
		Organization: "000a115d-2852-4b0a-9ce8-47f1134ba95a",
		Server:       "c2d8994f-1582-413e-8d48-c53076db06cc",
	}

	data := testOpenFixture(t, filepath.Join(fixtureDir, "ips_attach.json"))

	mux.HandleFunc(fmt.Sprintf("/ips/%s", inBody.ID), func(w http.ResponseWriter, r *http.Request) {
		v := new(IPRequest)
		json.NewDecoder(r.Body).Decode(v)

		testMethod(t, r, "PUT")
		if !reflect.DeepEqual(v, inBody) {
			t.Errorf("Request body = %+v, want %+v", v, inBody)
		}
		w.Header().Add("Content-Type", contentType)

		fmt.Fprint(w, string(data))
	})

	ip, _, err := client.IPs.Attach(inBody, inBody.ID)
	if err != nil {
		t.Errorf("IPs.Attach returned error: %v", err)
	}

	want := &IP{
		Address:      "212.47.226.88",
		ID:           "b50cd740-892d-47d3-8cbf-88510ef626e7",
		Organization: "000a115d-2852-4b0a-9ce8-47f1134ba95a",
		Server: &Server{
			ID:   "c2d8994f-1582-413e-8d48-c53076db06cc",
			Name: "default_server_name - acfb51",
		},
	}

	if !reflect.DeepEqual(ip, want) {
		t.Errorf("IPs.Attach returned %+v, want %+v", ip, want)
	}

}

func TestIPsService_Delete(t *testing.T) {

	setup()
	defer teardown()

	client.AuthToken = "654c95b0-2cf5-41a3-b3cc-733ffba4b4b7"

	ipID := "654c95b0-2cf5-41a3-b3cc-733ffba4b4b7"

	mux.HandleFunc(fmt.Sprintf("/ips/%s", ipID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.Header().Add("Content-Type", contentType)
	})

	_, err := client.IPs.Delete(ipID)
	if err != nil {
		t.Errorf("IPs.Delete returned error: %v", err)
	}

}
