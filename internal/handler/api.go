package handler

import (
	"encoding/json"
	"net/http"

	"github.com/ad9311/hito/internal/console"
	"github.com/ad9311/hito/internal/dbmodel"
)

// Landmarks returns all landmarks in the database as json.
func Landmarks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	lms, err := config.ConnDB.GetAllLandmarks()
	if err != nil {
		console.AssertError(err)
		lms.Entries = []dbmodel.Landmark{}
		res, err := json.Marshal(lms)
		console.AssertError(err)
		w.Write([]byte(res))
	} else {
		res, err := json.Marshal(lms)
		console.AssertError(err)
		w.Write([]byte(res))
	}
}

// Users returns a selected user from the database, all columns are included expect the password.
func Users(w http.ResponseWriter, r *http.Request) {
	dbmodel.ValidateBodyForUsers(r)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	us, err := config.ConnDB.GetUser("ad9311")
	if err != nil {
		console.AssertError(err)
		us.Entries = []dbmodel.User{}
		res, err := json.Marshal(us)
		console.AssertError(err)
		w.Write([]byte(res))
	} else {
		res, err := json.Marshal(us)
		console.AssertError(err)
		w.Write([]byte(res))
	}
}
