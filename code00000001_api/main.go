package main

import (
	"encoding/json"
	"net/http"
)

var users = map[string]*User{}

type User struct {
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
}

func jsonContentTypeMiddleware(next http.Handler) http.Handler {

	// Response Header add Content-Type
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", "application/json")

		// 전달 받은 http.Handler를 호출한다.
		next.ServeHTTP(rw, r)
	})
}

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("hello"))
	})
	mux := http.NewServeMux()

	userHandler := http.HandlerFunc(func(wr http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet: // select
			json.NewEncoder(wr).Encode(users)
		case http.MethodPost: // insert
			var user User
			user.Nickname = "js"
			user.Email = "js@cleancode"
			json.NewDecoder(r.Body).Decode(&user)

			users[user.Email] = &user

			json.NewEncoder(wr).Encode(user)
		}
	})

	// 3
	mux.Handle("/users", jsonContentTypeMiddleware(userHandler))
	http.ListenAndServe(":8080", mux)
}
