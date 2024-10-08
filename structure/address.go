package structure

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	Bid     string `gorm:"type:varchar(255)"`
	Address string `json:"address,omitempty" gorm:"type:varchar(255);unique;comment:地址"`
	// Type enum: 1 收款 receive 2 资金归集地址（打款，备付） gather
	Type string `json:"type,omitempty" gorm:"type:enum('receive','gather')"`
	// NetWorkKey enum: TRON  ETH  BSC  TRON-SHASTA
	NetworkKey string `json:"chain,omitempty" gorm:"type:varchar(255);comment:TRON  ETH  BSC  TRON-SHASTA"`
	// Status enum: 1 空闲（未分配给收款订单） free 2 占用 occupy  3 禁用 drop
	Status string `json:"status,omitempty" gorm:"type:enum('free','occupy','drop');default:free;"`
	// 用户ID 查询时经常作为条件，建立索引
	UserID int64 `json:"userID,omitempty" gorm:"default:0;index"`
}

const (
	AddressStatusFree   = "free"
	AddressStatusOccupy = "occupy"
	AddressStatusDrop   = "drop"
)

const (
	AddressTypeReceive = "receive"
	AddressTypeGather  = "gather"
)
