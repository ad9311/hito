package dbmodel

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
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
func ValidateAdmin(r *http.Request, currentUser string) error {
	formUser := r.PostFormValue("current-user")
	if formUser != currentUser {
		return fmt.Errorf("user %s is not allowed to perform this action as user %s", formUser, currentUser)
	}

	admin, err := strconv.ParseBool(r.PostFormValue("admin"))
	if err != nil || !admin {
		return fmt.Errorf("user %s is not an administrator user", currentUser)
	}

	return nil
}

// ValidateNewOrEditLandmark validates that all required fields are present
// for creating or editing a landmark.
func ValidateNewOrEditLandmark(r *http.Request) (Landmark, error) {
	required := []string{
		"id",
		"name",
		"native-name",
		"type",
		"description",
		"continent",
		"country",
		"city",
		"latitude",
		"longitude",
		"start-year",
		"end-year",
		"length",
		"width",
		"height",
		"wiki-url",
		"img-url",
		"user-id",
	}

	posError := "one or more requiered fields are missing"
	lm := Landmark{}

	for _, v := range required {
		if r.PostFormValue("mode") == "new" {
			if r.PostFormValue(v) != "" && r.PostFormValue(v) != "id" {
				return lm, errors.New(posError)
			}
		} else if r.PostFormValue("mode") == "edit" {
			if r.PostFormValue(v) != "" && r.PostFormValue(v) != "user-id" {
				return lm, errors.New(posError)
			}
		} else {
			return lm, errors.New("could not process form. Wrong format")
		}
	}

	return lm, nil
}
