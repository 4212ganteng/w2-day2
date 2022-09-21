package main

import (
	"fmt"
	"net/http"
	blogcontroller "personal-web/blogController"
	"personal-web/contactController"

	"github.com/gorilla/mux"
)

func main() {

	route := mux.NewRouter()

	// route path folder public
	route.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

	// routing blog
	route.HandleFunc("/", blogcontroller.Index).Methods("GET")
	route.HandleFunc("/add-project", blogcontroller.Add).Methods("GET")
	route.HandleFunc("/store-project", blogcontroller.Store).Methods("POST")


	// route contact
	route.HandleFunc("/contact", contactController.Index).Methods("GET")
	route.HandleFunc("/store-contact", contactController.Store).Methods("POST")


	fmt.Println("server running on port 5000")
	http.ListenAndServe("localhost:5000", route)

}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
