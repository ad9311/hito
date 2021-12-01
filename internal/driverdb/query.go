package driverdb

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/ad9311/hito/internal/console"
	"github.com/ad9311/hito/internal/dbmodel"
	"golang.org/x/crypto/bcrypt"
)

// ValidateLogin validates the user's username and password at login.
func (d *DB) ValidateLogin(username, password string) (dbmodel.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	user := dbmodel.User{}
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

// UpdateLastLogin updates the last login column from the users table.
func (d *DB) UpdateLastLogin(u *dbmodel.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "update users set last_login = $1 where id = $2"
	posError := "could not update last login for user"
	dt, err := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	if err != nil {
		console.AssertError(err)
		return fmt.Errorf("%s %s", posError, u.Username)
	}

	_, err = d.SQL.ExecContext(ctx, query, dt, u.ID)
	if err != nil {
		console.AssertError(err)
		return fmt.Errorf("%s %s", posError, u.Username)
	}
	u.LastLogin = dt

	return nil
}

// AddOrEditUser add a new user to the data base.
func (d *DB) AddOrEditUser(u dbmodel.User, password, mode string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	dt, errDt := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	if errDt != nil {
		console.AssertError(errDt)
	}

	posError := "could not add or edit user"
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		console.AssertError(err)
		return fmt.Errorf("%s, %s", posError, u.Name)
	}

	query1 := `insert into users
	("name",username,password,admin,last_login,created_at,updated_at)
	values ($1,$2,$3,$4,$5,$6,$7)
	`
	query2 := `update users
	set "name"=$1,username=$2,"password"=$3,updated_at=$4
	where id=$5;
	`

	if mode == "new" {
		_, err := d.SQL.ExecContext(
			ctx,
			query1,
			u.Name,
			u.Username,
			hashPassword,
			u.Admin,
			"0001-01-01 01:00:00",
			dt,
			dt,
		)
		if err != nil {
			console.AssertError(err)
			return fmt.Errorf("%s %s", posError, u.Name)
		}
	} else if mode == "edit" {
		_, err := d.SQL.ExecContext(
			ctx,
			query2,
			u.Name,
			u.Username,
			hashPassword,
			dt,
			u.ID,
		)
		if err != nil {
			console.AssertError(err)
			return fmt.Errorf("%s %s", posError, u.Name)
		}
	}

	return nil
}

// DeleteUser deletes a user from the database.
func (d *DB) DeleteUser(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	posError := "could not find user by id"
	u, err := d.findUserByID(id)
	if err != nil {
		console.AssertError(err)
		return fmt.Errorf("%s %d", posError, id)
	}
	if u.ID < 1 {
		err = fmt.Errorf("%s %d", posError, id)
		console.AssertError(err)
		return err
	}

	query := `delete from users where id=$1`

	_, err = d.SQL.ExecContext(ctx, query, u.ID)
	if err != nil {
		console.AssertError(err)
		return fmt.Errorf("could not delete from database user %s", u.Name)
	}
	return nil
}

// GetAllLandmarks returns all the landmarks in the database.
func (d *DB) GetAllLandmarks() (dbmodel.LandmarkSlice, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	lms := dbmodel.LandmarkSlice{}
	lm := dbmodel.Landmark{}
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

// AddOrEditLandmark adds a new landmark to the database.
func (d *DB) AddOrEditLandmark(lm dbmodel.Landmark, mode string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	posError := "could not add lanmark to database"
	dt, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	query1 := `
	insert into landmarks ("name",native_name,"type",description,continent,country,
	city,latitude,longitude,start_year,end_year,lengths,
	width,height,wiki_url,img_url,user_id,created_at,updated_at)
	values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19)
	`

	query2 := `
	update landmarks set "name"=$1,native_name=$2,"type"=$3,description=$4,continent=$5,country=$6,
	city=$7,latitude=$8,longitude=$9,start_year=$10,end_year=$11,lengths=$12,width=$13,height=$14,
	wiki_url=$15,img_url=$16,updated_at=$17 where id = $18
	`

	err := errors.New("")

	if mode == "new" {
		_, err = d.SQL.ExecContext(
			ctx,
			query1,
			lm.Name,
			lm.NativeName,
			lm.Type,
			lm.Description,
			lm.Continent,
			lm.Country,
			lm.City,
			lm.Latitude,
			lm.Longitude,
			lm.StartYear,
			lm.EndYear,
			lm.Length,
			lm.Width,
			lm.Height,
			lm.WikiURL,
			lm.ImgURL,
			lm.UserID,
			dt,
			dt,
		)
	}

	if mode == "edit" {
		_, err2 := d.SQL.ExecContext(
			ctx,
			query2,
			lm.Name,
			lm.NativeName,
			lm.Type,
			lm.Description,
			lm.Continent,
			lm.Country,
			lm.City,
			lm.Latitude,
			lm.Longitude,
			lm.StartYear,
			lm.EndYear,
			lm.Length,
			lm.Width,
			lm.Height,
			lm.WikiURL,
			lm.ImgURL,
			dt,
			lm.ID,
		)
		console.AssertError(err2)
	}

	if err != nil && err.Error() != "" {
		console.AssertError(err)
		return errors.New(posError)
	}

	return nil
}

// DeleteLandmark deletes a landmark from the database.
func (d *DB) DeleteLandmark(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	posError := "could not find landmark by id"
	lm, err := d.findLandmarkByID(id)
	if err != nil {
		console.AssertError(err)
		return fmt.Errorf("%s %d", posError, id)
	}
	if lm.ID < 1 {
		err = fmt.Errorf("%s %d", posError, id)
		console.AssertError(err)
		return err
	}

	query := "delete from landmarks where id = $1"

	_, err = d.SQL.ExecContext(ctx, query, lm.ID)
	if err != nil {
		console.AssertError(err)
		return fmt.Errorf("could not delete from database landmark %s", lm.Name)
	}

	return nil
}

// Non imported functions.

func (d *DB) findLandmarkByID(id int) (dbmodel.Landmark, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	lm := dbmodel.Landmark{}
	query := "select id, name from landmarks where id = $1"

	row := d.SQL.QueryRowContext(ctx, query, id)
	err := row.Scan(&lm.ID, &lm.Name)
	if err != nil {
		return lm, err
	}

	return lm, nil
}

func (d *DB) findUserByID(id int) (dbmodel.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	u := dbmodel.User{}
	query := "select id, name from users where id = $1"

	row := d.SQL.QueryRowContext(ctx, query, id)
	err := row.Scan(&u.ID, &u.Name)
	if err != nil {
		return u, err
	}

	return u, nil
}
