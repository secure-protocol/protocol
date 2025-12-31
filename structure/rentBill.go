package structure

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

const (
	// ResourceTypeEnergy 资源类型 能量
	ResourceTypeEnergy = "energy"
	// ResourceTypeBandwidth 资源类型 带宽
	ResourceTypeBandwidth = "bandwidth"
)

const (
	// RentStatusSuccess  租用订单状态 成功
	RentStatusSuccess = "success"
	// RentStatusFailed  租用订单状态 失败
	RentStatusFailed = "failed"
)

const (
	// RentProviderGasStation 租用提供方 gasStation
	RentProviderGasStation = "gasStation"
)

type RentBill struct {
	gorm.Model
	BillID           string `gorm:"uniqueIndex:bill_rentType;type:varchar(191)"`
	ThirdPartyBillID string `gorm:"type:varchar(191);index"`
	RelatedBillID    string `gorm:"type:varchar(191);index;comment:关联的业务账单ID"`

	ResourceType   string `gorm:"uniqueIndex:bill_rentType;type:varchar(191)"`
	ResourceNum    int64
	RentMinutes    int64
	ReceiveAddress string
	TokenKind      string
	Price          decimal.Decimal `gorm:"type:decimal(40,20)"`

	Status string
	Remark string

	Provider string `gorm:"type:varchar(191);index;comment:提供方"`
}
