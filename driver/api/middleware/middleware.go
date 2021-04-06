package middleware

import "net/http"

type Middleware interface {
	Execute() http.Handler
}
