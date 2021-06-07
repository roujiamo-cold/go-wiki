package middleware

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"

	"github.com/roujiamo-cold/go-wiki/learningMoreAboutGo/serverProgramming/gettingStarted/buildingYourOwnWebFramework/context"
)

type AppContext struct {
	db *sql.DB
}

func NewAppContext(db *sql.DB) *AppContext {
	return &AppContext{db: db}
}

func (c *AppContext) AuthHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		authToken := r.Header.Get("Authorization")
		user, err := getUser(c.db, authToken)

		if err != nil {
			http.Error(w, http.StatusText(401), 401)
			return
		}

		context.Set(r, "user", user)
		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func (c *AppContext) AdminHandler(w http.ResponseWriter, r *http.Request) {
	user := context.Get(r, "user")
	// Maybe other operations on the database
	_ = json.NewEncoder(w).Encode(user)
}

func getUser(db *sql.DB, name string) (map[string]interface{}, error) {
	u := make(map[string]interface{})
	u["name"] = "xiaoming"
	u["age"] = 13

	if name == "xiaoming" {
		return u, nil
	}
	return nil, errors.New("user not found")
}
