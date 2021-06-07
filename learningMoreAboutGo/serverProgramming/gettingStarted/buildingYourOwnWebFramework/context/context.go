package context

import (
	"log"
	"net/http"
	"sync"
)

var (
	mutex sync.RWMutex
	data  = make(map[*http.Request]map[interface{}]interface{})
)

func Set(r *http.Request, k, v interface{}) {
	mutex.Lock()
	defer mutex.Unlock()

	if _, ok := data[r]; !ok {
		data[r] = make(map[interface{}]interface{})
	}

	data[r][k] = v
}

func Get(r *http.Request, k interface{}) interface{} {
	mutex.RLock()
	defer mutex.RUnlock()

	v, ok := data[r]
	if !ok {
		return nil
	}
	return v[k]
}

func Clear(r *http.Request) {
	log.Println("Clear")
	mutex.Lock()
	defer mutex.Unlock()

	delete(data, r)
}

func ClearHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer Clear(r)

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
