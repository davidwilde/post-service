package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"VideoCommentsShow",
		"GET",
		"/video/{videoId}/comments",
		VideoCommentsShow,
	},
	Route{
		"VideoCommentCreate",
		"POST",
		"/video/{videoId}/comments",
		VideosCommentsCreate,
	},
}
