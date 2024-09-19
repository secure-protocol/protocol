package structure

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

// Token
// 使用id作为业务主键，id自定义(不自增)，不删除
type Token struct {
	gorm.Model
	//TokenID string `json:"tokenID" gorm:"type:varchar(255);comment:tokenID  {NetworkKey}-{Protocol}-{Symbol}"`
	// Network enum: tron
	NetworkKey string `json:"network,omitempty" gorm:"type:varchar(255);comment:链网络 enum:tron"`
	// Protocol enum:trc20
	Protocol string `json:"protocol,omitempty" gorm:"type:varchar(255);comment:链协议 enum:trc20"`
	// Symbol enum: USDT
	Symbol string `json:"token,omitempty" gorm:"type:varchar(20);comment:币种符号 enum: USDT"`

	FullName string `json:"fullName,omitempty" gorm:"type:varchar(50);comment:币种全名USDT:Tether USD BTC:Bitcoin"`

	ContractAddress string `json:"contract,omitempty" gorm:"type:varchar(255);comment:代币链上合约地址"`

	Logo string `json:"logo" gorm:"type:varchar(255);comment:logo图标"`

	// 单位
	DefaultUnit string `json:"defaultUnit" gorm:"type:varchar(255);comment:单位"`
	// 最小单位
	MinUnit string `json:"minUnit" gorm:"type:varchar(255);comment:最小单位"`

	// 展示精度
	DisplayDecimal uint `json:"displayDecimal" gorm:"type:varchar(255);comment:展示精度"`

	// 转账费率，部分token在转账时在合约中需要付出手续费
	FeeRate decimal.Decimal `json:"feeRate,omitempty" gorm:"type:varchar(255);comment:转账费率，部分token在转账时在合约中需要付出手续费"`
	// Status 启用状态 on，停用 off
	Status bool `json:"-" gorm:"comment:启用状态"`
}
