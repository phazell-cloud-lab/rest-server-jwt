package main

import (
	"log"
	"net/http"
)

func serve() {
	err := http.ListenAndServe(":8474", nil)
	if err != nil {
		log.Fatal("ListenAndServe failed ", err)
	}
}

func main() {

	Barry := User{"Barry", Password{"Chicken"}}
	Sheila := User{"Sheila", Password{"Dippers"}}
	authorised_users := []User{Barry, Sheila}

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			login(w, r, authorised_users)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	serve()
}
