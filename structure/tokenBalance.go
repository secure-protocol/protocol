package structure

import "github.com/shopspring/decimal"

type TokenBalance struct {
	TokenID     uint            `json:"tokenID" gorm:"primarykey;comment:tokenID"`
	Address     string          `json:"address" gorm:"primarykey;type:varchar(255)"`
	AddressType string          `json:"addressType" gorm:"type:enum('receive','gather')'"`
	Balance     decimal.Decimal `json:"balance" gorm:"type:decimal(40,20);comment:余额"`
	Occupy      decimal.Decimal `json:"occupy" gorm:"type:decimal(40,20);comment:占用"`
}
