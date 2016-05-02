package scaleway

import (
	"fmt"
	"net/http"
	"path/filepath"
	"reflect"
	"testing"
)

func TestUserService_Get(t *testing.T) {
	setup()
	defer teardown()

	client.AuthToken = "654c95b0-2cf5-41a3-b3cc-733ffba4b4b7"

	data := testOpenFixture(t, filepath.Join(fixtureDir, "users_get.json"))

	want := &User{
		ID:        "5bea0358-db40-429e-bd82-953016a7e2s7",
		Email:     "jsnow@got.com",
		Firstname: "John",
		Fullname:  "John Snow",
		Lastname:  "Snow",
		//SSHPubKeys: []string{},
	}

	mux.HandleFunc(fmt.Sprintf("/users/%s", want.ID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Header().Add("Content-Type", contentType)
		fmt.Fprint(w, string(data))
	})

	user, _, err := client.Users.Get(want.ID)
	if err != nil {
		t.Errorf("Users.Get returned error: %v", err)
	}

	if !reflect.DeepEqual(user, want) {
		t.Errorf("Users.Get returned %#v\n, want %#v", user, want)
	}
}
