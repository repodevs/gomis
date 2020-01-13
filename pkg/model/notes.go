package model

import (
	"database/sql"
)

type Note struct {
	Id      int64  `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

type NoteId struct {
	Id int64 `json:"id"`
}

// define query
const insertNoteQuery = "INSERT INTO notes(title, content, author) VALUES ($1, $2, $3) RETURNING id;"
const getNoteByIdQuery = "SELECT title, content, author FROM notes WHERE id=$1;"
const getAllNotes = "SELECT * FROM notes;"
const deleteNoteByIdQuery = "DELETE FROM notes WHERE id=$1 RETURNING id;"

// Add Note and return the id
func (note *Note) AddNote(db *sql.DB) error {
	var noteId int64
	if err := db.QueryRow(insertNoteQuery, note.Title, note.Content, note.Author).Scan(&noteId); err != nil {
		return err
	} else {
		note.Id = noteId
	}
	return nil
}

// Get Note by ID
func (note *Note) GetNoteById(db *sql.DB) error {
	return db.QueryRow(getNoteByIdQuery, note.Id).Scan(&note.Title, &note.Content, &note.Author)
}

// Get All Notes
func GetAllNotes(db *sql.DB) ([]Note, error) {
	result, err := db.Query(getAllNotes)
	if err != nil {
		return nil, err
	}
	defer result.Close()
	notes := []Note{}

	for result.Next() {
		var note Note

		if err = result.Scan(&note.Id, &note.Title, &note.Content, &note.Author); err != nil {
			return nil, err
		}
		notes = append(notes, note)
	}

	return notes, nil
}

// Delete Note by ID
func (note *Note) DeleteNoteById(db *sql.DB) error {
	return db.QueryRow(deleteNoteByIdQuery, note.Id).Scan(&note.Id)
}
