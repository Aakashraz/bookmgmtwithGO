package controllers

import (
	"encoding/json"
	"github.com/Aakashraz/book_mgmt_GO/pkg/models"
	"github.com/Aakashraz/book_mgmt_GO/pkg/utils"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

var newBook models.Book

func ListBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write(res)
	if err != nil {
		log.Printf("error while writing into response: %s", err)
	}
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"] //this "bookId" is corresponding to the id mapped in the bookstore-router
	Id, err := strconv.ParseInt(bookId, 0, 0)
	//0, 0: The second and third arguments are the base and a bit size for the conversion.
	//Setting both to 0 means that ParseInt will automatically detect the base (e.g., hexadecimal, octal, or decimal)
	//and use the bit size of the platform (32 or 64 bits).
	if err != nil {
		log.Println("Error while parsing, str conversion to int")
	}

	bookDetail, _ := models.GetBookById(Id)
	res, err := json.Marshal(bookDetail)
	if err != nil {
		log.Printf("error while marshalling %s:", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(res)
	if err != nil {
		log.Printf("error while writing into response: %s", err)
	}
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	// Create an instance of the models.Book struct
	createBook := &models.Book{}
	// Parse the request body and populate the createBook struct
	utils.ParseBody(r, createBook)
	// Call the CreateBook method on the createBook instance
	b := createBook.CreateBook()

	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Write the JSON response to the response writer
	_, err := w.Write(res)
	if err != nil {
		log.Printf("error while writing into response: %s", err)
	}
}

func DeleteBook() {

}
