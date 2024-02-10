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
	newBook := &models.Book{}
	// Parse the request body and populate the newBook struct
	utils.ParseBody(r, newBook)
	log.Printf("marshal value: %v", r)
	// Call the CreateBook method on the newBook instance
	b := newBook.CreateBook()

	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Write the JSON response to the response writer
	_, err := w.Write(res)
	if err != nil {
		log.Printf("error while writing into response: %s", err)
	}
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		log.Println("error while parsing into integer")
	}

	book := models.DeleteBook(ID)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(res)
	if err != nil {
		log.Printf("error while writing into response: %s", err)
	}
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	// Create an instance of the models.Book struct
	updateBook := &models.Book{}
	utils.ParseBody(r, updateBook)
	// to test the marshalled value
	log.Printf("marshal value for update: %v", r)

	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		log.Println("error while parsing into integer id")
	}

	bookDetails, db := models.GetBookById(ID)
	// print to check the value in updateBook and Bookdetails for debugging
	log.Printf("Bookdetails: %+v", bookDetails)
	log.Printf("updateBook: %+v", updateBook)

	//	to fetch the details of book for updating
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}

	db.Save(&bookDetails)

	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(res)
	if err != nil {
		log.Printf("error while writing into response: %s", err)
	}
}
