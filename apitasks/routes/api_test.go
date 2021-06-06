package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/vsergio011/apitasks/models"
)

func TestUsersCRUD(t *testing.T) {
	idAux := ""
	t.Run("TEST GET USER", func(t *testing.T) {
		req, err := http.NewRequest("GET", "api/user", nil)
		req = mux.SetURLVars(req, map[string]string{"ID": "M1VXKsoMHqO9vMFZaREVSVKJxQX2"})

		if err != nil {
			t.Fatalf("could not created request: %v", err)
		}
		a := new(api)
		rec := httptest.NewRecorder()
		a.fetchUser(rec, req)

		res := rec.Result()
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			t.Errorf("expected %d, got: %d", http.StatusOK, res.StatusCode)
		}

		b, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatalf("could not read response: %v", err)
		}
		var got *models.User
		err = json.Unmarshal(b, &got)
		fmt.Println()
		if err != nil {
			t.Fatalf("could not unmarshall response %v", err)
		}

		if got.Email == "" {
			t.Errorf("expected 5 gophers, got: %v gopher", got.Email)
		}
	})
	t.Run("TEST LIST USERS", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/api/users", nil)
		if err != nil {
			t.Fatalf("could not created request: %v", err)
		}
		a := new(api)
		rec := httptest.NewRecorder()
		a.fetchUsers(rec, req)

		res := rec.Result()
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			t.Errorf("expected %d, got: %d", http.StatusOK, res.StatusCode)
		}

		b, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatalf("could not read response: %v", err)
		}
		var got []*models.User
		err = json.Unmarshal(b, &got)
		fmt.Println()
		if err != nil {
			t.Fatalf("could not unmarshall response %v", err)
		}

		if len(got) == 0 {
			t.Errorf("expected 5 gophers, got: %v gopher", len(got))
		}
	})
	t.Run("TEST CREATE USER", func(t *testing.T) {
		bodyReader := strings.NewReader(`{
			"Id": "prueba",
			"Name": "usuarioDePrueba",
			"Surname": "surnameDeprueba",
			"Email": "prueba@gmal.com",
			"PhoneNumber": "+34658456365",
			"Password": "passwordprueba",
			"DisplayName": "namePruba",
			"Token": ""
		}`)
		req, err := http.NewRequest("POST", "api/user", bodyReader)

		if err != nil {
			t.Fatalf("could not created request: %v", err)
		}
		a := new(api)
		rec := httptest.NewRecorder()
		a.addUser(rec, req)

		res := rec.Result()
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			t.Errorf("expected %d, got: %d", http.StatusOK, res.StatusCode)
		}

		b, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatalf("could not read response: %v", err)
		}
		var got *models.User
		err = json.Unmarshal(b, &got)
		fmt.Println()
		if err != nil {
			t.Fatalf("could not unmarshall response %v", err)
		}
		idAux = got.RawID
		if got.Email != "prueba@gmal.com" {
			t.Errorf("expected 5 gophers, got: %v gopher", got.RawID)
		}
	})
	t.Run("TEST UPDATE USER", func(t *testing.T) {
		bodyReader := strings.NewReader(`{
			"Id": "prueba",
			"Name": "updateusuarioDePrueba",
			"Surname": "surnameDeprueba",
			"Email": "pruebaupd@gmal.com",
			"PhoneNumber": "+34658456365",
			"Password": "passwordprueba",
			"DisplayName": "namePruba",
			"Token": ""
		}`)

		req, err := http.NewRequest("POST", "api/upduser", bodyReader)
		req = mux.SetURLVars(req, map[string]string{"ID": idAux})

		if err != nil {
			t.Fatalf("could not created request: %v", err)
		}
		a := new(api)
		rec := httptest.NewRecorder()
		a.updateUser(rec, req)

		res := rec.Result()
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			t.Errorf("expected %d, got: %d", http.StatusOK, res.StatusCode)
		}

		b, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatalf("could not read response: %v", err)
		}
		var got *models.User
		err = json.Unmarshal(b, &got)
		fmt.Println()
		if err != nil {
			t.Fatalf("could not unmarshall response %v", err)
		}
		idAux = got.RawID
		if got.Email != "pruebaupd@gmal.com" {
			t.Errorf("expected email pruebaupd@gmal.com, got: %v email", got.Email)
		}
	})
	t.Run("TEST DELETE USER", func(t *testing.T) {
		req, err := http.NewRequest("GET", "api/rmuser", nil)
		req = mux.SetURLVars(req, map[string]string{"ID": idAux})

		if err != nil {
			t.Fatalf("could not created request: %v", err)
		}
		a := new(api)
		rec := httptest.NewRecorder()
		a.removeUser(rec, req)

		res := rec.Result()
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			t.Errorf("expected %d, got: %d", http.StatusOK, res.StatusCode)
		}

		b, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatalf("could not read response: %v", err)
		}

		if b == nil {
			t.Errorf("expected 5 gophers, got:  gopher")
		}
	})

}
