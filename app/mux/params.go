package mux

import (
	"github.com/gorilla/sessions"
	"github.com/lcaballero/kublai/app/settings"
	"net/http"
	"github.com/lcaballero/kublai/app/handlers"
)

// Params represent the most common parametrs to a function capable
// of servicing an http request.
type Params interface {
	Response() http.ResponseWriter
	Request() *http.Request
	Settings() settings.Settings
	Store() sessions.Store
	Appender() chan *handlers.PubEvent
	Session() (*sessions.Session, error)
	SessionName() string
	TemporaryRedirect(path string) error
}

type params struct {
	res         http.ResponseWriter
	req         *http.Request
	settings    settings.Settings
	store       sessions.Store
	sessionKey  string
	sessionName string
}

func (p *params) Response() http.ResponseWriter {
	return p.res
}
func (p *params) Request() *http.Request {
	return p.req
}
func (p *params) Settings() settings.Settings {
	return p.settings
}
func (p *params) Store() sessions.Store {
	return p.store
}
func (p *params) Session() (*sessions.Session, error) {
	session, err := p.Store().Get(p.Request(), p.SessionName())
	if err != nil {
		return nil, err
	}
	return session, nil
}
func (p *params) SessionName() string {
	return p.sessionName
}
func (p *params) TemporaryRedirect(path string) error {
	http.Redirect(
		p.Response(),
		p.Request(),
		path,
		http.StatusTemporaryRedirect,
	)
	return EndedByRedirect
}
func (p *params) Appender() chan *handlers.PubEvent {
	return nil
}