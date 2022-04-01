package model

// import "time"

type Content struct {
	Title string
	Body  string
}
type Person struct {
	Firstname string
	Lastname  string
}

type Article struct {
	ID      string
	Content        // Promoted fields
	Author  Person // Nested struct
	// CreatedAt *time.Time
}
