package models

import (
	"html"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int16     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
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

func (u *User) Prepare() {
	u.ID = 0
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

func (u *User) FindUserByID(id int16) (*User, error) {
	db, err := dbutils.dbConn()

	if err != nil {
		return &User{}, err
	}

	selDB, err := db.Query("SELECT * FROM USers WHERE id=?", id)
	if err != nil {
		return &User{}, err
	}

	u = &User{}
	for selDB.Next() {
		var id int16
		var name, email, username, password string
		var createdAt, updatedAt, deletedAt time.Time
		err = selDB.Scan(&id, &name, &email, &username, &password, &createdAt, &updatedAt, &deletedAt)
		if err != nil {
			panic(err.Error())
		}
		u.ID = id
		u.Name = name
		u.Email = email
		u.Username = username
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
