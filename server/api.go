package main

import "net/http"

type APIError struct {
	Status int
	Msg    string
}

func (e APIError) Error() string {
	return e.Msg
}

func NotFoundError(e string) APIError {
	return APIError{
		Status: http.StatusNotFound,
		Msg:    e + " not found",
	}
}
