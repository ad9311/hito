package handler

import (
	"net/http"

	"github.com/ad9311/hito/internal/app"
	"github.com/ad9311/hito/internal/console"
	"github.com/ad9311/hito/internal/driverdb"
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
		config.ConnDB.UpdateLastLogin(&data.CurrentUser)
		writeTemplate(w, "home.tmpl.html")
	} else {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}

// PostHome is the home page handler for post action.
func PostHome(w http.ResponseWriter, r *http.Request) {
	if isLoggedIn(r) {
		switch r.PostFormValue("_method") {
		case "post":
			postModel(w, r)
			break
		case "patch":
			patchModel(w, r)
			break
		case "delete":
			deleteModel(w, r)
			break
		default:
			w.Write([]byte("default"))
		}
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
	console.AssertError(r.ParseForm())
	u, err := config.ConnDB.ValidateLogin(r)
	if err != nil {
		console.AssertError(err)
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	} else {
		config.Session.Put(r.Context(), "loggedIn", true)
		data.CurrentUser = u
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

// Logout handles the logout action.
func Logout(w http.ResponseWriter, r *http.Request) {
	_ = config.Session.Destroy(r.Context())
	_ = config.Session.RenewToken(r.Context())
	data.CurrentUser = driverdb.User{}
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

// Unexported functions

func isLoggedIn(r *http.Request) bool {
	loggedIn := config.Session.GetBool(r.Context(), "loggedIn")
	return loggedIn
}

func postModel(w http.ResponseWriter, r *http.Request) {
	switch r.PostFormValue("model") {
	case "user":
		err := config.ConnDB.CreateUser(r, data.CurrentUser)
		if err != nil {
			console.AssertError(err)
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
		break
	case "landmark":
		err := config.ConnDB.AddLandmark(r, data.CurrentUser)
		if err != nil {
			console.AssertError(err)
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
		break
	default:
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func patchModel(w http.ResponseWriter, r *http.Request) {
	switch r.PostFormValue("model") {
	case "user":
		err := config.ConnDB.EditUser(r, data.CurrentUser)
		if err != nil {
			console.AssertError(err)
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
		break
	case "landmark":
		err := config.ConnDB.EditLandmark(r, data.CurrentUser)
		if err != nil {
			console.AssertError(err)
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
		break
	default:
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func deleteModel(w http.ResponseWriter, r *http.Request) {
	switch r.PostFormValue("model") {
	case "user":
		err := config.ConnDB.DeleteUser(r, data.CurrentUser)
		if err != nil {
			console.AssertError(err)
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
		break
	case "landmark":
		err := config.ConnDB.DeleteLandmark(r, data.CurrentUser)
		if err != nil {
			console.AssertError(err)
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
		break
	default:
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
