package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"os"
)

var (
	Host   = reqEnv("HOST")
	Port   = reqEnv("PORT")
	ApiKey = reqEnv("API_KEY")
)

type ResponseDTO struct {
	Message string `json:"message"`
}

func main() {
	r := chi.NewRouter()

	// use default chi logger & set default content type to application/json
	r.Use(middleware.Logger)
	r.Use(contentTypeJsonMw)

	// public routes
	r.Group(func(r chi.Router) {
		r.Get("/alive", handleGetAlive())
	})

	// require api key to access private routes
	r.Group(func(r chi.Router) {
		r.Use(apiKeyMw)
		r.Get("/private", handleGetPrivate())
	})

	err := http.ListenAndServe(fmt.Sprintf("%s:%s", Host, Port), r)
	if err != nil {
		panic(fmt.Sprintf("failed to start server: %s", err))
		return
	}
}

// handleGetAlive returns a http.HandlerFunc that writes a json response or an error if marshalling or writing fails.
//
// see [ResponseDTO], [json.Marshal], [http.ResponseWriter.Write]
func handleGetAlive() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		writeJSONResponse(w, ResponseDTO{Message: "I'm alive!"})
	}
}

func handleGetPrivate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		writeJSONResponse(w, ResponseDTO{Message: "I'm private!"})
	}
}

func writeJSONResponse(w http.ResponseWriter, object any) {
	b, err := json.Marshal(object)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to marshal response: %s", err), http.StatusInternalServerError)
		return
	}
	_, _ = w.Write(b)
}

func contentTypeJsonMw(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		handler.ServeHTTP(w, r)
	})
}

func apiKeyMw(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-API-KEY") != ApiKey {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		handler.ServeHTTP(w, r)
	})
}

// reqEnv tries to look up an env value using os.LookupEnv, returns a value if it exists and panics if not
func reqEnv(key string) string {
	if env, ok := os.LookupEnv(key); ok {
		return env
	}
	panic(fmt.Sprintf("%s not set", key))
}
