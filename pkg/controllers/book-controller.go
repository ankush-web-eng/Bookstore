package controllers

import (
	"net/http"
	"strconv"

	"github.com/ankush-web-eng/Bookstore/pkg/config"
	"github.com/ankush-web-eng/Bookstore/pkg/models"
	"github.com/ankush-web-eng/Bookstore/pkg/utils"
	"github.com/gorilla/mux"
)

// var newBook models.Book

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	books := models.GetAllBooks()
	utils.JsonResponse(w, http.StatusOK, books)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookId := params["bookId"]
	ID, err := strconv.ParseInt(bookId, 10, 64)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	bookDetails, err := models.GetBookById(ID)
	if err != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	utils.JsonResponse(w, http.StatusOK, bookDetails)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	newBook := &models.Book{}
	utils.ParseBody(r, newBook)
	createdBook := newBook.CreateBook()
	utils.JsonResponse(w, http.StatusCreated, createdBook)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookId := params["bookId"]
	ID, err := strconv.ParseInt(bookId, 10, 64)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	deletedBook := models.DeleteBook(ID)
	utils.JsonResponse(w, http.StatusOK, deletedBook)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	updatedBook := &models.Book{}
	utils.ParseBody(r, updatedBook)
	params := mux.Vars(r)
	bookId := params["bookId"]
	ID, err := strconv.ParseInt(bookId, 10, 64)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	bookDetails, err := models.GetBookById(ID)
	if err != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	// Update bookDetails with updatedBook fields
	if updatedBook.Name != "" {
		bookDetails.Name = updatedBook.Name
	}
	if updatedBook.Author != "" {
		bookDetails.Author = updatedBook.Author
	}
	if updatedBook.Publication != "" {
		bookDetails.Publication = updatedBook.Publication
	}

	db := config.GetDB()
	db.Save(&bookDetails)

	// Return updated bookDetails as JSON response
	utils.JsonResponse(w, http.StatusOK, bookDetails)
}
