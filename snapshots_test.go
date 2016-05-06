package scaleway

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"
	"reflect"
	"testing"
	"time"
)

func TestSnapshotsService_Create(t *testing.T) {

	setup()
	defer teardown()

	client.AuthToken = "654c95b0-2cf5-41a3-b3cc-733ffba4b4b7"

	inBody := &SnapshotRequest{
		Name:         "snapshot-0-1",
		Organization: "000a115d-2852-4b0a-9ce8-47f1134ba95a",
		Volume:       "701a8946-ff9d-4579-95e3-1c2c2d0f892d",
	}

	data := testOpenFixture(t, filepath.Join(fixtureDir, "snapshots_create.json"))

	mux.HandleFunc("/snapshots", func(w http.ResponseWriter, r *http.Request) {
		v := new(SnapshotRequest)
		json.NewDecoder(r.Body).Decode(v)

		testMethod(t, r, "POST")
		if !reflect.DeepEqual(v, inBody) {
			t.Errorf("Request body = %+v, want %+v", v, inBody)
		}
		w.Header().Add("Content-Type", contentType)

		fmt.Fprint(w, string(data))
	})

	snapshot, _, err := client.Snapshots.Create(inBody)
	if err != nil {
		t.Errorf("Snapshots.Create returned error: %v", err)
	}
	creationDate, _ := time.Parse("2006-01-02T15:04:05.999999Z07:00", "2014-05-22T12:10:05.596769+00:00")
	want := &Snapshot{
		BaseVolume: &Volume{
			ID:   "701a8946-ff9d-4579-95e3-1c2c2d0f892d",
			Name: "vol simple snapshot",
		},
		CreationDate: Ntime(creationDate),
		ID:           "f0361e7b-cbe4-4882-a999-945192b7171b",
		Name:         "snapshot-0-1",
		Organization: "000a115d-2852-4b0a-9ce8-47f1134ba95a",
		Size:         10000000000,
		State:        "snapshotting",
		Type:         "l_ssd",
	}
	if !reflect.DeepEqual(snapshot, want) {
		t.Errorf("Snapshots.Create returned %+v, want %+v", snapshot, want)
	}
}

func TestSnapshotsService_List(t *testing.T) {

	setup()
	defer teardown()

	client.AuthToken = "654c95b0-2cf5-41a3-b3cc-733ffba4b4b7"

	data := testOpenFixture(t, filepath.Join(fixtureDir, "snapshots_list.json"))

	mux.HandleFunc("/snapshots", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Header().Add("Content-Type", contentType)
		fmt.Fprint(w, string(data))
	})

	snapshots, _, err := client.Snapshots.List()
	if err != nil {
		t.Errorf("Snaphosts.List returned error: %v", err)
	}
	creationDate, _ := time.Parse("2006-01-02T15:04:05.999999Z07:00", "2014-05-22T12:11:06.055998+00:00")
	want := []*Snapshot{
		{
			BaseVolume: &Volume{
				ID:   "09a4184c-733b-43c8-99c3-f1dde30536fe",
				Name: "vol simple snapshot",
			},
			CreationDate: Ntime(creationDate),
			ID:           "6f418e5f-b42d-4423-a0b5-349c74c454a4",
			Name:         "snapshot-0-1",
			Organization: "000a115d-2852-4b0a-9ce8-47f1134ba95a",
			Size:         10000000000,
			State:        "snapshotting",
			Type:         "l_ssd",
		},
	}
	if !reflect.DeepEqual(snapshots, want) {
		t.Errorf("Snapshots.List returned %+v\n, want %+v", snapshots, want)
	}

}

func TestSnapshotsService_Get(t *testing.T) {

	setup()
	defer teardown()

	client.AuthToken = "654c95b0-2cf5-41a3-b3cc-733ffba4b4b7"

	data := testOpenFixture(t, filepath.Join(fixtureDir, "snapshots_get.json"))

	creationDate, _ := time.Parse("2006-01-02T15:04:05.999999Z07:00", "2014-05-22T12:11:06.055998+00:00")
	want := &Snapshot{
		BaseVolume: &Volume{
			ID:   "09a4184c-733b-43c8-99c3-f1dde30536fe",
			Name: "vol simple snapshot",
		},
		CreationDate: Ntime(creationDate),
		ID:           "6f418e5f-b42d-4423-a0b5-349c74c454a4",
		Name:         "snapshot-0-1",
		Organization: "000a115d-2852-4b0a-9ce8-47f1134ba95a",
		Size:         10000000000,
		State:        "snapshotting",
		Type:         "l_ssd",
	}

	mux.HandleFunc(fmt.Sprintf("/snapshots/%s", want.ID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Header().Add("Content-Type", contentType)
		fmt.Fprint(w, string(data))
	})

	snapshot, _, err := client.Snapshots.Get(want.ID)
	if err != nil {
		t.Errorf("Snapshots.Get returned error: %v", err)
	}
	if !reflect.DeepEqual(snapshot, want) {
		t.Errorf("Snapshots.Get returned %+v\n, want %+v", snapshot, want)
	}

}

func TestSnapshotsService_Update(t *testing.T) {

	setup()
	defer teardown()

	client.AuthToken = "654c95b0-2cf5-41a3-b3cc-733ffba4b4b7"

	data := testOpenFixture(t, filepath.Join(fixtureDir, "snapshots_update.json"))

	creationDate, _ := time.Parse("2006-01-02T15:04:05.999999Z07:00", "2014-05-22T12:11:06.055998+00:00")
	want := &Snapshot{
		BaseVolume: &Volume{
			ID:   "09a4184c-733b-43c8-99c3-f1dde30536fe",
			Name: "vol simple snapshot",
		},
		CreationDate: Ntime(creationDate),
		ID:           "6f418e5f-b42d-4423-a0b5-349c74c454a4",
		Name:         "snapshot-0-1",
		Organization: "000a115d-2852-4b0a-9ce8-47f1134ba95a",
		Size:         10000000000,
		State:        "snapshotting",
		Type:         "l_ssd",
	}

	mux.HandleFunc(fmt.Sprintf("/snapshots/%s", want.ID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		w.Header().Add("Content-Type", contentType)
		fmt.Fprint(w, string(data))
	})

	inBody := &SnapshotRequest{
		Organization: "000a115d-2852-4b0a-9ce8-47f1134ba95a",
	}

	snapshot, _, err := client.Snapshots.Update(want.ID, inBody)
	if err != nil {
		t.Errorf("Snapshots.Update returned error: %v", err)
	}
	if !reflect.DeepEqual(snapshot, want) {
		t.Errorf("Snapshots.Update returned %+v\n, want %+v", snapshot, want)
	}

}

func TestSnapshotsService_Delete(t *testing.T) {

	setup()
	defer teardown()

	client.AuthToken = "654c95b0-2cf5-41a3-b3cc-733ffba4b4b7"

	snapshotID := "6f418e5f-b42d-4423-a0b5-349c74c454a4"

	mux.HandleFunc(fmt.Sprintf("/snapshots/%s", snapshotID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.Header().Add("Content-Type", contentType)
	})

	_, err := client.Snapshots.Delete(snapshotID)
	if err != nil {
		t.Errorf("Snapshots.Delete returned error: %v", err)
	}

}
