package routers

import (
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"

	"github.com/Latias94/taskmanager/common"
	"github.com/Latias94/taskmanager/controllers"
)

func SetNoteRoutes(router *mux.Router) *mux.Router {
	noteRoter := mux.NewRouter()
	noteRoter.HandleFunc("/notes", controllers.CreateNote).Methods("POST")
	noteRoter.HandleFunc("/notes/{id}", controllers.UpdateNote).Methods("PUT")
	noteRoter.HandleFunc("/notes/{id}", controllers.GetNoteById).Methods("GET")
	noteRoter.HandleFunc("/notes", controllers.GetNotes).Methods("GET")
	noteRoter.HandleFunc("/notes/tasks/{id}", controllers.GetNotesByTask).Methods("GET")
	noteRoter.HandleFunc("/notes/{id}", controllers.DeleteNote).Methods("DELETE")
	router.PathPrefix("/notes").Handler(negroni.New(
		negroni.HandlerFunc(common.Authorize),
		negroni.Wrap(noteRoter),
	))
	return router

}
