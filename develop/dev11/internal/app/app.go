package app

import (
	"log"
	"net/http"
	"time"

	"github.com/wb-third-lvl/develop/dev11/internal/transport/router"
)

func Run() {
	mux := http.NewServeMux()
	router := router.NewRouter()
	router.Register(mux)

	handler := Logging(mux)

	log.Fatal(http.ListenAndServe("localhost:8080", handler))
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, req)
		log.Printf("%s %s %s", req.Method, req.RequestURI, time.Since(start))
	})
}
