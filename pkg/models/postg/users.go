package postg

import (
	"database/sql"
	"errors"
	"log"
	"os"
	"strings"
	"webui/pkg/models"

	"golang.org/x/crypto/bcrypt"
)

type UserModel struct {
	DB *sql.DB
}

// We'll use the Insert method to add a new record to the users table.
func (m *UserModel) Insert(name, email, password string) error {
	// Create a bcrypt hash of the plain-text password.
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}
	stmt := `INSERT INTO webui.users (name, email, hashed_password, created)
			 VALUES($1, $2, $3, NOW())`

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	// Use the Exec() method to insert the user details and hashed password
	// into the users table.
	_, err = m.DB.Exec(stmt, name, email, string(hashedPassword))
	if err != nil {
		//duplicate key value violates unique constraint
		infoLog.Printf("err !=nil [%s]", err)

		if strings.Contains(err.Error(), "users_uc_email") {
			infoLog.Printf("err duplicate email [%s]", err.Error())
			return models.ErrDuplicateEmail
		}
		return err
	}

	return nil

}

// We'll use the Authenticate method to verify whether a user exists with
// the provided email address and password. This will return the relevant
// user ID if they do.
func (m *UserModel) Authenticate(email, password string) (int, error) {
	// Retrieve the id and hashed password associated with the given email. If no
	// matching email exists, or the user is not active, we return the
	// ErrInvalidCredentials error.
	var id int
	var hashedPassword []byte
	stmt := "SELECT id, hashed_password FROM webui.users WHERE email = $1 AND active = TRUE"
	row := m.DB.QueryRow(stmt, email)
	err := row.Scan(&id, &hashedPassword)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, models.ErrInvalidCredentials
		} else {
			return 0, err
		}
	}
	// Check whether the hashed password and plain-text password provided match.
	// If they don't, we return the ErrInvalidCredentials error.
	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return 0, models.ErrInvalidCredentials
		} else {
			return 0, err
		}
	}
	// Otherwise, the password is correct. Return the user ID.
	return id, nil
}

// We'll use the Get method to fetch details for a specific user based
// on their user ID.
func (m *UserModel) Get(id int) (*models.User, error) {
	u := &models.User{}
	stmt := `SELECT id, name, email, created, active FROM webui.users WHERE id = $1`
	err := m.DB.QueryRow(stmt, id).Scan(&u.ID, &u.Name, &u.Email, &u.Created, &u.Active)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}
	return u, nil
}

func (m *UserModel) UserChangePassword(id int, currentPassword, newPassword string) error {
	var currentHashedPassword []byte
	stmt := `SELECT hashed_password FROM webui.users WHERE id = $1`
	row := m.DB.QueryRow(stmt, id)
	err := row.Scan(&currentHashedPassword)
	if err != nil {
		return err
	}
	err = bcrypt.CompareHashAndPassword(currentHashedPassword, []byte(currentPassword))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return models.ErrInvalidCredentials
		} else {
			return err
		}
	}
	newHashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), 12)
	if err != nil {
		return err
	}
	stmt = `UPDATE webui.users SET hashed_password = $1 WHERE id = $2`
	_, err = m.DB.Exec(stmt, string(newHashedPassword), id)
	return err
}
