package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users = []User{
	{ID: 1, Name: "Ong", Age: 20},
	{ID: 2, Name: "nong", Age: 30},
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		if b, err := json.Marshal(users); err == nil {
			w.Header().Add("Content-Type", "application/json")
			w.Write(b)
			return
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
	}

	if r.Method == "POST" {
		user := User{}
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		if err := json.Unmarshal(body, &user); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		users = append(users, user)
		fmt.Fprintf(w, "Hello %s", user.Name)
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func logMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("Server http middlware: %s %s %s %s", r.RemoteAddr, r.Method, r.URL, time.Since(start))
	}
}

type Logger struct {
	Handler http.Handler
}

func (l Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	l.Handler.ServeHTTP(w, r)
	log.Printf("Server http middlware: %s %s %s %s", r.RemoteAddr, r.Method, r.URL, time.Since(start))
}

type Auth struct {
	Handler http.Handler
}

func (a Auth) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	u, p, ok := r.BasicAuth()
	if !ok {
		w.WriteHeader(401)
		w.Write([]byte(`Can't parse the basic auth`))
		return
	}

	if u != "ong" || p != "pass" {
		w.WriteHeader(401)
		w.Write([]byte(`Username/Password incorrect.`))
		return
	}
	log.Println("Auth passed.")
	a.Handler.ServeHTTP(w, r)
}

func AUthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u, p, ok := r.BasicAuth()
		if !ok {
			w.WriteHeader(401)
			w.Write([]byte(`Can't parse the basic auth`))
			return
		}

		if u != "ong" || p != "pass" {
			w.WriteHeader(401)
			w.Write([]byte(`Username/Password incorrect.`))
			return
		}
		log.Println("Auth passed.")
		next.ServeHTTP(w, r)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/users", handler)
	mux.HandleFunc("/health", healthHandler)

	authMux := Auth{Handler: mux}
	logMux := Logger{Handler: authMux}

	srv := http.Server{
		Addr:    ":2565",
		Handler: logMux,
	}

	log.Println("Server up by port 2565")
	log.Fatal(srv.ListenAndServe())
	log.Println("Bye")
}
