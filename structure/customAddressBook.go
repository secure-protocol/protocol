package structure

import "gorm.io/gorm"

type CustomAddressBook struct {
	gorm.Model  `json:"-"`
	UserID      int64  `json:"userID" gorm:"comment:userID"`
	AddressID   string `json:"addressID" gorm:"comment:AddressID"`
	Network     string `json:"network,omitempty" gorm:"comment:所属网络"`
	Address     string `json:"address,omitempty" gorm:"comment:地址"`
	Name        string `json:"name,omitempty" gorm:"comment:名称"`
	TokenSymbol string `json:"tokenSymbol,omitempty" gorm:"comment:币种Symbol"`
	Token       uint   `json:"token,omitempty" gorm:"comment:币种"`
	Timestamp   int64  `json:"timestamp" gorm:"comment:创建时间"`
}
