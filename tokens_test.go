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

func TestTokensService_Create(t *testing.T) {

	setup()
	defer teardown()

	credentials := NewCredentials("foo@bar.com", "foobar")
	tokenExpires := true

	inBody := &tokenRequest{
		Credentials: credentials,
		Expires:     tokenExpires,
	}

	data := testOpenFixture(t, filepath.Join(fixtureDir, "tokens_create.json"))

	mux.HandleFunc("/tokens", func(w http.ResponseWriter, r *http.Request) {
		v := new(tokenRequest)
		json.NewDecoder(r.Body).Decode(v)

		testMethod(t, r, "POST")
		if !reflect.DeepEqual(v, inBody) {
			t.Errorf("Request body = %+v, want %+v", v, inBody)
		}
		w.Header().Add("Content-Type", contentType)

		fmt.Fprint(w, string(data))
	})

	token, _, err := client.Tokens.Create(credentials, tokenExpires)
	if err != nil {
		t.Errorf("Tokens.Create returned error: %v", err)
	}
	creationDate, _ := time.Parse("2006-01-02T15:04:05.999999Z07:00", "2014-05-22T08:06:51.742826+00:00")
	expires, _ := time.Parse("2006-01-02T15:04:05.999999Z07:00", "2014-05-20T14:05:06.393875+00:00")
	want := &Token{
		CreationDate:      Ntime(creationDate),
		Expires:           Ntime(expires),
		ID:                "654c95b0-2cf5-41a3-b3cc-733ffba4b4b7",
		InheritsUserPerms: true,
		UserID:            "5bea0358-db40-429e-bd82-953016a7e2s7",
		Permission:        []string{},
	}
	if !reflect.DeepEqual(token, want) {
		t.Errorf("Tokens.Create returned %+v\n, want %+v", token, want)
	}
}

func TestTokensService_List(t *testing.T) {
	setup()
	defer teardown()

	client.AuthToken = "654c95b0-2cf5-41a3-b3cc-733ffba4b4b7"

	data := testOpenFixture(t, filepath.Join(fixtureDir, "tokens_list.json"))

	mux.HandleFunc("/tokens", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Header().Add("Content-Type", contentType)
		fmt.Fprint(w, string(data))
	})

	tokens, _, err := client.Tokens.List()
	if err != nil {
		t.Errorf("Tokens.List returned error: %v", err)
	}
	creationDate, _ := time.Parse("2006-01-02T15:04:05.999999Z07:00", "2014-05-22T08:06:51.742826+00:00")
	want := []*Token{
		{
			CreationDate:      Ntime(creationDate),
			Expires:           Ntime{},
			ID:                "654c95b0-2cf5-41a3-b3cc-733ffba4b4b7",
			InheritsUserPerms: true,
			UserID:            "5bea0358-db40-429e-bd82-953016a7e2s7",
			Permission:        []string{},
		},
	}
	if !reflect.DeepEqual(tokens, want) {
		t.Errorf("Tokens.List returned %+v\n, want %+v", tokens, want)
	}
}

func TestTokensService_Get(t *testing.T) {
	setup()
	defer teardown()

	client.AuthToken = "654c95b0-2cf5-41a3-b3cc-733ffba4b4b7"

	data := testOpenFixture(t, filepath.Join(fixtureDir, "tokens_get.json"))

	creationDate, _ := time.Parse("2006-01-02T15:04:05.999999Z07:00", "2014-05-22T08:06:51.742826+00:00")
	expires, _ := time.Parse("2006-01-02T15:04:05.999999Z07:00", "2014-05-20T14:05:06.393875+00:00")
	want := &Token{
		CreationDate:      Ntime(creationDate),
		Expires:           Ntime(expires),
		ID:                "654c95b0-2cf5-41a3-b3cc-733ffba4b4b7",
		InheritsUserPerms: true,
		UserID:            "5bea0358-db40-429e-bd82-953016a7e2s7",
		Permission:        []string{},
	}

	mux.HandleFunc(fmt.Sprintf("/tokens/%s", want.ID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Header().Add("Content-Type", contentType)
		fmt.Fprint(w, string(data))
	})

	token, _, err := client.Tokens.Get(want.ID)
	if err != nil {
		t.Errorf("Tokens.Get returned error: %v", err)
	}
	if !reflect.DeepEqual(token, want) {
		t.Errorf("Tokens.Get returned %+v\n, want %+v", token, want)
	}
}

func TestTokensService_Update(t *testing.T) {
	setup()
	defer teardown()

	client.AuthToken = "654c95b0-2cf5-41a3-b3cc-733ffba4b4b7"

	data := testOpenFixture(t, filepath.Join(fixtureDir, "tokens_update.json"))

	creationDate, _ := time.Parse("2006-01-02T15:04:05.999999Z07:00", "2014-05-22T08:06:51.742826+00:00")
	expires, _ := time.Parse("2006-01-02T15:04:05.999999Z07:00", "2014-05-22T11:18:07.786841+00:00")
	want := &Token{
		CreationDate:      Ntime(creationDate),
		Expires:           Ntime(expires),
		ID:                "654c95b0-2cf5-41a3-b3cc-733ffba4b4b7",
		InheritsUserPerms: true,
		UserID:            "5bea0358-db40-429e-bd82-953016a7e2s7",
		Permission:        []string{},
	}

	mux.HandleFunc(fmt.Sprintf("/tokens/%s", want.ID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PATCH")
		w.Header().Add("Content-Type", contentType)
		fmt.Fprint(w, string(data))
	})

	token, _, err := client.Tokens.Update(want.ID)
	if err != nil {
		t.Errorf("Tokens.Update returned error: %v", err)
	}
	if !reflect.DeepEqual(token, want) {
		t.Errorf("Tokens.Update returned %+v\n, want %+v", token, want)
	}
}

func TestTokensService_Delete(t *testing.T) {
	setup()
	defer teardown()

	client.AuthToken = "654c95b0-2cf5-41a3-b3cc-733ffba4b4b7"

	tokenID := "654c95b0-2cf5-41a3-b3cc-733ffba4b4b7"

	mux.HandleFunc(fmt.Sprintf("/tokens/%s", tokenID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.Header().Add("Content-Type", contentType)
	})

	_, err := client.Tokens.Delete(tokenID)
	if err != nil {
		t.Errorf("Tokens.Delete returned error: %v", err)
	}
}
