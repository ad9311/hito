package driverdb

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

// User holds data for a selected user.
type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	Admin     bool      `json:"admin"`
	LastLogin time.Time `json:"lastLogin"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// UserSlice holds multiple users.
type UserSlice struct {
	Entries []User `json:"data"`
	Message string `json:"message"`
}

// LandmarkSlice holds multiple landmarks.
type LandmarkSlice struct {
	Entries []Landmark `json:"data"`
	Message string     `json:"message"`
}

// Landmark holds data for selected landmark.
type Landmark struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	NativeName  string    `json:"nativeName"`
	Type        string    `json:"type"`
	Description string    `json:"description"`
	Continent   string    `json:"continent"`
	Country     string    `json:"country"`
	City        string    `json:"city"`
	Latitude    float64   `json:"latitude"`
	Longitude   float64   `json:"longitude"`
	StartYear   int       `json:"startYear"`
	EndYear     int       `json:"endYear"`
	Length      float64   `json:"length"`
	Width       float64   `json:"width"`
	Height      float64   `json:"height"`
	WikiURL     string    `json:"wikiURL"`
	ImgURL      string    `json:"imgURL"`
	UserID      int       `json:"userID"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// ValidateLoginForm validates that all required fields are present for logging in.
func ValidateLoginForm(r *http.Request) error {
	required := []string{"username", "password"}
	for _, v := range required {
		if r.PostFormValue(v) == "" {
			return errors.New("username or password are required")
		}
	}
	return nil
}

// ValidateAdmin validates that a user's requests are from an administrator user
// and that the currently logged-in user is the actual sender.
func ValidateAdmin(r *http.Request, currentUser User) error {
	formUsername := r.PostFormValue("current-user")
	if formUsername != currentUser.Username {
		return fmt.Errorf("user %s is not allowed to perform this action as user %s", formUsername, currentUser.Username)
	}

	if !currentUser.Admin {
		return fmt.Errorf("user %s is not an administrator user", currentUser.Username)
	}

	return nil
}

// ValidateAddUser validates the required fields for adding a user are present.
func ValidateAddUser(r *http.Request, currentUser User) (User, error) {
	u := User{}
	// required := []string{"name", "username", "password", "password-confirmation", "admin"}
	// posError := "one or more requiered fields are missing"
	// passMismatch := "passwords mismatch"

	// err := ValidateAdmin(r, currentUser)
	// if err != nil {
	// 	return u, err
	// }

	// for _, v := range required {
	// 	if r.PostFormValue(v) == "" {
	// 		return u, errors.New(posError)
	// 	}
	// }

	// password := r.PostFormValue("password")
	// passConfirm := r.PostFormValue("password-confirmation")

	// if password != passConfirm {
	// 	return u, errors.New(passMismatch)
	// }

	// u.Name = r.PostFormValue("name")
	// u.Username = r.PostFormValue("username")
	// admin, err := strconv.ParseBool(r.PostFormValue("admin"))
	// if err != nil {
	// 	return User{}, errors.New("could not process user's privilages")
	// }
	// u.Admin = admin
	// u.Password = password

	return u, nil
}

// ValidateBodyForSingleUsers validates that the ajax request's body sent contains the required data.
func ValidateBodyForSingleUsers(r *http.Request, username, csrfToken string) error {
	type required struct {
		Username  string `json:"username"`
		CSRFToken string `json:"csrf-token"`
	}
	req := required{}
	var unmarshalErr *json.UnmarshalTypeError

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&req)
	if err != nil {
		if errors.As(err, &unmarshalErr) {
			return errors.New("could not process request. Wrong data type provided")
		}
		return errors.New("could not process request")
	}

	if username != req.Username {
		return fmt.Errorf("user %s is not current user or is not logged-in", username)
	}

	if csrfToken != req.CSRFToken {
		return errors.New("CSRFToken is not valid")
	}

	return nil
}

// ValidateNewOrEditLandmark validates that all required fields are present
// for creating or editing a landmark.
// func ValidateNewOrEditLandmark(r *http.Request) (Landmark, error) {
// 	required := []string{
// 		"id",
// 		"name",
// 		"native-name",
// 		"type",
// 		"description",
// 		"continent",
// 		"country",
// 		"city",
// 		"latitude",
// 		"longitude",
// 		"start-year",
// 		"end-year",
// 		"length",
// 		"width",
// 		"height",
// 		"wiki-url",
// 		"img-url",
// 		"user-id",
// 	}

// 	posError := "one or more requiered fields are missing"
// 	lm := Landmark{}

// 	for _, v := range required {
// 		if r.PostFormValue("mode") == "new" {
// 			if r.PostFormValue(v) != "" && r.PostFormValue(v) != "id" {
// 				return lm, errors.New(posError)
// 			}
// 		} else if r.PostFormValue("mode") == "edit" {
// 			if r.PostFormValue(v) != "" && r.PostFormValue(v) != "user-id" {
// 				return lm, errors.New(posError)
// 			}
// 		} else {
// 			return lm, errors.New("could not process form. Wrong format")
// 		}
// 	}

// 	return lm, nil
// }
