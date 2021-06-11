package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"

	"github.com/roujiamo-cold/go-wiki/learningMoreAboutGo/serverProgramming/gettingStarted/buildingYourOwnWebFramework/context"
	"github.com/roujiamo-cold/go-wiki/learningMoreAboutGo/serverProgramming/gettingStarted/buildingYourOwnWebFramework/middleware"
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

func teaHandler(w http.ResponseWriter, r *http.Request) {
	params := context.Get(r, "params").(httprouter.Params)
	tea := getTea(params.ByName("id"))
	json.NewEncoder(w).Encode(tea)
}

func getTea(id string) string {
	return fmt.Sprintf("this is tea[%s]", id)
}

func init() {

}

func wrapHandler(h http.Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		context.Set(r, "params", ps)
		h.ServeHTTP(w, r)
	}
}

func main() {
	db, _ := sql.Open("postgres", "...")
	appC := middleware.NewAppContext(db)
	commonHandlers := alice.New(context.ClearHandler, middleware.LoggingHandler, middleware.RecoverHandler)

	router := httprouter.New()
	router.GET("/admin", wrapHandler(commonHandlers.Append(appC.AuthHandler).ThenFunc(appC.AdminHandler)))
	router.GET("/about", wrapHandler(commonHandlers.ThenFunc(aboutHandler)))
	router.GET("/", wrapHandler(commonHandlers.ThenFunc(indexHandler)))
	router.GET("/panic", wrapHandler(commonHandlers.ThenFunc(panicHandler)))
	router.GET("/teas/:id", wrapHandler(commonHandlers.ThenFunc(teaHandler)))
	log.Fatal(http.ListenAndServe(":8080", router))
}
