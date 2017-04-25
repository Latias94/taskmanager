package routers

import "github.com/gorilla/mux"

func InitRoutes() *mux.Router {
	router:=mux.NewRouter().StrictSlash(false)
	// Routes for the User entity
	router=SetUserRoutes(router)
	router=SetTaskRoutes(router)
	router=SetNoteRoutes(router)
	return router
}
