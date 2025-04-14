package router

import (
	"github.com/gorilla/mux"
	"github.com/vamshireddy02/go-postgres/controllers"
)



func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("api/stock/{id}", controllers.GetStock).Methods("GET", "OPTIONS")
	router.HandleFunc("api/stock", controllers.GetAllStock).Methods("GET", "OPTION")
	router.HandleFunc("api/newstock", controllers.CreateStock).Methods("POST", "OPTION")
	router.HandleFunc("api/stock/{id}", controllers.UpdateStock).Methods("PUT", "OPTION")
	router.HandleFunc("api/deletestock/{id}", controllers.DeleteStock).Methods("DELETE", "OPTION")

	return router
}
