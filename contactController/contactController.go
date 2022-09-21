package contactController

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func Index(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")

	theme, err := template.ParseFiles("views/contact/contact.html")
	if err != nil {
		panic(err)
	}
	theme.Execute(res, nil)
}

func Store(res http.ResponseWriter, req *http.Request)  {
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