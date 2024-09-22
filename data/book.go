package data

import (
    "time"
)

// Book represents a book with relevant details.
type Book struct {
    ID          int       `json:"id"`
    Title       string    `json:"title"`
    Author      string    `json:"author"`
    PublishedAt time.Time `json:"published_at"`
    Pages       int       `json:"pages"`
}

// NewBook creates a new Book instance.
func NewBook(id int, title, author string, publishedAt time.Time, pages int) *Book {
    return &Book{
        ID:          id,
        Title:       title,
        Author:      author,
        PublishedAt: publishedAt,
        Pages:       pages,
    }
}

// UpdateTitle updates the title of the book.
func (b *Book) UpdateTitle(newTitle string) {
    b.Title = newTitle
}

// UpdateAuthor updates the author of the book.
func (b *Book) UpdateAuthor(newAuthor string) {
    b.Author = newAuthor
}

// UpdatePublishedAt updates the published date of the book.
func (b *Book) UpdatePublishedAt(newDate time.Time) {
    b.PublishedAt = newDate
}

// UpdatePages updates the number of pages of the book.
func (b *Book) UpdatePages(newPages int) {
    b.Pages = newPages
}