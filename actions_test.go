package scaleway

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"
	"reflect"
	"testing"
)

func TestActionsService_Exec(t *testing.T) {

	setup()
	defer teardown()

	client.AuthToken = "654c95b0-2cf5-41a3-b3cc-733ffba4b4b7"

	serverID := "741db378-6b87-46d4-a8c5-4e46a09ab1f8"

	inBody := &ActionRequest{
		"poweroff",
	}

	data := testOpenFixture(t, filepath.Join(fixtureDir, "actions_exec.json"))

	mux.HandleFunc(fmt.Sprintf("/servers/%s/action", serverID), func(w http.ResponseWriter, r *http.Request) {
		v := new(ActionRequest)
		json.NewDecoder(r.Body).Decode(v)

		testMethod(t, r, "POST")
		if !reflect.DeepEqual(v, inBody) {
			t.Errorf("Request body = %+v, want %+v", v, inBody)
		}
		w.Header().Add("Content-Type", contentType)

		fmt.Fprint(w, string(data))
	})

	task, _, err := client.Actions.Exec(serverID, inBody)
	if err != nil {
		t.Errorf("Actions.Exec returned error: %v", err)
	}

	want := &Task{
		Description: "server_poweroff",
		HrefFrom:    "/servers/741db378-6b87-46d4-a8c5-4e46a09ab1f8/action",
		ID:          "a8a1775c-0dda-4f52-87b2-4e8101d68d6e",
		Progress:    "0",
		Status:      "pending",
	}

	if !reflect.DeepEqual(task, want) {
		t.Errorf("Actions.Exec returned %+v\n, want %+v", task, want)
	}
}

func TestActionsService_List(t *testing.T) {
	setup()
	defer teardown()

	client.AuthToken = "654c95b0-2cf5-41a3-b3cc-733ffba4b4b7"

	serverID := "741db378-6b87-46d4-a8c5-4e46a09ab1f8"

	data := testOpenFixture(t, filepath.Join(fixtureDir, "actions_list.json"))

	mux.HandleFunc(fmt.Sprintf("/servers/%s/action", serverID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Header().Add("Content-Type", contentType)
		fmt.Fprint(w, string(data))
	})

	actions, _, err := client.Actions.List(serverID)
	if err != nil {
		t.Errorf("Actions.List returned error: %v", err)
	}

	want := []string{"poweron", "poweroff", "reboot"}

	if !reflect.DeepEqual(actions, want) {
		t.Errorf("Actions.List returned %+v\n, want %+v", actions, want)
	}
}
