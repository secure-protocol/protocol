package platform

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Network struct {
	gorm.Model
	// ？
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

	ChainServerGrpcURL string
}

type Address struct {
	gorm.Model
	Bid     string `gorm:"type:varchar(255)"`
	Address string `json:"address,omitempty" gorm:"type:varchar(255);unique;comment:地址"`
	// Type enum: 1 收款 receive 2 资金归集地址（打款，备付） gather
	Type string `json:"type,omitempty" gorm:"type:varchar(255);comment:1 收款 receive 2 资金归集地址（打款，备付） gather 3. 资源地址 resource 4. 商家提现地址 withdraw"`
	// NetWorkKey enum: TRON  ETH  BSC  TRON-SHASTA
	NetworkKey string `json:"chain,omitempty" gorm:"type:varchar(255);comment:TRON  ETH  BSC  TRON-SHASTA"`
	// Status enum: 1 空闲（未分配给收款订单） free 2 占用 occupy  3 禁用 drop
	Status string `json:"status,omitempty" gorm:"type:varchar(255);comment:1 空闲 free 2 占用 occupy  3 禁用 drop"`
}

type ResourceAddress struct {
	gorm.Model
	Address string `json:"address,omitempty" gorm:"type:varchar(255);comment:地址"`
	// Type enum: 1 收款 receive 2 资金归集地址（打款，备付） gather 3. 资源地址 resource            4. 商家提现地址 5. 冷钱包地址
	//Type string `json:"type,omitempty" gorm:"type:varchar(255);comment:1 收款 receive 2 资金归集地址（打款，备付） gather 3. 资源地址 resource 4. 商家提现地址 withdraw"`
	// ChaNetWorkKey enum: TRON  ETH  BSC  TRON-SHASTA
	NetworkKey string `json:"chain,omitempty" gorm:"type:varchar(255);comment:TRON  ETH  BSC  TRON-SHASTA"`
	// Status enum: 1 空闲 free 2 占用 occupy  3 禁用 drop
	Status string `json:"status,omitempty" gorm:"type:varchar(255);comment:1 空闲 free 2 占用 occupy  3 禁用 drop"`
	// stake
	StakeEnergy    uint64
	StakeBandwidth uint64
	// 资源余量
	BandWidth uint64
	Energy    uint64
	USDT      uint64
	//
}

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
	// Status 启用状态
	Status bool `json:"-" gorm:"comment:启用状态"`
}

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
