package scaleway

import (
	"fmt"
	"net/http"
	"path/filepath"
	"reflect"
	"testing"
)

func TestOrganizationService_List(t *testing.T) {

	setup()
	defer teardown()

	client.AuthToken = "654c95b0-2cf5-41a3-b3cc-733ffba4b4b7"

	data := testOpenFixture(t, filepath.Join(fixtureDir, "organizations_list.json"))

	mux.HandleFunc("/organizations", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Header().Add("Content-Type", contentType)
		fmt.Fprint(w, string(data))
	})

	org, _, err := client.Organizations.List()
	if err != nil {
		t.Errorf("Organizations.List returned error: %v", err)
	}

	want := []*Organization{
		&Organization{
			ID:   "000a115d-2852-4b0a-9ce8-47f1134ba95a",
			Name: "jsnow@got.com",
			Users: []*User{
				&User{
					Email:      "jsnow@got.com",
					Firstname:  "John",
					Fullname:   "John Snow",
					ID:         "59a98700-8622-4495-a11a-e1efbfac5972",
					Lastname:   "Snow",
					SSHPubKeys: []string{},
				},
			},
		},
	}
	if !reflect.DeepEqual(org, want) {
		t.Errorf("Organization.List returned %+v\n, want %+v", org, want)
	}
}
