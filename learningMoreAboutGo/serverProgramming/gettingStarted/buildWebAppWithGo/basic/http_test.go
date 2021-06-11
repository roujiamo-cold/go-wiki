package main

import (
	"fmt"
	"log"
	"net/http"
	"testing"
)

func TestNilMap(t *testing.T) {
	type user struct {
		m map[string]string
	}

	u := &user{}

	if d, exist := u.m["name"]; !exist {
		fmt.Println("not exist")
	} else {
		fmt.Println(d)
	}
}

func TestHttpServer(t *testing.T) {

	host := "prozac.com"
	testHost := "test.prozac.com"
	docsHost := "docs.prozac.com"

	userHandle := func(w http.ResponseWriter, r *http.Request) {
		name := r.FormValue("name")
		if name == "" {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		fmt.Fprintf(w, "hello %s", name)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello world")
	})
	http.Handle("/user", http.HandlerFunc(userHandle))

	http.HandleFunc("/path/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, r.URL.Path)
	})

	http.HandleFunc("/path", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "this is /path")
	})

	http.HandleFunc(host+"/some", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(host)
		fmt.Fprintf(w, "%s%s", r.Host, r.URL.Path)
	})
	http.HandleFunc(testHost+"/some", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(testHost)
		fmt.Fprintf(w, "%s%s", r.Host, r.URL.Path)
	})
	http.HandleFunc(docsHost+"/some", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(docsHost)
		fmt.Fprintf(w, "%s%s", r.Host, r.URL.Path)
	})

	log.Fatal(http.ListenAndServe(":8888", nil))
}
