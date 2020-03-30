package models

import (
	"errors"

	"github.com/gofrs/uuid"
	"github.com/silvano-bergamasco/business6sense/backend/utils/dbutils"
)

type FinancialStatement struct {
	ID               string  `json:"id"`
	Assets           float32 `json:"assets"`
	NetFinancialDebt float32 `json:"net_financial_debt"`
	NetAssets        float32 `json:"net_assets"`
	ActiveCurrent    float32 `json:"active_current"`
	PassiveCurrent   float32 `json:"passive_current"`
	NetIncome        float32 `json:"net_income"`
	Edibt            float32 `json:"edibt"`
}

func (fs *FinancialStatement) Prepare() error {
	if len(fs.ID) == 0 {
		uuid, err := uuid.NewV4()
		if err != nil {
			return err
		}
		fs.ID = uuid.String()
	}
	return nil
}

func (fs *FinancialStatement) Validate() error {
	if fs.Assets == 0 {
		return errors.New("Required Assets")
	}
	if fs.NetFinancialDebt == 0 {
		return errors.New("Required Net Financial Debt")
	}
	if fs.NetAssets == 0 {
		return errors.New("Required Net Assets")
	}
	if fs.ActiveCurrent == 0 {
		return errors.New("Required Active Current")
	}
	if fs.PassiveCurrent == 0 {
		return errors.New("Required Passive Current")
	}
	if fs.NetIncome == 0 {
		return errors.New("Required Net Income")
	}
	if fs.Edibt == 0 {
		return errors.New("Required Edibt")
	}
	return nil
}

func (fs *FinancialStatement) FindFinancialStatementByID(id string) (*FinancialStatement, error) {
	db, err := dbutils.DbConn()

	if err != nil {
		return &FinancialStatement{}, err
	}

	selDB, err := db.Query("SELECT * FROM finantial_statements WHERE id = ?", id)
	if err != nil {
		return &FinancialStatement{}, err
	}

	//c = &Company{}
	for selDB.Next() {
		var id string
		var assets, netFinancialDebt, netAssets, activeCurrent, passiveCurrent, netIncome, edibt float32
		err = selDB.Scan(&id, &assets, &netFinancialDebt, &netAssets, &activeCurrent, &passiveCurrent, &netIncome, &edibt)
		if err != nil {
			panic(err.Error())
		}
		fs.ID = id
		fs.Assets = assets
		fs.NetFinancialDebt = netFinancialDebt
		fs.NetAssets = netAssets
		fs.ActiveCurrent = activeCurrent
		fs.PassiveCurrent = passiveCurrent
		fs.NetIncome = netIncome
		fs.Edibt = edibt
	}

	defer db.Close()
	return fs, nil
}

func (fs *FinancialStatement) InsertFinancialStatement() error {
	db, err := dbutils.DbConn()
	if err != nil {
		return err
	}

	insForm, err := db.Prepare("INSERT INTO financial_statements(id, assets, net_financial_bebt, net_assets, active_current, passive_current, net_income, edibt) VALUES(?,?,?,?,?,?,?,?)")
	if err != nil {
		return err
	}

	insForm.Exec(fs.ID, fs.Assets, fs.NetFinancialDebt, fs.NetAssets, fs.ActiveCurrent, fs.PassiveCurrent, fs.NetIncome, fs.Edibt)

	//log.Println("INSERT: Name: " + name + " | City: " + city)

	defer db.Close()

	return nil
}
