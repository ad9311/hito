package driverdb

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/ad9311/hito/internal/console"
	"golang.org/x/crypto/bcrypt"
)

// UpdateLastLogin updates the currently logged-in user's last login date.
func (d *DB) UpdateLastLogin(u *User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	posErr := fmt.Errorf("could not update last login for user %s", u.Username)
	query := "update users set last_login = $1 where id = $2"
	dt, err := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	if err != nil {
		console.AssertError(err)
		return posErr
	}

	_, err = d.SQL.ExecContext(ctx, query, dt, u.ID)
	if err != nil {
		console.AssertError(err)
		return posErr
	}
	u.LastLogin = dt

	return nil
}

// GetLandmarks returns all the landmarks in the database.
func (d *DB) GetLandmarks() (LandmarkSlice, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	lms := LandmarkSlice{}
	lm := Landmark{}
	posError := "could not get landmarks from database"
	query := "select * from landmarks"

	rows, err := d.SQL.QueryContext(ctx, query)
	if err != nil {
		rows.Close()
		console.AssertError(err)
		lms.Message = posError
		return lms, errors.New(posError)
	}

	for rows.Next() {
		err := rows.Scan(
			&lm.ID,
			&lm.Name,
			&lm.NativeName,
			&lm.Type,
			&lm.Description,
			&lm.Continent,
			&lm.Country,
			&lm.City,
			&lm.Latitude,
			&lm.Longitude,
			&lm.StartYear,
			&lm.EndYear,
			&lm.Length,
			&lm.Width,
			&lm.Height,
			&lm.WikiURL,
			&lm.ImgURL,
			&lm.UserID,
			&lm.CreatedAt,
			&lm.UpdatedAt,
		)
		if err != nil {
			rows.Close()
			console.AssertError(err)
			lms.Message = posError
			return lms, errors.New(posError)
		}
		lms.Entries = append(lms.Entries, lm)
	}
	lms.Message = "success"
	return lms, nil
}

// Unexported functions

func (d *DB) getUserAndComparePasswords(username, password string) (User, error) {
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
		return user, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(password))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return user, err
	} else if err != nil {
		return user, err
	}

	return user, nil
}

func (d *DB) getUserByUsername(username string) (User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	u := User{}
	query := `select id, "name",username,admin,last_login,
	created_at,updated_at from users where username = $1`

	row := d.SQL.QueryRowContext(ctx, query, username)
	err := row.Scan(
		&u.ID,
		&u.Name,
		&u.Username,
		&u.Admin,
		&u.LastLogin,
		&u.CreatedAt,
		&u.UpdatedAt,
	)

	if err != nil {
		return u, err
	}
	return u, nil
}

func (d *DB) getUserPasswordByUsername(username string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var password string
	query := `select password from users where username = $1`

	row := d.SQL.QueryRowContext(ctx, query, username)
	err := row.Scan(&password)

	if err != nil {
		return password, err
	}
	return password, nil
}

func (d *DB) addUserToDB(r *http.Request) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	dt, errDt := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	if errDt != nil {
		console.AssertError(errDt)
	}

	hashPassword, err := bcrypt.GenerateFromPassword(
		[]byte(r.PostFormValue("password")),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return err
	}

	query := `insert into users
	("name",username,password,admin,last_login,created_at,updated_at)
	values ($1,$2,$3,$4,$5,$6,$7)
	`

	_, err = d.SQL.ExecContext(
		ctx,
		query,
		r.PostFormValue("name"),
		r.PostFormValue("username"),
		hashPassword,
		r.PostFormValue("admin"),
		"0001-01-01 01:00:00",
		dt,
		dt,
	)

	if err != nil {
		return err
	}

	return nil
}

func (d *DB) editDBUser(r *http.Request) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	dt, errDt := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	if errDt != nil {
		console.AssertError(errDt)
	}

	hashPassword, err := bcrypt.GenerateFromPassword(
		[]byte(r.PostFormValue("new-password")),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return err
	}

	query := `update users set "name"=$1,
	username=$2,password=$3,updated_at=$4 where username=$5`

	_, err = d.SQL.ExecContext(
		ctx,
		query,
		r.PostFormValue("name"),
		r.PostFormValue("username"),
		hashPassword,
		dt,
		r.PostFormValue("current-user"),
	)

	if err != nil {
		return err
	}

	return nil
}

func (d *DB) deleteUserFromDB(u User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "delete from users where username=$1"
	_, err := d.SQL.ExecContext(ctx, query, u.Username)
	if err != nil {
		return err
	}

	return nil
}

func (d *DB) addLandmarkToDB(r *http.Request, u User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	dt, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	query := `
	insert into landmarks ("name",native_name,"type",description,continent,country,
	city,latitude,longitude,start_year,end_year,lengths,
	width,height,wiki_url,img_url,user_id,created_at,updated_at)
	values ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19)
	`

	err := errors.New("")

	_, err = d.SQL.ExecContext(
		ctx,
		query,
		r.PostFormValue("name"),
		r.PostFormValue("native-name"),
		r.PostFormValue("type"),
		r.PostFormValue("description"),
		r.PostFormValue("continent"),
		r.PostFormValue("country"),
		r.PostFormValue("city"),
		r.PostFormValue("latitude"),
		r.PostFormValue("longitude"),
		r.PostFormValue("start-year"),
		r.PostFormValue("end-year"),
		r.PostFormValue("length"),
		r.PostFormValue("width"),
		r.PostFormValue("height"),
		r.PostFormValue("wiki-url"),
		r.PostFormValue("img-url"),
		u.ID,
		dt,
		dt,
	)

	if err != nil {
		return err
	}

	return nil
}

func (d *DB) editLandmarkToDB(r *http.Request) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	dt, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	query := `
	update landmarks set "name"=$1,native_name=$2,"type"=$3,description=$4,continent=$5,country=$6,
	city=$7,latitude=$8,longitude=$9,start_year=$10,end_year=$11,lengths=$12,width=$13,height=$14,
	wiki_url=$15,img_url=$16,updated_at=$17 where id=$18
	`

	err := errors.New("")

	_, err = d.SQL.ExecContext(
		ctx,
		query,
		r.PostFormValue("name"),
		r.PostFormValue("native-name"),
		r.PostFormValue("type"),
		r.PostFormValue("description"),
		r.PostFormValue("continent"),
		r.PostFormValue("country"),
		r.PostFormValue("city"),
		r.PostFormValue("latitude"),
		r.PostFormValue("longitude"),
		r.PostFormValue("start-year"),
		r.PostFormValue("end-year"),
		r.PostFormValue("length"),
		r.PostFormValue("width"),
		r.PostFormValue("height"),
		r.PostFormValue("wiki-url"),
		r.PostFormValue("img-url"),
		dt,
		r.PostFormValue("landmark-id"),
	)

	if err != nil {
		return err
	}

	return nil
}
