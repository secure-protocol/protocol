package structure

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Network struct {
	gorm.Model
	NetworkKey string `json:"networkKey,omitempty" gorm:"type:varchar(255);comment:唯一标识符"`
	AccountKey string `json:"accountKey,omitempty" gorm:"type:varchar(255);comment:关联的私钥地址生成类型key"`
	ChainID    string `json:"chainID,omitempty" gorm:"type:varchar(255);comment:某些网络对应的chainID"`
	// 网络原生币信息
	Decimal        uint64 `json:"decimal,omitempty" gorm:"comment:网络代币精度"`
	DisplayDecimal uint64 `json:"displayDecimal,omitempty" gorm:"comment:网络代币展示精度"`
	Logo           string `json:"logo,omitempty" gorm:"type:varchar(255);comment:网络代币logo"`
	Symbol         string `json:"symbol,omitempty" gorm:"type:varchar(255);comment:网络代币符号"`
	FullSymbol     string `json:"fullSymbol,omitempty" gorm:"type:varchar(255);comment:网络代币完整符号"`

	ArchType string `json:"archType,omitempty" gorm:"type:varchar(255);comment:网络对应链的架构类型"`

	RpcURL              string `json:"rpcURL,omitempty" gorm:"type:varchar(255);comment:网络rpcUrl，逗号分隔存储"`
	ConfirmStrategy     uint64 `json:"confirmStrategy,omitempty" gorm:"comment:网络确认数字"`
	BlockGenerationRate uint64 `json:"blockGenerationRate,omitempty" gorm:"comment:区块生成周期 s/block"`
	ScanURL             string `json:"ScanURL,omitempty" gorm:"comment:区块链浏览器"`

	SystemAvailableTypes uint64 `json:"UseTypes,omitempty" gorm:"comment:系统使用模块"`

	HangTime uint64 `json:"orderExpireSecond" gorm:"comment:挂起时间 单位minutes;default:10"`

	ChainServerGrpcURL string
	Status             bool `json:"status" gorm:"comment:network 启用状态;default:true"`

	NetFee decimal.Decimal `json:"netFee,omitempty" gorm:"type:decimal(40,20);comment:手续费"`
}

const (
	CreateMask = 1 << iota
	StakingMask
	ReceiveMask
	PayMask
)
