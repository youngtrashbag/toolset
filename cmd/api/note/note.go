package note

import (
	"time"

	_ "github.com/go-sql-driver/mysql" // this is needed for mysql
)

// Note : Struct used for writing note
type Note struct {
	ID           int64     `json:"id"`
	Title        string    `json:"title"`
	Content      string    `json:"content"`
	CreationDate time.Time `json:"creation_date"`
	AuthorID     int64     `json:"author_id"`
}

// NewNote : a constructor for the Note struct
func NewNote(t string, c string, a int64) Note {
	var n Note

	n.Title = t
	n.Content = c
	n.AuthorID = a
	n.CreationDate = time.Now()

	return n
}
