package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/gofrs/uuid"
	"github.com/silvano-bergamasco/business6sense/backend/utils/dbutils"
)

type Company struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	URL       string `json:"url"`
	Sector    string `json:"sector"`
	Industry  string `json:"industry"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}

func (c *Company) New(name, url, sector, industry string) error {
	uuid, err := uuid.NewV4()
	if err != nil {
		return err
	}
	c.ID = uuid.String()
	c.Name = name
	c.URL = url
	c.Sector = sector
	c.Industry = industry
	c.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	c.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")

	return nil
}

func (c *Company) Prepare() error {
	if len(c.ID) == 0 {
		uuid, err := uuid.NewV4()
		if err != nil {
			return err
		}
		c.ID = uuid.String()
	}
	c.Name = html.EscapeString(strings.TrimSpace(c.Name))
	c.URL = html.EscapeString(strings.TrimSpace(c.URL))
	c.Sector = html.EscapeString(strings.TrimSpace(c.Sector))
	c.Industry = html.EscapeString(strings.TrimSpace(c.Industry))
	c.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	c.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")

	return nil
}

func (c *Company) Validate() error {

	if c.Name == "" {
		return errors.New("Required Name")
	}
	if c.URL == "" {
		return errors.New("Required URL")
	}
	if c.Sector == "" {
		return errors.New("Required Sector")
	}
	if c.Industry == "" {
		return errors.New("Required Industry")
	}
	return nil
}

func (c *Company) FindCompanyByID(id string) (*Company, error) {
	db, err := dbutils.DbConn()

	if err != nil {
		return &Company{}, err
	}

	selDB, err := db.Query("SELECT * FROM companies WHERE id = ?", id)
	if err != nil {
		return &Company{}, err
	}

	//c = &Company{}
	for selDB.Next() {
		var id, name, url, sector, industry, createdAt, updatedAt, deletedAt string
		err = selDB.Scan(&id, &name, &url, &sector, &industry, &createdAt, &updatedAt, &deletedAt)
		if err != nil {
			panic(err.Error())
		}
		c.ID = id
		c.Name = name
		c.URL = url
		c.Sector = sector
		c.Industry = industry
		c.CreatedAt = createdAt
		c.UpdatedAt = updatedAt
		c.DeletedAt = deletedAt
	}

	defer db.Close()
	return c, nil
}

func (c *Company) InsertCompany() error {
	db, err := dbutils.DbConn()
	if err != nil {
		return err
	}

	insForm, err := db.Prepare("INSERT INTO companies(id, name, url, sector, industry, created_at, updated_at) VALUES(?,?,?,?,?,?,?)")
	if err != nil {
		return err
	}

	insForm.Exec(c.ID, c.Name, c.URL, c.Sector, c.Industry, c.CreatedAt, c.UpdatedAt)

	//log.Println("INSERT: Name: " + name + " | City: " + city)

	defer db.Close()

	return nil
}
