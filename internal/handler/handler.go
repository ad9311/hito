package handler

import (
	"net/http"

	"github.com/ad9311/hito/internal/app"
	"github.com/ad9311/hito/internal/console"
	"github.com/ad9311/hito/internal/dbmodel"
	"github.com/justinas/nosurf"
)

var config *app.InitConfig
var data *app.InitData

// New saves the app configuration into a local struct.
func New(cfg *app.InitConfig, dat *app.InitData) {
	config = cfg
	data = dat
}

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	if isLoggedIn(r) {
		writeTemplate(w, "home.tmpl.html")
	} else {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}

// PostHome is the home page handler for post action.
func PostHome(w http.ResponseWriter, r *http.Request) {
	if isLoggedIn(r) {
	} else {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}

// Login is the login page handler
func Login(w http.ResponseWriter, r *http.Request) {
	if isLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		data.CSRFToken = nosurf.Token(r)
		writeTemplate(w, "login.tmpl.html")
	}
}

// PostLogin is the login page handler for post action.
func PostLogin(w http.ResponseWriter, r *http.Request) {
	_ = config.Session.RenewToken(r.Context())
	err := r.ParseForm()
	console.AssertError(err)
	err = dbmodel.ValidateLoginForm(r)
	if err != nil {
		console.AssertError(err)
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	data.CurrentUser, err = config.ConnDB.ValidateLogin(username, password)
	if err != nil {
		console.AssertError(err)
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	} else {
		config.Session.Put(r.Context(), "loggedIn", true)
		err = config.ConnDB.UpdateLastLogin(&data.CurrentUser)
		console.AssertError(err)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

// Logout handles the logout action.
func Logout(w http.ResponseWriter, r *http.Request) {
	_ = config.Session.Destroy(r.Context())
	_ = config.Session.RenewToken(r.Context())
	data.CurrentUser = dbmodel.User{}
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

// Unexported functions

func isLoggedIn(r *http.Request) bool {
	loggedIn := config.Session.GetBool(r.Context(), "loggedIn")
	return loggedIn
}