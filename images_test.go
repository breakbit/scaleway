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

func TestImagesService_Create(t *testing.T) {

	setup()
	defer teardown()

	inBody := &ImageRequest{
		Organization: "000a115d-2852-4b0a-9ce8-47f1134ba95a",
		Arch:         "arm",
		Name:         "my_image",
		RootVolume:   "f0361e7b-cbe4-4882-a999-945192b7171b",
	}

	data := testOpenFixture(t, filepath.Join(fixtureDir, "images_create.json"))

	mux.HandleFunc("/images", func(w http.ResponseWriter, r *http.Request) {
		v := new(ImageRequest)
		json.NewDecoder(r.Body).Decode(v)

		testMethod(t, r, "POST")
		if !reflect.DeepEqual(v, inBody) {
			t.Errorf("Request body = %+v, want %+v", v, inBody)
		}
		w.Header().Add("Content-Type", contentType)

		fmt.Fprint(w, string(data))
	})

	image, _, err := client.Images.Create(inBody)
	if err != nil {
		t.Errorf("Images.Create returned error: %v", err)
	}
	creationDate, _ := time.Parse("2006-01-02T15:04:05.999999Z07:00", "2014-05-22T12:56:56.984011+00:00")
	modificationDate, _ := time.Parse("2006-01-02T15:04:05.999999Z07:00", "2014-05-22T12:56:56.984011+00:00")

	want := &Image{
		CreationDate:     Ntime(creationDate),
		ModificationDate: Ntime(modificationDate),
		Arch:             "arm",
		ID:               "98bf3ac2-a1f5-471d-8c8f-1b706ab57ef0",
		//ExtraVolumes:     []string{},
		FromImage:      "",
		FromServer:     "",
		MarketplaceKey: "",
		Name:           "my_image",
		Organization:   "000a115d-2852-4b0a-9ce8-47f1134ba95a",
		Public:         false,
		RootVolume: &Volume{
			Name: "vol-0-1",
			ID:   "f0361e7b-cbe4-4882-a999-945192b7171b",
		},
	}
	if !reflect.DeepEqual(image, want) {
		t.Errorf("Images.Create returned %+v\n, want %+v", image, want)
	}
}

func TestImagesService_List(t *testing.T) {
	setup()
	defer teardown()

	client.AuthToken = "654c95b0-2cf5-41a3-b3cc-733ffba4b4b7"

	data := testOpenFixture(t, filepath.Join(fixtureDir, "images_list.json"))

	mux.HandleFunc("/images", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Header().Add("Content-Type", contentType)
		fmt.Fprint(w, string(data))
	})

	images, _, err := client.Images.List()
	if err != nil {
		t.Errorf("Images.List returned error: %v", err)
	}

	creationDate, _ := time.Parse("2006-01-02T15:04:05.999999Z07:00", "2014-05-22T12:56:56.984011+00:00")
	modificationDate, _ := time.Parse("2006-01-02T15:04:05.999999Z07:00", "2014-05-22T12:56:56.984011+00:00")

	want := []*Image{
		{
			CreationDate:     Ntime(creationDate),
			ModificationDate: Ntime(modificationDate),
			Arch:             "arm",
			ID:               "98bf3ac2-a1f5-471d-8c8f-1b706ab57ef0",
			//ExtraVolumes:     []string{},
			FromImage:      "",
			FromServer:     "",
			MarketplaceKey: "",
			Name:           "my_image",
			Organization:   "000a115d-2852-4b0a-9ce8-47f1134ba95a",
			Public:         false,
			RootVolume: &Volume{
				Name: "vol-0-1",
				ID:   "f0361e7b-cbe4-4882-a999-945192b7171b",
			},
		},
	}
	if !reflect.DeepEqual(images, want) {
		t.Errorf("Images.List returned %+v\n, want %+v", images, want)
	}
}

func TestImagesService_Get(t *testing.T) {
	setup()
	defer teardown()

	client.AuthToken = "654c95b0-2cf5-41a3-b3cc-733ffba4b4b7"

	data := testOpenFixture(t, filepath.Join(fixtureDir, "images_get.json"))

	creationDate, _ := time.Parse("2006-01-02T15:04:05.999999Z07:00", "2014-05-22T12:56:56.984011+00:00")
	modificationDate, _ := time.Parse("2006-01-02T15:04:05.999999Z07:00", "2014-05-22T12:56:56.984011+00:00")

	want := &Image{
		CreationDate:     Ntime(creationDate),
		ModificationDate: Ntime(modificationDate),
		Arch:             "arm",
		ID:               "98bf3ac2-a1f5-471d-8c8f-1b706ab57ef0",
		//ExtraVolumes:     []string{},
		FromImage:      "",
		FromServer:     "",
		MarketplaceKey: "",
		Name:           "my_image",
		Organization:   "000a115d-2852-4b0a-9ce8-47f1134ba95a",
		Public:         false,
		RootVolume: &Volume{
			Name: "vol-0-1",
			ID:   "f0361e7b-cbe4-4882-a999-945192b7171b",
		},
	}

	mux.HandleFunc(fmt.Sprintf("/images/%s", want.ID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Header().Add("Content-Type", contentType)
		fmt.Fprint(w, string(data))
	})

	image, _, err := client.Images.Get(want.ID)
	if err != nil {
		t.Errorf("Images.Get returned error: %v", err)
	}
	if !reflect.DeepEqual(image, want) {
		t.Errorf("Images.Get returned %+v\n, want %+v", image, want)
	}
}

func TestImagesService_Delete(t *testing.T) {
	setup()
	defer teardown()

	client.AuthToken = "654c95b0-2cf5-41a3-b3cc-733ffba4b4b7"

	imageID := "654c95b0-2cf5-41a3-b3cc-733ffba4b4b7"

	mux.HandleFunc(fmt.Sprintf("/images/%s", imageID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.Header().Add("Content-Type", contentType)
	})

	_, err := client.Images.Delete(imageID)
	if err != nil {
		t.Errorf("Images.Delete returned error: %v", err)
	}
}
