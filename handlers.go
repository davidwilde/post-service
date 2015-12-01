package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/davidwilde/post-service/redis_backend"
	"github.com/davidwilde/post-service/repository"
	"github.com/gorilla/mux"
)

var redisClient = redis_backend.NewRedisClient()
var postRepository = repository.NewPostRepository(redisClient)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func VideosCommentsCreate(w http.ResponseWriter, r *http.Request) {
	var comment repository.Comment
	vars := mux.Vars(r)
	videoId := vars["videoId"]
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &comment); err != nil {
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	postRepository.SavePostById(videoId, comment)
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
}

func VideoCommentsShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	videoId := vars["videoId"]
	posts := postRepository.FetchPostsForId(videoId)
	js, err := json.Marshal(posts)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.Write(js)
}
