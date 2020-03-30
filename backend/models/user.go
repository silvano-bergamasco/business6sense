package models

import (
	"html"
	"strings"
	"time"

	"github.com/silvano-bergamasco/business6sense/backend/utils/dbutils"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int16  `json:"id"`
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

func (u *User) Prepare() {
	u.ID = 0
	u.Nickname = html.EscapeString(strings.TrimSpace(u.Nickname))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	u.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
}

func (u *User) FindUserByID(id int16) (*User, error) {
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
		var id int16
		var name, nickname, email, password, createdAt, updatedAt, deletedAt string
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
