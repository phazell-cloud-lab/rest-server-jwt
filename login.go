package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type User struct {
	Name  string
	Pword Password
}

type Password struct {
	Value string `json:"password"`
}

func login(w http.ResponseWriter, r *http.Request, allowed []User) {
	login_with := getCredentials(r)
	authorised := checkAuthenticatedUsers(login_with, allowed)

	if authorised {
		// create JWT
		// dummy good response
		dummy_response := []byte("Your in Sunshine...\n")
		w.Write(dummy_response)

		// return JWT
	} else {
		// return unauthorised error message
		w.WriteHeader(401)
	}
}

func getCredentials(r *http.Request) User {
	// read username from parameter
	username := r.URL.Query().Get("uname")

	// read password from body
	var pword Password
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	body := buf.String()
	json.Unmarshal([]byte(body), &pword)

	// if (err != nil) || (pword.Value == "") {
	// 	// return error
	// }

	return User{username, pword}
}

func checkAuthenticatedUsers(u User, authorised []User) bool {
	for _, x := range authorised {
		if x == u {
			return true
		}
	}
	return false
}
