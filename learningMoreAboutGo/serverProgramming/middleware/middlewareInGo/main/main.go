package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/justinas/alice"
)

func loggingHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Println("loggingHandler")
		t1 := time.Now()
		next.ServeHTTP(w, r)
		t2 := time.Now()
		log.Printf("[%s] %q %v\n", r.Method, r.URL.String(), t2.Sub(t1))
	}

	return http.HandlerFunc(fn)
}

func recoverHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Println("recoverHandler")
		defer func() {
			if err := recover(); err != nil {
				log.Printf("panic: %+v", err)
				http.Error(w, http.StatusText(500), 500)
			}
		}()

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("aboutHandler")
	fmt.Fprintf(w, "You are on the about page.")
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("indexHandler")
	fmt.Fprintf(w, "Welcome!")
}

func panicHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("panicHandler")
	panic("panic occurred")
}

func main() {
	commonHandlers := alice.New(loggingHandler, recoverHandler)
	http.Handle("/about", commonHandlers.ThenFunc(aboutHandler))
	http.Handle("/", commonHandlers.ThenFunc(indexHandler))
	http.Handle("/panic", commonHandlers.ThenFunc(panicHandler))
	http.ListenAndServe(":8080", nil)
}
