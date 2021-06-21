package main

import (
	"html/template"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

func init() {
	http.Handle("/view", appHandler(viewRecord))
}

type appError struct {
	Error   error
	Message string
	Code    int
}

type appHandler func(http.ResponseWriter, *http.Request) *appError

func (fn appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if e := fn(w, r); e != nil { // e is *appError, not os.Error.
		http.Error(w, e.Message, e.Code)
	}
}

type Record struct {
	ID      string
	Content string
}

func viewRecord(w http.ResponseWriter, r *http.Request) *appError {
	c := appengine.NewContext(r)
	key := datastore.NewKey(c, "Record", r.FormValue("id"), 0, nil)
	record := new(Record)
	if err := datastore.Get(c, key, record); err != nil {
		return &appError{err, "Record not found", 404}
	}
	if err := viewTemplate.Execute(w, record); err != nil {
		return &appError{err, "Can't display record", 500}
	}
	return nil
}

var viewTemplate = template.Must(template.New("results").Parse(`
<html>
<head/>
<body>
  <p>ID: {{.ID}} Content: {{.Content}}</p>
</body>
</html>
`))
