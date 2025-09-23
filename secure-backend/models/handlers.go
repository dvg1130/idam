package models

import (
	"net/http"
)

type Handlers struct {
	Handler http.HandlerFunc
	Login   http.HandlerFunc
}
