package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"github.com/gofrs/uuid"
	"github.com/silvano-bergamasco/business6sense/backend/utils/dbutils"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Nickname  string `json:"nickname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (u *User) BeforeSave() error {
	hashedPassword, err := Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) Prepare(action string) error {

	if strings.ToLower(action) == dbutils.Insert {
		uuid, err := uuid.NewV4()
		if err != nil {
			return err
		}
		u.ID = uuid.String()
		u.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	}

	u.Name = html.EscapeString(strings.TrimSpace(u.Name))
	u.Nickname = html.EscapeString(strings.TrimSpace(u.Nickname))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")

	if strings.ToLower(action) == dbutils.Delete {
		u.DeletedAt = time.Now().Format("2006-01-02 15:04:05")
	}

	return nil
}

func (u *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case dbutils.Update:
		if u.Nickname == "" {
			return errors.New("Required Nickname")
		}
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}

		return nil
	case "login":
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil

	default:
		if u.Nickname == "" {
			return errors.New("Required Nickname")
		}
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil
	}
}

func (u *User) InsertUser() error {
	db, err := dbutils.DbConn()
	if err != nil {
		return err
	}

	insForm, err := db.Prepare("INSERT INTO users(id, name, nickname, email, password, created_at, updated_at) VALUES(?,?,?,?,?,?,?)")
	if err != nil {
		return err
	}

	insForm.Exec(u.ID, u.Name, u.Nickname, u.Email, u.Password, u.CreatedAt, u.UpdatedAt)

	//log.Println("INSERT: Name: " + name + " | City: " + city)

	defer db.Close()

	return nil
}

func (u *User) FindUserByID(id string) (*User, error) {
	db, err := dbutils.DbConn()

	if err != nil {
		return &User{}, err
	}

	selDB, err := db.Query("SELECT * FROM users WHERE id = ?", id)
	if err != nil {
		return &User{}, err
	}

	u = &User{}
	for selDB.Next() {
		var id, name, nickname, email, password, createdAt, updatedAt, deletedAt string
		err = selDB.Scan(&id, &name, &nickname, &email, &password, &createdAt, &updatedAt, &deletedAt)
		if err != nil {
			panic(err.Error())
		}
		u.ID = id
		u.Name = name
		u.Nickname = nickname
		u.Email = email
		u.Password = password
		u.CreatedAt = createdAt
		u.UpdatedAt = updatedAt
		u.DeletedAt = deletedAt
	}

	//if gorm.IsRecordNotFoundError(err) {
	//	return &User{}, errors.New("User Not Found")
	//}
	return u, err
}

func (u *User) FindUserByEmail(email string) (error) {
	db, err := dbutils.DbConn()

	if err != nil {
		return err
	}

	selDB, err := db.Query("SELECT * FROM users WHERE email = ? AND deleted_at IS NULL", email)
	if err != nil {
		return err
	}

	u = &User{}
	for selDB.Next() {
		var id, name, nickname, email, password, createdAt, updatedAt, deletedAt string
		err = selDB.Scan(&id, &name, &nickname, &email, &password, &createdAt, &updatedAt, &deletedAt)
		if err != nil {
			panic(err.Error())
		}
		u.ID = id
		u.Name = name
		u.Nickname = nickname
		u.Email = email
		u.Password = password
		u.CreatedAt = createdAt
		u.UpdatedAt = updatedAt
		u.DeletedAt = deletedAt
	}

	if u.ID == "" {
		return errors.New("User Not Found")
	}
	return err
}