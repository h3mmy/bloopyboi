package models

// import "github.com/h3mmy/bloopyboi/ent"

type BookSearchRequest struct {
	// Isbn is a 10 or 13 digit ISBN
	ISBN string `json:"isbn,omitempty"`
	// Title is the title of the book
	Title string `json:"title,omitempty"`
	// Author is the author of the book
	Author string `json:"author,omitempty"`
	// Publisher is the publisher of the book
	Publisher string `json:"publisher,omitempty"`
	// Year is the year the book was published
	Year int `json:"year,omitempty"`
	// TextSnippet is a free-form text search
	TextSnippet string `json:"text,omitempty"`
	// DiscordUser is the Discord user who made the request
	// DiscordUser *ent.DiscordUser `json:"discorduser,omitempty"`
}
