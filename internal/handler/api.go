package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ad9311/hito/internal/console"
)

// Landmarks returns all landmarks in the database as json.
func Landmarks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	lms, err := config.ConnDB.GetLandmarks()
	if err != nil {
		console.AssertError(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("{\"data\":[],\"message\":\"%s\"}", err.Error())))
	} else {
		res, err := json.Marshal(lms)
		console.AssertError(err)
		w.Write([]byte(res))
	}
}

// PostCurrentUser returns a selected user from the database,
// all columns are included expect the password.
func PostCurrentUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	us, err := config.ConnDB.GetUser(r, data.CurrentUser.Username, data.CSRFToken)
	if err != nil {
		console.AssertError(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("{\"data\":[],\"message\":\"%s\"}", err.Error())))
	} else {
		res, err := json.Marshal(us)
		console.AssertError(err)
		w.Write(res)
	}
}
