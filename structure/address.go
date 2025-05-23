package structure

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	Bid     string `gorm:"type:varchar(255)"`
	Address string `json:"address,omitempty" gorm:"unique;type:varchar(255);unique;comment:地址"`
	// Type enum: 1 收款 receive 2 资金归集地址（打款，备付） gather
	Type string `json:"type,omitempty" gorm:"type:enum('receive','gather')"`
	// NetWorkKey enum: TRON  ETH  BSC  TRON-SHASTA
	NetworkKey string `json:"chain,omitempty" gorm:"type:varchar(255);comment:TRON  ETH  BSC  TRON-SHASTA"`
	// SOLANA网络地址与币种绑定
	TokenKind string `gorm:"type:varchar(255)"`
	// Status enum: 1 空闲（未分配给收款订单） free 2 占用 occupy  3 禁用 drop
	Status string `json:"status,omitempty" gorm:"type:enum('free','occupy','drop','taken');default:free;"`
	// 用户ID 查询时经常作为条件，建立索引
	UserID int64 `json:"userID,omitempty" gorm:"default:0;index"`

	PublicKey string `json:"publicKey,omitempty" gorm:"type:varchar(255);comment:16进制公钥"`
}

const (
	AddressStatusFree   = "free"
	AddressStatusOccupy = "occupy"
	AddressStatusDrop   = "drop"
	AddressStatusTaken  = "taken"
)

const (
	AddressTypeReceive = "receive"
	AddressTypeGather  = "gather"
)
