package scaleway

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"
	"reflect"
	"testing"
)

func TestServersService_Create(t *testing.T) {

	setup()
	defer teardown()

	inBody := &ServerRequest{
		Organization: "000a115d-2852-4b0a-9ce8-47f1134ba95a",
		Name:         "my_image",
		Image:        "85917034-46b0-4cc5-8b48-f0a2245e357e",
		Tags:         []string{"test", "www"},
	}

	data := testOpenFixture(t, filepath.Join(fixtureDir, "servers_create.json"))

	mux.HandleFunc("/servers", func(w http.ResponseWriter, r *http.Request) {
		v := new(ServerRequest)
		json.NewDecoder(r.Body).Decode(v)

		testMethod(t, r, "POST")
		if !reflect.DeepEqual(v, inBody) {
			t.Errorf("Request body = %+v, want %+v", v, inBody)
		}
		w.Header().Add("Content-Type", contentType)

		fmt.Fprint(w, string(data))
	})

	server, _, err := client.Servers.Create(inBody)
	if err != nil {
		t.Errorf("Servers.Create returned error: %v", err)
	}

	want := &Server{
		BootScript:      "",
		DynamicPublicIP: false,
		ID:              "3cb18e2d-f4f7-48f7-b452-59b88ae8fc8c",
		Image: &Image{
			ID:   "85917034-46b0-4cc5-8b48-f0a2245e357e",
			Name: "archlinux working",
		},
		Name:         "my_server",
		Organization: "000a115d-2852-4b0a-9ce8-47f1134ba95a",
		PrivateIP:    "",
		PublicIP:     "",
		State:        "running",
		Tags:         []string{"test", "www"},
		Volumes: map[string]*Volume{
			"0": &Volume{
				ExportURI:    "",
				ID:           "d9257116-6919-49b4-a420-dcfdff51fcb1",
				Name:         "vol simple snapshot",
				Organization: "000a115d-2852-4b0a-9ce8-47f1134ba95a",
				Size:         10000000000,
				Type:         "l_ssd",
			},
		},
	}

	if !reflect.DeepEqual(server, want) {
		t.Errorf("Servers.Create returned %+v\n, want %+v", server, want)
	}
}
