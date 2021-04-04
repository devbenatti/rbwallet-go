package action

import "net/http"

type HandlerFunc func(w http.ResponseWriter, r *http.Request)

type Action interface {
	Handle() HandlerFunc
}
