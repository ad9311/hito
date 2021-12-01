package handler

import (
	"encoding/json"
	"net/http"

	"github.com/ad9311/hito/internal/console"
	"github.com/ad9311/hito/internal/dbmodel"
)

// Landmarks returns all landmarks in the database as json
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
