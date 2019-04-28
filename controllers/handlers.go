package controllers

import (
    "encoding/json"
    "fmt"
    "net/http"
    "io"
    "io/ioutil"
    "GO_REST_API/utils"
    "GO_REST_API/models"
    "GO_REST_API/repo"
    "github.com/gorilla/mux"
)

func Ping(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Ping")
}

func PostBlog(w http.ResponseWriter, r *http.Request){
	//fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	var blog models.Blog
    body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
    if err != nil {
        panic(err)
    }
    if err := r.Body.Close(); err != nil {
        panic(err)
    }
    if err := json.Unmarshal(body, &blog); err != nil {
        panic(err)
        w.WriteHeader(422) // unprocessable entity
        utils.Respond(w, utils.Message(false,"Error while processing the entity"))
    }
 
    if err := repo.CreateBlog(blog); err !=nil {
        panic(err)
        utils.Respond(w, utils.Message(false, "error when posting Blog"))
    }
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusCreated)
    utils.Respond(w, utils.Message(true,"Success"))
}

func GetBlog(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
    blogId := vars["title"]
    fmt.Println(" title is :"+ blogId)
	blog := repo.FindBlog(blogId)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(blog); err != nil {
		panic(err)
	}
}

func GetBlogs(w http.ResponseWriter, r *http.Request){
	blogs := models.Blogs{
		models.Blog{Title : "Title 1",  Body : "My First blog"},
		models.Blog{Title : "Title 2",  Body : "My Second blog"},
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(blogs); err != nil {
		panic(err)
	}
}