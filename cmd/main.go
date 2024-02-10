package main

import (
	"github.com/Aakashraz/book_mgmt_GO/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	routes.BookStoreRoutes(r)

	//This line associates the router r with the root ("/") path. It means that any incoming HTTP requests will be handled by the routes configured in your r router.
	http.Handle("/", r)

	if err := http.ListenAndServe("localhost:9010", r); err != nil {
		log.Fatal(err)
	}
}
