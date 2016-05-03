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
}

func TestVolumesService_List(t *testing.T) {
	setup()
	defer teardown()

	client.AuthToken = "654c95b0-2cf5-41a3-b3cc-733ffba4b4b7"
}

func TestVolumesService_Delete(t *testing.T) {
	setup()
	defer teardown()

	client.AuthToken = "654c95b0-2cf5-41a3-b3cc-733ffba4b4b7"
}
