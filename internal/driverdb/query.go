package driverdb

import (
	"context"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// ValidateLogin validates the user's username and password at login.
func (d *DB) ValidateLogin(username, password string) (User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	user := User{}
	var storedPassword string
	query := "select * from users where username = $1"
	row := d.SQL.QueryRowContext(ctx, query, username)
	err := row.Scan(
		&user.ID,
		&user.Name,
		&user.Username,
		&storedPassword,
		&user.Admin,
		&user.LastLogin,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return user, errors.New("invalid username or password")
	}
	err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(password))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return user, errors.New("invalid username or password")
	} else if err != nil {
		return user, err
	}

	return user, nil
}
