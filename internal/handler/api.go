package handler

import (
	"net/http"
)

// Landmarks returns all landmarks in the database as json.
func Landmarks(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json; charset=utf-8")
	// lms, err := config.ConnDB.GetLandmarks()
	// if err != nil {
	// 	console.AssertError(err)
	// 	lms.Entries = []driverdb.Landmark{}
	// 	res, err := json.Marshal(lms)
	// 	console.AssertError(err)
	// 	w.Write([]byte(res))
	// } else {
	// 	res, err := json.Marshal(lms)
	// 	console.AssertError(err)
	// 	w.Write([]byte(res))
	// }
}

// PostUsers returns a selected user from the database, all columns are included expect the password.
func PostUsers(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json; charset=utf-8")
	// err := driverdb.ValidateBodyForSingleUsers(r, data.CurrentUser.Username, data.CSRFToken)
	// if err != nil {
	// 	console.AssertError(err)
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	w.Write([]byte(fmt.Sprintf("{\"data\":[],\"message\":\"%s\"}", err.Error())))
	// } else {
	// 	us, err := config.ConnDB.GetUser(data.CurrentUser.Username)
	// 	if err != nil {
	// 		console.AssertError(err)
	// 		us.Entries = []driverdb.User{}
	// 		res, err := json.Marshal(us)
	// 		console.AssertError(err)
	// 		w.WriteHeader(http.StatusBadRequest)
	// 		w.Write([]byte(res))
	// 	} else {
	// 		w.WriteHeader(http.StatusAccepted)
	// 		res, err := json.Marshal(us)
	// 		console.AssertError(err)
	// 		w.Write([]byte(res))
	// 	}
	// }
}
