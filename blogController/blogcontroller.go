package blogcontroller

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func Index(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	theme, err := template.ParseFiles("views/blog/index.html")

	if err != nil {
		res.Write([]byte("Hacker jangan menyerang! : " + err.Error()))
		return
	}

	theme.Execute(res, nil)
}

func Add(res http.ResponseWriter, req *http.Request)  {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	theme, err := template.ParseFiles("views/blog/addproject.html")

	if err != nil {
		res.Write([]byte("Hacker jangan menyerang! : " + err.Error()))
		return
	}

	theme.Execute(res, nil)
}

func Store(res http.ResponseWriter, req *http.Request)  {
	err := req.ParseForm()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("title : " + req.PostForm.Get("title"))
	fmt.Println("start date : " + req.PostForm.Get("start-date"))
	fmt.Println("end date : " + req.PostForm.Get("end-date"))
	fmt.Println("deskripsi : " + req.PostForm.Get("desc"))
	fmt.Println("node : " + req.PostForm.Get("node"))
	fmt.Println("react : " + req.PostForm.Get("react"))
	fmt.Println("laravel : " + req.PostForm.Get("laravel"))
	fmt.Println("Golang : " + req.PostForm.Get("golang"))
	fmt.Println("Gambar : " + req.PostForm.Get("gambar"))
	
	http.Redirect(res, req, "/add-project", http.StatusMovedPermanently)

}