package service

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/repodevs/gomis/pkg/util"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type Server struct {
	Router *mux.Router
	DBCon  *sql.DB
}

// Init Service
func (service *Server) Init() {
	service.Router = mux.NewRouter()
	service.initResource()
}

// Connect to database
func (service *Server) ConnectDB(dbHost, dbPort, dbUser, dbPass, dbName string) {
	service.DBInit(dbHost, dbPort, dbUser, dbPass, dbName)
}

// Start the server in given port
func (service *Server) Start(port int) {
	log.Println(fmt.Sprintf(util.StartingServer, port))
	err := http.ListenAndServe(fmt.Sprintf(":%v", port), service.Router)

	if err != nil {
		log.Fatal(err)
	}

	// Init Route
	http.Handle("/", service.Router)
}

// Resource URLs
func (service *Server) initResource() {
	log.Println(util.DeployingResources)

	service.Router.HandleFunc("/ping", service.pingService).Methods(http.MethodGet)

	// get All Notes
	service.Router.HandleFunc("/notes", service.getAllNotes).Methods(http.MethodGet)
	// get note by id: GET /notes/{id}
	service.Router.HandleFunc("/notes/{id}", service.getNoteByID).Methods(http.MethodGet)
	// add note
	service.Router.HandleFunc("/notes", service.addNote).Methods(http.MethodPost)
	// delete note by id: DELETE /notes/{id}
	service.Router.HandleFunc("/notes/{id}", service.deleteNoteByID).Methods(http.MethodDelete)

	log.Println(util.DeployingResourcesSuccess)
}

// handle ping request
func (service *Server) pingService(w http.ResponseWriter, r *http.Request) {
	log.Println("Got Ping Request! ", r.Header, r.Body)

	// write response
	w.WriteHeader(200)
	pongMsg := []byte("PONG!")
	_, err := w.Write(pongMsg)
	if err != nil {
		log.Fatal(err)
	}
}

// Failure Response
func (service *Server) FailureResponse(w http.ResponseWriter, code int, message string) {
	service.WriteResponse(w, code, message, nil)
}

// Write Response
func (service *Server) WriteResponse(w http.ResponseWriter, code int, message string, data interface{}) {
	response := Response{Status: code, Message: message, Data: data}
	payload, _ := json.Marshal(response)

	w.Header().Set(util.ContentType, util.ApplicationJson)
	w.WriteHeader(code)

	_, err := w.Write(payload)
	if err != nil {
		switch err {
		case http.ErrBodyNotAllowed:
			fmt.Println(err)
			break
		default:
			log.Fatal(err)
			break
		}
	}
}
