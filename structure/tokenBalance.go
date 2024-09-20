package structure

import "github.com/shopspring/decimal"

type TokenBalance struct {
	Address string `gorm:"type:varchar(255);primary_key"`
	TokenID string `gorm:"type:varchar(255);primary_key"`
	Balance decimal.Decimal
}
