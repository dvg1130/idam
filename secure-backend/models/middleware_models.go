package models

import "net/http"

type StatusRecorder struct {
	http.ResponseWriter
	status int
}
