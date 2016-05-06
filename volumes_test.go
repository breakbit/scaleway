package scaleway

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"
	"reflect"
	"testing"
)

func TestVolumesService_Create(t *testing.T) {
	setup()
	defer teardown()

	inBody := &VolumeRequest{
		"volume-0-3",
		"000a115d-2852-4b0a-9ce8-47f1134ba95a",
		"l_ssd",
		10000000000,
	}

	client.AuthToken = "654c95b0-2cf5-41a3-b3cc-733ffba4b4b7"

	data := testOpenFixture(t, filepath.Join(fixtureDir, "volumes_create.json"))

	mux.HandleFunc("/volumes", func(w http.ResponseWriter, r *http.Request) {
		v := new(VolumeRequest)
		json.NewDecoder(r.Body).Decode(v)

		testMethod(t, r, "POST")
		if !reflect.DeepEqual(v, inBody) {
			t.Errorf("Request body = %+v, want %+v", v, inBody)
		}
		w.Header().Add("Content-Type", contentType)

		fmt.Fprint(w, string(data))
	})

	volume, _, err := client.Volumes.Create(inBody)
	if err != nil {
		t.Errorf("Tokens.Create returned error: %v", err)
	}
	want := &Volume{
		Name:         "volume-0-3",
		ID:           "c675f420-cfeb-48ff-ba2a-9d2a4dbe3fcd",
		ExportURI:    "",
		Organization: "000a115d-2852-4b0a-9ce8-47f1134ba95a",
		Size:         10000000000,
		Type:         "l_ssd",
	}
	if !reflect.DeepEqual(volume, want) {
		t.Errorf("Volumes.Create returned %+v\n, want %+v", volume, want)
	}
}

func TestVolumesService_Get(t *testing.T) {
	setup()
	defer teardown()

	client.AuthToken = "654c95b0-2cf5-41a3-b3cc-733ffba4b4b7"

	data := testOpenFixture(t, filepath.Join(fixtureDir, "volumes_get.json"))

	want := &Volume{
		Name:         "volume-0-1",
		ID:           "f929fe39-63f8-4be8-a80e-1e9c8ae22a76",
		ExportURI:    "",
		Organization: "000a115d-2852-4b0a-9ce8-47f1134ba95a",
		Size:         10000000000,
		Type:         "l_ssd",
	}

	mux.HandleFunc(fmt.Sprintf("/volumes/%s", want.ID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Header().Add("Content-Type", contentType)
		fmt.Fprint(w, string(data))
	})

	volume, _, err := client.Volumes.Get(string(want.ID))
	if err != nil {
		t.Errorf("Volumes.Get returned error: %v", err)
	}
	if !reflect.DeepEqual(volume, want) {
		t.Errorf("Volumes.Get returned %+v\n, want %+v", volume, want)
	}
}

func TestVolumesService_List(t *testing.T) {
	setup()
	defer teardown()

	client.AuthToken = "654c95b0-2cf5-41a3-b3cc-733ffba4b4b7"

	data := testOpenFixture(t, filepath.Join(fixtureDir, "volumes_list.json"))

	mux.HandleFunc("/volumes", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Header().Add("Content-Type", contentType)
		fmt.Fprint(w, string(data))
	})

	volumes, _, err := client.Volumes.List()
	if err != nil {
		t.Errorf("Volumes.List returned error: %v", err)
	}
	want := []*Volume{
		{
			ExportURI:    "",
			ID:           "f929fe39-63f8-4be8-a80e-1e9c8ae22a76",
			Name:         "volume-0-1",
			Organization: "000a115d-2852-4b0a-9ce8-47f1134ba95a",
			Size:         10000000000,
			Type:         "l_ssd",
		},
		{
			ExportURI:    "",
			ID:           "0facb6b5-b117-441a-81c1-f28b1d723779",
			Name:         "volume-0-2",
			Organization: "000a115d-2852-4b0a-9ce8-47f1134ba95a",
			Size:         20000000000,
			Type:         "l_ssd",
		},
	}
	if !reflect.DeepEqual(volumes, want) {
		t.Errorf("Volumes.List returned %+v\n, want %+v", volumes, want)
	}
}

func TestVolumesService_Delete(t *testing.T) {
	setup()
	defer teardown()

	client.AuthToken = "654c95b0-2cf5-41a3-b3cc-733ffba4b4b7"

	volumeID := "0facb6b5-b117-441a-81c1-f28b1d723779"

	mux.HandleFunc(fmt.Sprintf("/volumes/%s", volumeID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.Header().Add("Content-Type", contentType)
	})

	_, err := client.Volumes.Delete(volumeID)
	if err != nil {
		t.Errorf("Volumes.Delete returned error: %v", err)
	}
}
