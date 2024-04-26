package router

import (
	"github.com/amrendrayadav/mongoapi/controllers"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/v1/movies", controllers.GetMyAllMovies).Methods("GET")
	router.HandleFunc("/api/v1/movie", controllers.CreateMovie).Methods("POST")
	router.HandleFunc("/api/v1/movie/{ID}", controllers.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/v1/movie/{ID}", controllers.DeleteAMovie).Methods("DELETE")
	router.HandleFunc("/api/v1/deleteall", controllers.DeleteAllMovie).Methods("DELETE")

	return router
}
