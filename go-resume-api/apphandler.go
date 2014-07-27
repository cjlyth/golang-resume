package api

import (
	"appengine"
	"appengine/user"
	"encoding/json"
	"net/http"
	"time"
)

type AppError struct {
	Error   error
	Message string
	Code    int
}

func NewScope(r *http.Request) *Scope {
	s := new(Scope)
	s.Created = time.Now()
	s.Context = appengine.NewContext(r)
	return s
}

type Scope struct {
	appengine.Context `json:"-"`
	Created           time.Time `json:"asOf"`
	User              *User     `json:"user,omitempty"`
	Links             *Links    `json:"links,omitempty"`
	AppError          *AppError `json:"error,omitempty"`
}

type User struct {
	Email      string `json:"email,omitempty"`
	AuthDomain string `json:"domain,omitempty"`
	Admin      bool   `json:"isAdmin"`
}

type Links struct {
	LoginURL  string `json:"login,omitempty"`
	LogoutURL string `json:"logout,omitempty"`
}

type AppHandler func(http.ResponseWriter, *Scope)

func (scope *Scope) HasError() bool {
	return scope.AppError == nil
}
func (scope *Scope) SendResponse(w http.ResponseWriter) {
	e := scope.AppError
	if e == nil {
		json.NewEncoder(w).Encode(scope)
		return
	}
	if e.Error != nil {
		scope.Errorf("AppError while handling a request \n%#v\n", e.Error)
	}
	http.Error(w, e.Message, e.Code)
}

func (fn AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	scope := NewScope(r)
	defer scope.SendResponse(w)
	scope.LoadUser().LoadLinks()
	fn(w, scope)

}
func (s *Scope) LoadUser() *Scope {
	s.trace("LoadUser")
	u := user.Current(s.Context)
	if u != nil {
		s.User = &User{u.Email, u.AuthDomain, u.Admin}
	}
	return s
}
func (s *Scope) LoadLinks() *Scope {
	s.trace("LoadLinks")
	l := new(Links)
	if s.User == nil {
		loginURL, err := user.LoginURL(s.Context, "/")
		if err != nil {
			s.AppError = &AppError{err, "Error getting login url!", 500}
		}
		l.LoginURL = loginURL
	} else {
		logoutURL, err := user.LogoutURL(s.Context, "/")
		if err != nil {
			s.AppError = &AppError{err, "Error getting logout url!", 500}
		}
		l.LogoutURL = logoutURL
	}
	if s.HasError() {
		s.Links = l
	}
	return s
}
