package routes

import (
	"github.com/Aakashraz/book_mgmt_GO/pkg/controllers"
	"github.com/gorilla/mux"
)

var BookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.ListBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
