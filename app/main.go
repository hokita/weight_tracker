package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	os.Exit(0)
}

func run() error {
	fmt.Println("start server")

	r := mux.NewRouter()
	r.HandleFunc("/", handle).Methods("GET")
	r.HandleFunc("/weights/", postHandle).Methods("POST")

	if err := http.ListenAndServe(":8080", r); err != nil {
		return err
	}

	return nil
}

func handle(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("templates/index.html"))

	b, err := ioutil.ReadFile("data.txt")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	m := map[string]string{
		"Result": string(b),
	}
	tpl.Execute(w, m)
}

func postHandle(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("templates/index.html"))

	f, err := os.OpenFile("data.txt", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer func() {
		if err := f.Close(); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}()

	fmt.Fprintln(f, r.FormValue("weight"))

	b, err := ioutil.ReadFile("data.txt")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	m := map[string]string{
		"Result": string(b),
	}
	tpl.Execute(w, m)
}
