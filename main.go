package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {

	route := mux.NewRouter()

	// route path folder public
	route.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

	// routing blog
	route.HandleFunc("/", index).Methods("GET")
	route.HandleFunc("/add-project", add).Methods("GET")
	route.HandleFunc("/store-project", store).Methods("POST")
	route.HandleFunc("/detail/{id}", detail).Methods("GET")


	// route contact
	route.HandleFunc("/contact", contact).Methods("GET")
	route.HandleFunc("/store-contact", contactStore).Methods("POST")


	fmt.Println("server running on port 5000")
	http.ListenAndServe("localhost:5000", route)

}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}


// ===========+++++++++++++=============

// contact controller

func contact(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")

	theme, err := template.ParseFiles("views/contact/contact.html")
	if err != nil {
		panic(err)
	}
	theme.Execute(res, nil)
}

func contactStore(res http.ResponseWriter, req *http.Request)  {
	err := req.ParseForm()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("name : " + req.PostForm.Get("name"))
	fmt.Println("email : " + req.PostForm.Get("email"))
	fmt.Println("no hp : " + req.PostForm.Get("hp"))
	fmt.Println("subject : " + req.PostForm.Get("select"))
	fmt.Println("pesan : " + req.PostForm.Get("message"))

	http.Redirect(res, req, "/contact", http.StatusMovedPermanently)
}

// controller Blog
func index(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	theme, err := template.ParseFiles("views/blog/index.html")

	if err != nil {
		res.Write([]byte("Hacker jangan menyerang! : " + err.Error()))
		return
	}

	responPenampung := map[string]interface{}{
		"penampungdata" : penampung,

	}
	theme.Execute(res, responPenampung)
}

func add(res http.ResponseWriter, req *http.Request)  {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	theme, err := template.ParseFiles("views/blog/addproject.html")

	if err != nil {
		res.Write([]byte("Hacker jangan menyerang! : " + err.Error()))
		return
	}

	theme.Execute(res, nil)
}

func store(res http.ResponseWriter, req *http.Request)  {
	err := req.ParseForm()

	if err != nil {
		log.Fatal(err)
	}


	// tampung varaibel dulu cuyyy
	title := req.PostForm.Get("title")
	startDate := req.PostForm.Get("start-date")
	endDate := req.PostForm.Get("end-date")
	deskripsi := req.PostForm.Get("desc")
	node := req.PostForm.Get("node")
	react := req.PostForm.Get("react")
	laravel := req.PostForm.Get("laravel")
	golang := req.PostForm.Get("golang")
	gambar := req.PostForm.Get("gambar")

	dataObj := struktur {
		Title : title,
		StartDate : startDate,
		EndDate : endDate,
		Deskripsi : deskripsi,
		Node : node,
		React : react,
		Laravel : laravel,
		Golang : golang,
		Gambar : gambar,

	}
	// push obj ke penampung arr
	penampung = append(penampung,dataObj)

	
	http.Redirect(res, req, "/", http.StatusMovedPermanently)

}

func detail(res http.ResponseWriter, req *http.Request){
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	theme, err := template.ParseFiles("views/blog/detail.html")

	if err != nil {
		res.Write([]byte("Hacker jangan menyerang! :" + err.Error()))
		return
	}
	var blogDetail = struktur{}

	index, _ := strconv.Atoi(mux.Vars(req)["id"])

	for i, data := range penampung{
		if index == i{
			blogDetail = struktur{
				Title : data.Title,
				Deskripsi : data.Deskripsi,
			}
		}
	}

	data := map[string]interface{}{
		"blog" : blogDetail,
	}
	fmt.Println(data)
	theme.Execute(res, data)
}

// ini arr penampung
var penampung = []struktur{}


// ini struktur
type struktur struct{
	Title string
	StartDate string
	EndDate string
	Deskripsi string
	Node string
	React string
	Laravel string
	Golang string
	Gambar string
}

