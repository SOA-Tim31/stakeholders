package routing

import (
	"net/http"
	"stakeholders/handler"

	"github.com/gorilla/mux"
)

func SetupRoutes(handler *handler.UserHandler, handler2 *handler.PersonHandler, handler3 *handler.AppRatingHandler) http.Handler {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/users/getAll", handler.FindAllUsers).Methods("GET", "OPTIONS")
	router.HandleFunc("/users/block", handler.BlockOrUblock).Methods("PUT", "OPTIONS")
	router.HandleFunc("/users/register", handler.Register).Methods("POST", "OPTIONS")
	router.HandleFunc("/people/get/{id}", handler2.Get).Methods("GET", "OPTIONS")
	router.HandleFunc("/people/update", handler2.Update).Methods("PUT", "OPTIONS")
	router.HandleFunc("/ratings/create", handler3.Create).Methods("POST", "OPTIONS")

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))

	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization")
			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}
			next.ServeHTTP(w, r)
		})
	})

	return router
}
