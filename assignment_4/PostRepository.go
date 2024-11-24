package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	var posts []Post

	GetConnection().Find(&posts)

	if len(posts) == 0 {
		http.Error(w, http.StatusText(404), http.StatusNotFound)
		log.Println("Posts don't exist or unreachable")
		return
	}

	err := json.NewEncoder(w).Encode(posts)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
	} else {
		w.WriteHeader(http.StatusOK)
		log.Println("Posts retrieved successfully")
	}
}

func GetPostById(w http.ResponseWriter, r *http.Request) {
	httpParams := mux.Vars(r)
	PostId := httpParams["id"]

	if !VaidateId(PostId) {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		log.Println("Identificator should be numeric")
		return
	}

	var CheckPost Post

	GetConnection().First(&CheckPost, "id = ?", PostId)

	if CheckPost == (Post{}) {
		http.Error(w, http.StatusText(404), http.StatusNotFound)
		log.Println("Post can not be found")
		return
	}

	err := json.NewEncoder(w).Encode(CheckPost)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
	} else {
		w.WriteHeader(http.StatusOK)
		log.Println("Post retrieved successfully")
	}
}

func AddPost(w http.ResponseWriter, r *http.Request) {
	var post Post
	err := json.NewDecoder(r.Body).Decode(&post)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}
	if !ValidateString(post.Title) {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		log.Println("The content of title may contain SQL injection code")
		return
	}
	if !ValidateString(post.Content) {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		log.Println("The content of content field may contain SQL injection code")
		return
	}
	if !ValidateString(post.PostedAt) {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		log.Println("The content of post date may contain SQL injection code")
		return
	}

	// This will be checked by middleware
	AuthorId := r.Context().Value("user_id").(uint)
	post.AuthorId = AuthorId
	GetConnection().Create(&post)
	err = json.NewEncoder(w).Encode(post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
	} else {
		w.WriteHeader(http.StatusCreated)
		log.Println("Post has successfully been added")
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	httpParams := mux.Vars(r)
	PostId := httpParams["id"]
	if !VaidateId(PostId) {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		log.Println("Identificator should be numeric")
		return
	}
	var CheckPost Post
	GetConnection().First(&CheckPost, "id = ?", PostId)
	if CheckPost == (Post{}) {
		http.Error(w, http.StatusText(404), http.StatusNotFound)
		log.Println("Post can not be found")
		return
	}
	var UpdatePost Post
	if !ValidateString(UpdatePost.Title) {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		log.Println("The content of title may contain SQL injection code")
		return
	}
	if !ValidateString(UpdatePost.Content) {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		log.Println("The content may contain SQL injection code")
		return
	}
	if !ValidateString(UpdatePost.PostedAt) {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		log.Println("The content of post date may contain SQL injection code")
		return
	}
	err := json.NewDecoder(r.Body).Decode(&UpdatePost)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}
	GetConnection().Updates(&UpdatePost)
	err = json.NewEncoder(w).Encode(UpdatePost)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
	} else {
		w.WriteHeader(http.StatusOK)
		log.Println("Post updated successfully")
	}

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	httpParams := mux.Vars(r)
	PostId := httpParams["id"]

	if !VaidateId(PostId) {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		log.Println("Identificator should be numeric")
		return
	}

	var CheckPost Post

	GetConnection().First(&CheckPost, "id = ?", PostId)

	if CheckPost == (Post{}) {
		http.Error(w, http.StatusText(404), http.StatusNotFound)
		log.Println("Post can not be found")
		return
	}

	GetConnection().Delete(&Post{}, CheckPost.Id)

	log.Println("Post deleted successfully")
}
