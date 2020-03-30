package models

import (
	"errors"
	"html"
	"strings"

	"github.com/gofrs/uuid"
	"github.com/silvano-bergamasco/business6sense/backend/utils/dbutils"
)

type StockMarket struct {
	ID         string `json:"id"`
	City       string `json:"city"`
	Nation     string `json:"nation"`
	NationIso2 string `json:"nation_iso2"`
	Active     bool   `json:"active"`
}

func (m *StockMarket) Prepare() error {
	if len(m.ID) == 0 {
		uuid, err := uuid.NewV4()
		if err != nil {
			return err
		}
		m.ID = uuid.String()
	}
	m.City = html.EscapeString(strings.TrimSpace(m.City))
	m.Nation = html.EscapeString(strings.TrimSpace(m.Nation))
	m.NationIso2 = html.EscapeString(strings.TrimSpace(m.NationIso2))

	return nil
}

func (m *StockMarket) Validate() error {
	if m.City == "" {
		return errors.New("Required City")
	}
	if m.Nation == "" {
		return errors.New("Required Nation")
	}
	if len(m.NationIso2) != 2 {
		return errors.New("Required Correct Nation ISO2")
	}
	return nil
}

func (m *StockMarket) FindStockMarketByID(id string) (*StockMarket, error) {
	db, err := dbutils.DbConn()

	if err != nil {
		return &StockMarket{}, err
	}

	selDB, err := db.Query("SELECT * FROM stock_markets WHERE id = ?", id)
	if err != nil {
		return &StockMarket{}, err
	}

	//c = &Company{}
	for selDB.Next() {
		var id, city, nation, nationIso2 string
		var active bool
		err = selDB.Scan(&id, &city, &nation, &nationIso2, &active)
		if err != nil {
			panic(err.Error())
		}
		m.ID = id
		m.City = city
		m.Nation = nation
		m.NationIso2 = nationIso2
		m.Active = active
	}

	defer db.Close()
	return m, nil
}

func (m *StockMarket) InsertMarket() error {
	db, err := dbutils.DbConn()
	if err != nil {
		return err
	}

	insForm, err := db.Prepare("INSERT INTO stock_markets(id, city, nation, nation_iso2, active) VALUES(?,?,?,?,?)")
	if err != nil {
		return err
	}

	insForm.Exec(m.ID, m.City, m.Nation, m.NationIso2, m.Active)

	//log.Println("INSERT: Name: " + name + " | City: " + city)

	defer db.Close()

	return nil
}
