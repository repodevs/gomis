package service

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/repodevs/gomis/pkg/model"
	"github.com/repodevs/gomis/pkg/util"
)

// Retrieve all notes
func (service *Server) getAllNotes(w http.ResponseWriter, r *http.Request) {
	notes, err := model.GetAllNotes(service.DBCon)

	if err != nil {
		service.FailureResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	service.WriteResponse(w, http.StatusOK, util.SuccessMsg, notes)
}

// Retrieve Note by ID
func (service *Server) getNoteByID(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 0, 0)

	if err != nil {
		service.FailureResponse(w, http.StatusBadRequest, util.InvalidNoteIdMsg)
		return
	}

	note := model.Note{Id: id}
	err = note.GetNoteById(service.DBCon)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			service.FailureResponse(w, http.StatusNotFound, util.NoteNotFoundMsg)
			break
		default:
			service.FailureResponse(w, http.StatusInternalServerError, err.Error())
			break
		}
		return
	}
	service.WriteResponse(w, http.StatusOK, util.SuccessMsg, note)
}

// Add Notes
func (service *Server) addNote(w http.ResponseWriter, r *http.Request) {
	var note model.Note

	body := json.NewDecoder(r.Body)
	if err := body.Decode(&note); err != nil {
		fmt.Println(err)
		service.FailureResponse(w, http.StatusBadRequest, util.BadRequestMsg)
		return
	}

	defer r.Body.Close()

	if note.Title == "" || note.Content == "" || note.Author == "" {
		service.FailureResponse(w, http.StatusBadRequest, util.BadRequestMsg)
		return
	}

	if err := note.AddNote(service.DBCon); err != nil {
		fmt.Println(err)
		service.FailureResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	service.WriteResponse(w, http.StatusCreated, util.SuccessMsg, model.NoteId{Id: note.Id})

}

// Delete Note by ID
func (service *Server) deleteNoteByID(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 0, 0)

	if err != nil {
		service.FailureResponse(w, http.StatusBadRequest, util.InvalidNoteIdMsg)
		return
	}

	note := model.Note{Id: id}
	// TODO: Should Check if Note exists?
	err = note.DeleteNoteById(service.DBCon)
	if err != nil {
		service.FailureResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	service.WriteResponse(w, http.StatusNoContent, util.SuccessMsg, model.Note{})
}
