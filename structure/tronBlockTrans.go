package structure

import "gorm.io/gorm"

type TronBlockTrans struct {
	gorm.Model

	TxID          string `json:"txID,omitempty" gorm:"type:varchar(255);comment:交易id或者交易hash"`
	ChainName     string `json:"chainName,omitempty" gorm:"type:varchar(255);comment:链"`
	NetworkKey    string `json:"networkKey,omitempty" gorm:"type:varchar(255);comment:网络唯一标识符"`
	From          string `json:"from,omitempty" gorm:"type:varchar(255);comment:发送地址"`
	To            string `json:"to,omitempty" gorm:"type:varchar(255);comment:接收地址"`
	NativeValue   string `json:"nativeValue,omitempty" gorm:"type:varchar(255) ;comment: 主代币转账消耗金额，最小精度单位"`
	TokenAmount   string `json:"tokenAmount,omitempty" gorm:"type:varchar(255);comment: 代币转账金额，最小精度单位"`
	NativeDecimal uint64 `json:"nativeDecimal,omitempty" gorm:"comment: 主代币精度"`
	TokenDecimal  uint64 `json:"tokenDecimal,omitempty" gorm:"comment: 合约代币精度"`

	TransType       uint64 `json:"transType,omitempty" gorm:"comment: 波场交易类型 ，0普通trx交易 ，1 trc10 交易 ， 2 trc20交易"`
	ContractAddress string `json:"contractAddress,omitempty" gorm:"type: varchar(255);comment: 如果是合约交易 合约地址 不为空"`
	CoinID          uint64 `json:"coinID,omitempty" gorm:"type:varchar(255);comment:代币id"`
	CoinSymbol      string `json:"coinSymbol,omitempty" gorm:"type:varchar(255);comment:代币符号"`

	EnergyUsed    string `json:"energyUsed,omitempty" gorm:"type:varchar(255);comment: 能量花费"`
	BandWidthUsed string `json:"BandWidthUsed,omitempty" gorm:"type:varchar(255);comment: 带宽花费"`
	TrxUsed       string `json:"trxUsed,omitempty" gorm:"type:varchar(255);comment: 交易Trx花费TRX burned for resources"`

	Confirmation uint64 `json:"confirmation,omitempty" gorm:"comment: 交易确认数"`
	Block        uint64 `json:"block,omitempty" gorm:"comment: 区块高度"`
	Status       int    `json:"status,omitempty" gorm:"comment: 0 未上链 1 pending 2.成功 3 失败"`

	IsSystemSendTrans bool `json:"isSystemTrans,omitempty" gorm:"comment: 是否是我们自己发送的交易"`

	TransactionCreateTime uint64 `json:"transactionCreateTime,omitempty" gorm:"comment: 交易生成时间"`

	///一般是交易创建后60s
	TransactionExpireTime uint64 `json:"transactionExpireTime,omitempty" gorm:"comment: 交易过期时间"`

	TransactionBlockTime uint64 `json:"transactionBlockTime,omitempty" gorm:"comment: 交易所以区块时间"`
}

const (
	TronTransStatusNotOnChain = iota
	TronTransStatusPending
	TronTransStatusSuccess
	TronTransStatusFailed
)
