package models

import (
	"net/http"
)

type Handlers struct {
	Login http.HandlerFunc
}
