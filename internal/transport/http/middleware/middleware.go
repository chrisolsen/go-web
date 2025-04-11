package middleware

import (
	"net/http"
)

type mw = func(http.Handler) http.Handler

type Middleware struct {
	fns []mw
}

func New(fns ...mw) *Middleware {
	return &Middleware{
		fns: fns,
	}
}

func (m *Middleware) Then(last http.Handler) http.Handler {
	for i := len(m.fns) - 1; i >= 0; i-- {
		last = m.fns[i](last)
	}
	return last
}
