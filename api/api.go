package api

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/Ashfmate/library_api/data"
)

var (
	books []data.Book
	mu    sync.Mutex
)
// PostData handles adding or updating a book.
func PostData(w http.ResponseWriter, r *http.Request) {
	var newBook data.Book
	err := json.NewDecoder(r.Body).Decode(&newBook)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	// Look for the same book
	for i, book := range books {
		if book.ID == newBook.ID {
			// Update existing book
			books[i] = newBook
			w.WriteHeader(http.StatusOK)
			return
		}
	}

	// Add new book
	books = append(books, newBook)
	w.WriteHeader(http.StatusCreated)
}

// GetData handles retrieving all books.
func GetData(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(books)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}