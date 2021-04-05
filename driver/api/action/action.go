package action

import "net/http"

type Action interface {
	Execute() http.HandlerFunc
}
