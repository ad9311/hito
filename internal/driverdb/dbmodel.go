package driverdb

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/ad9311/hito/internal/console"
	"golang.org/x/crypto/bcrypt"
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

// ValidateLogin validates user's credentials when logging in.
func (d *DB) ValidateLogin(r *http.Request) (User, error) {
	u := User{}
	required := []string{"username", "password"}
	for _, v := range required {
		if r.PostFormValue(v) == "" {
			return u, errors.New("username or password are required")
		}
	}

	password := r.PostFormValue("password")
	username := r.PostFormValue("username")

	u, err := d.getUserAndComparePasswords(username, password)
	if err != nil {
		console.AssertError(err)
		return u, fmt.Errorf("invalid username or password")
	}

	return u, nil
}

// CreateUser verify a post request form and creates a new user if valid.
func (d *DB) CreateUser(r *http.Request, u User) error {
	err := validateAdmin(r, u)
	if err != nil {
		return err
	}

	missingErr := errors.New("one or more requiered fields are missing")
	failedErr := fmt.Errorf("could not create user %s", r.PostFormValue("username"))
	passMismatch := errors.New("passwords mismatch")
	required := []string{
		"name",
		"username",
		"password",
		"password-confirmation",
		"admin",
	}

	for _, v := range required {
		if r.PostFormValue(v) == "" {
			return missingErr
		}
	}

	password := r.PostFormValue("password")
	passConfirm := r.PostFormValue("password-confirmation")
	if password != passConfirm {
		return passMismatch
	}

	err = d.addUserToDB(r)
	if err != nil {
		console.AssertError(err)
		return failedErr
	}

	return nil
}

// EditUser verify a post request form and creates a new user if valid.
func (d *DB) EditUser(r *http.Request, u User) error {
	err := validateCurrentUser(r, u)
	if err != nil {
		console.AssertError(err)
		return err
	}

	missingErr := errors.New("one or more requiered fields are missing")
	failedErr := fmt.Errorf("could not edit user %s", r.PostFormValue("username"))
	currPasswdMismatch := errors.New("incorrent current password")
	passMismatch := errors.New("new passwords mismatch")
	required := []string{
		"name",
		"username",
		"current-password",
		"new-password",
		"new-password-confirmation",
	}

	for _, v := range required {
		if r.PostFormValue(v) == "" {
			fmt.Println(v)
			return missingErr
		}
	}

	currentPassword, err := d.getUserPasswordByUsername(u.Username)
	if err != nil {
		return failedErr
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(currentPassword),
		[]byte(r.PostFormValue("current-password")),
	)
	if err != nil {
		return currPasswdMismatch
	}

	newPassword := r.PostFormValue("new-password")
	newPassConfirm := r.PostFormValue("new-password-confirmation")
	if newPassword != newPassConfirm {
		return passMismatch
	}

	err = d.editDBUser(r)
	if err != nil {
		console.AssertError(err)
		return failedErr
	}

	return nil
}

// GetUser validates request's origin and returns a slice of users with a single entry
// being that the current user.
func (d *DB) GetUser(r *http.Request, username, csrfToken string) (UserSlice, error) {
	type required struct {
		Username  string `json:"username"`
		CSRFToken string `json:"csrf-token"`
	}

	req := required{}
	us := UserSlice{}
	var unmarshalErr *json.UnmarshalTypeError

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&req)
	if err != nil {
		if errors.As(err, &unmarshalErr) {
			return us, errors.New("could not process request. Wrong data type provided")
		}
		return us, errors.New("could not process request")
	}

	if username != req.Username {
		return us, fmt.Errorf("user %s is not current user or is not logged-in", username)
	}

	if csrfToken != req.CSRFToken {
		return us, errors.New("CSRFToken is not valid")
	}

	u, err := d.getUserByUsername(username)
	if err != nil {
		console.AssertError(err)
		return us, fmt.Errorf("could not find user %s", username)
	}

	us.Entries = append(us.Entries, u)
	us.Message = "success"

	return us, nil
}

// Unexported functions

func validateAdmin(r *http.Request, u User) error {
	formUsername := r.PostFormValue("current-user")
	errS := "is not allowed to perform this action as user"
	if formUsername != u.Username {
		return fmt.Errorf("user %s %s %s", formUsername, errS, u.Username)
	}

	if !u.Admin {
		return fmt.Errorf("user %s is not an administrator user", u.Username)
	}

	return nil
}

func validateCurrentUser(r *http.Request, u User) error {
	formUsername := r.PostFormValue("current-user")
	errS := "is not allowed to perform this action as user"
	if formUsername != u.Username {
		return fmt.Errorf("user %s %s %s", formUsername, errS, u.Username)
	}

	return nil
}
