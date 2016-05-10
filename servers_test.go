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

func TestServersService_List(t *testing.T) {
	setup()
	defer teardown()

	client.AuthToken = "654c95b0-2cf5-41a3-b3cc-733ffba4b4b7"

	data := testOpenFixture(t, filepath.Join(fixtureDir, "servers_list.json"))

	mux.HandleFunc("/servers", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Header().Add("Content-Type", contentType)
		fmt.Fprint(w, string(data))
	})

	servers, _, err := client.Servers.List()
	if err != nil {
		t.Errorf("Servers.List returned error: %v", err)
	}

	want := []*Server{
		{
			BootScript:      "",
			DynamicPublicIP: false,
			ID:              "741db378-6b87-46d4-a8c5-4e46a09ab1f8",
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
					ID:           "c1eb8f3a-4f0b-4b95-a71c-93223e457f5a",
					Name:         "vol simple snapshot",
					Organization: "000a115d-2852-4b0a-9ce8-47f1134ba95a",
					Size:         10000000000,
					Type:         "l_ssd",
				},
			},
		},
	}

	if !reflect.DeepEqual(servers, want) {
		t.Errorf("Servers.List returned %+v\n, want %+v", servers, want)
	}
}

func TestServersService_Get(t *testing.T) {
	setup()
	defer teardown()

	client.AuthToken = "654c95b0-2cf5-41a3-b3cc-733ffba4b4b7"

	data := testOpenFixture(t, filepath.Join(fixtureDir, "servers_get.json"))

	want := &Server{
		BootScript:      "",
		DynamicPublicIP: false,
		ID:              "741db378-6b87-46d4-a8c5-4e46a09ab1f8",
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
				ID:           "c1eb8f3a-4f0b-4b95-a71c-93223e457f5a",
				Name:         "vol simple snapshot",
				Organization: "000a115d-2852-4b0a-9ce8-47f1134ba95a",
				Size:         10000000000,
				Type:         "l_ssd",
			},
		},
	}

	mux.HandleFunc(fmt.Sprintf("/servers/%s", want.ID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Header().Add("Content-Type", contentType)
		fmt.Fprint(w, string(data))
	})

	server, _, err := client.Servers.Get(want.ID)
	if err != nil {
		t.Errorf("Servers.Get returned error: %v", err)
	}
	if !reflect.DeepEqual(server, want) {
		t.Errorf("Servers.Get returned %+v\n, want %+v", server, want)
	}
}

func TestServersService_Delete(t *testing.T) {
	setup()
	defer teardown()

	client.AuthToken = "654c95b0-2cf5-41a3-b3cc-733ffba4b4b7"

	serverID := "654c95b0-2cf5-41a3-b3cc-733ffba4b4b7"

	mux.HandleFunc(fmt.Sprintf("/servers/%s", serverID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.Header().Add("Content-Type", contentType)
	})

	_, err := client.Servers.Delete(serverID)
	if err != nil {
		t.Errorf("Servers.Delete returned error: %v", err)
	}
}
