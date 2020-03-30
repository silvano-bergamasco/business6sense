package models

import (
	"errors"

	"github.com/gofrs/uuid"
	"github.com/silvano-bergamasco/business6sense/backend/utils/dbutils"
)

type Stock struct {
	ID           string `json:"id"`
	CompanyID    string `json:"company_id"`
	StocksNumber int16  `json:"stocks_number"`
	ValidFrom    string `json:"valid_from"`
	ValidTo      string `json:"valid_to"`
}

func (s *Stock) Prepare() error {
	if len(s.ID) == 0 {
		uuid, err := uuid.NewV4()
		if err != nil {
			return err
		}
		s.ID = uuid.String()
	}

	return nil
}

func (s *Stock) Validate() error {
	if len(s.CompanyID) != 36 {
		return errors.New("Wrong Company ID")
	}
	if s.StocksNumber == 0 {
		return errors.New("Required Stocks Number")
	}
	if s.ValidFrom == "" {
		return errors.New("Required stocks number emitting date")
	}
	return nil
}

func (s *Stock) FindStocksByID(id string) (*Stock, error) {
	db, err := dbutils.DbConn()

	if err != nil {
		return &Stock{}, err
	}

	selDB, err := db.Query("SELECT * FROM stocks WHERE id = ?", id)
	if err != nil {
		return &Stock{}, err
	}

	//c = &Company{}
	for selDB.Next() {
		var id, companyID, validFrom, validTo string
		var stocksNumber int16
		err = selDB.Scan(&id, &companyID, &stocksNumber, &validFrom, &validTo)
		if err != nil {
			panic(err.Error())
		}
		s.ID = id
		s.CompanyID = companyID
		s.StocksNumber = stocksNumber
		s.ValidFrom = validFrom
		s.ValidTo = validTo
	}

	defer db.Close()
	return s, nil
}

func (s *Stock) InsertStock() error {
	db, err := dbutils.DbConn()
	if err != nil {
		return err
	}

	insForm, err := db.Prepare("INSERT INTO stocks(id, company_id, stocks_number,valid_from, valid_to) VALUES(?,?,?,?,?)")
	if err != nil {
		return err
	}

	insForm.Exec(s.ID, s.CompanyID, s.StocksNumber, s.ValidFrom, s.ValidFrom)

	//log.Println("INSERT: Name: " + name + " | City: " + city)

	defer db.Close()

	return nil
}
