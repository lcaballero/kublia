package mux

import (
	"fmt"
	gmux "github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/lcaballero/kublai/app/settings"
	"net/http"
)

var NoSessionFoundError = fmt.Errorf("no session found error")
var EndedByRedirect = fmt.Errorf("ended by redirect")

// Handle type is function capable of servicing an http.Request
type Handler func(Params) error

// NoopHandler accepts the given params and return nil.
func NoopHandler(Params) error {
	return nil
}

type ErrHandler func(error, Params)

func DefaultErrorHandler(err error, p Params) {
	msg := fmt.Sprintf("%s path: %s", err.Error(), p.Request().URL.Path)
	http.Error(p.Response(), msg, http.StatusInternalServerError)
}

// Mux is a decorator of Handlers to provide a common signature for
// providing common http request contexts.
type Mux struct {
	router  *gmux.Router
	onError ErrHandler
	store   *sessions.CookieStore
}

const (
	SessionName = "sid"
	sessionKey  = "some-secret-key"
)

// NewMux allocates a new Mux route decorator.
func NewMux() *Mux {
	return &Mux{
		store:   sessions.NewCookieStore([]byte(sessionKey)),
		router:  gmux.NewRouter(),
		onError: DefaultErrorHandler,
	}
}

// Handle registers the path with the http Handler provider.
func (m *Mux) Handle(path string, s settings.Settings, h ...Handler) *gmux.Route {
	return m.router.HandleFunc(path, m.WithConfig(s, h...))
}

// OnError registers the given ErrHandler to use when an internal
// error is detected.
func (m *Mux) OnError(h ErrHandler) {
	if h != nil {
		h = DefaultErrorHandler
	}
	m.onError = h
}

// WithConfig produces an http.HandlerFunc over the given config and
// request Handler.
func (m *Mux) WithConfig(s settings.Settings, handlers ...Handler) http.HandlerFunc {
	newHandler := func(res http.ResponseWriter, req *http.Request) {
		newParams := &params{
			res:         res,
			req:         req,
			settings:    s,
			store:       m.store,
			sessionName: SessionName,
		}

		for _, handler := range handlers {
			err := handler(newParams)
			if err == EndedByRedirect {
				break
			}
			if err != nil {
				m.onError(err, newParams)
				break
			}
		}
	}
	return newHandler
}

// ToHandler turns this Mux into an http.Handler which can serve the
// registered paths when used as parameter to http.ListenAndServe.
func (m *Mux) ToHandler() http.Handler {
	return m.router
}
