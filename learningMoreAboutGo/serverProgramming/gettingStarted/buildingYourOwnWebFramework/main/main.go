package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/roujiamo-cold/go-wiki/learningMoreAboutGo/serverProgramming/gettingStarted/buildingYourOwnWebFramework/middleware"

	"github.com/roujiamo-cold/go-wiki/learningMoreAboutGo/serverProgramming/gettingStarted/buildingYourOwnWebFramework/context"

	"github.com/justinas/alice"
)

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

func init() {

}

func main() {
	db, _ := sql.Open("postgres", "...")
	appC := middleware.NewAppContext(db)
	commonHandlers := alice.New(context.ClearHandler, middleware.LoggingHandler, middleware.RecoverHandler)
	http.Handle("/admin", commonHandlers.Append(appC.AuthHandler).ThenFunc(appC.AdminHandler))
	http.Handle("/about", commonHandlers.ThenFunc(aboutHandler))
	http.Handle("/", commonHandlers.ThenFunc(indexHandler))
	http.Handle("/panic", commonHandlers.ThenFunc(panicHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
