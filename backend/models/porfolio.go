package models

import (
	"errors"

	"github.com/gofrs/uuid"
	"github.com/silvano-bergamasco/business6sense/backend/utils/dbutils"
)

type Porfolio struct {
	ID               string  `json:"id"`
	UserID           string  `json:"user_id"`
	CompanyID        string  `json:"company_id"`
	StockMarketID    string  `json:"stock_market_id"`
	StocksNumber     float32 `json:"stocks_number"`
	PurchasePrice    float32 `json:"purchase_price"`
	PurchaseCurrency string  `json:"purchase_currency"`
	PurchaseDate     string  `json:"purchase_date"`
	SellPrice        float32 `json:"sell_price"`
	SellCurrency     string  `json:"sell_currency"`
	SellDate         string  `json:"sell_date"`
}

func (p *Porfolio) Prepare() error {
	if len(p.ID) == 0 {
		uuid, err := uuid.NewV4()
		if err != nil {
			return err
		}
		p.ID = uuid.String()
	}

	return nil
}

func (p *Porfolio) Validate(action string) error {
	if len(p.UserID) != 36 {
		return errors.New("Wrong User ID")
	}
	if len(p.CompanyID) != 36 {
		return errors.New("Wrong Company ID")
	}
	if p.StocksNumber == 0 {
		return errors.New("Required Stocks Number")
	}
	if p.PurchasePrice == 0 {
		return errors.New("Required purchase price")
	}
	if p.PurchaseCurrency == "" {
		return errors.New("Required purchase currency")
	}
	if p.PurchaseDate == "" {
		return errors.New("Required purchase date")
	}

	if action == "update" {
		if p.SellPrice == 0 {
			return errors.New("Required sell price")
		}
		if p.SellCurrency == "" {
			return errors.New("Required sell currency")
		}
		if p.SellDate == "" {
			return errors.New("Required sell date")
		}
	}

	return nil
}

func (p *Porfolio) FindPortfolioByID(id string) (*Porfolio, error) {
	db, err := dbutils.DbConn()

	if err != nil {
		return &Porfolio{}, err
	}

	selDB, err := db.Query("SELECT * FROM portfolios WHERE id = ?", id)
	if err != nil {
		return &Porfolio{}, err
	}

	for selDB.Next() {
		var id, userId, companyId, stockMarketId, purchaseCurrency, purchaseDate, sellCurrency, sellDate string
		var stocksNumber, purchasePrice, sellPrice float32
		err = selDB.Scan(&id, &userId, &companyId, &stockMarketId, &stocksNumber, &purchasePrice, &purchaseCurrency, &purchaseDate, &sellPrice, &sellCurrency, &sellDate)
		if err != nil {
			panic(err.Error())
		}
		p.ID = id
		p.UserID = userId
		p.CompanyID = companyId
		p.StockMarketID = stockMarketId
		p.StocksNumber = stocksNumber
		p.PurchasePrice = purchasePrice
		p.PurchaseCurrency = purchaseCurrency
		p.PurchaseDate = purchaseDate
		p.SellPrice = sellPrice
		p.SellCurrency = sellCurrency
		p.SellDate = sellDate
	}

	defer db.Close()
	return p, nil
}

func (p *Porfolio) InsertPortfolio() error {
	db, err := dbutils.DbConn()
	if err != nil {
		return err
	}

	insForm, err := db.Prepare("INSERT INTO portfolios(id, user_id, company_id, stock_market_id, stocks_number, purchase_price, purchase_currency, purchase_date) VALUES(?,?,?,?,?,?,?,?)")
	if err != nil {
		return err
	}

	insForm.Exec(p.ID, p.UserID, p.CompanyID, p.StockMarketID, p.StocksNumber, p.PurchasePrice, p.PurchaseCurrency, p.PurchaseDate)

	//log.Println("INSERT: Name: " + name + " | City: " + city)

	defer db.Close()

	return nil
}
