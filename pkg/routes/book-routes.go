package routes

import (
	"github.com/gorilla/mux"
	"github.com/sifatistiak/book-management-system/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook)
}
