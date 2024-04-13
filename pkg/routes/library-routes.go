package routes

//go lang has absolute paths i.e paths start from root folder
import (
	"github.com/Karthik-S-L/library-management-golan/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes= func (router *mux.Router)  {
	router.HandleFunc("/book",controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book",controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}",controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}",controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}",controllers.DeleteBook).Methods("DELETE")
}
