package structure

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Transfer struct {
	gorm.Model
	TxHash      string `gorm:"type:varchar(255);comment:solana为sig+index"`
	BlockHeight uint64 `gorm:"comment:solana为slot"`
	Timestamp   int64
	NetworkKey  string `gorm:"type:enum('TRON','TRON-NILE','TRON-SHASTA','SOL','SOL-DEVNET','BTC','ETH')"`
	TokenSymbol string `json:"type:enum('USDT','BTC','ETH','SOL')"`

	Sender   string `gorm:"comment:资金发送者"`
	Receiver string `gorm:"comment:资金接收者"`

	IsSystemSendTrans bool `gorm:"comment: 是否是我们自己发送的交易"`
	IsSendOutTrans    bool `gorm:"comment: 是否是发送到外部的交易"`

	PendingStatus string `gorm:"type:enum('new','pending','solid');comment:块确认状态"`
	DealStatus    string `gorm:"type:enum('new','pendingDealed','solidDealed');comment:交易处理状态"`
	Status        string `gorm:"type:enum('new','failed','success');comment:交易状态"`

	Amount decimal.Decimal `gorm:"comment:金额"`
	Unit   string          `gorm:"comment:单位"`

	SenderBalance   decimal.Decimal `gorm:"comment:资金发送者发送成功后的资金余额"`
	ReceiverBalance decimal.Decimal `gorm:"comment:资金接收者 交易成功后的资金余额"`

	BillID string `gorm:"comment:账单ID"`
}

const (
	TransferDealStatusNew     = "new"
	TransferDealStatusPending = "pendingDealed"
	TransferDealStatusSolid   = "solidDealed"
)

const (
	TransferStatusNew     = "new"
	TransferStatusSuccess = "success"
	TransferStatusFailed  = "failed"
)

const (
	TransferPendingStatus        = "TronTransPendingStatus"
	TransferPendingStatusNew     = "new"
	TransferPendingStatusPending = "pending"
	TransferPendingStatusSolid   = "solid"
)

type SolanaBlockTrans struct {
	gorm.Model

	Signature  string `json:"txID,omitempty" gorm:"unique;not null;type:varchar(255);comment:交易id或者交易hash"`
	Index      int64
	ChainName  string `json:"chainName,omitempty" gorm:"type:varchar(255);comment:链"`
	NetworkKey string `json:"networkKey,omitempty" gorm:"type:varchar(255);comment:网络唯一标识符"`
	Payer      string `json:"from,omitempty" gorm:"type:varchar(255);comment:区块链上的发送地址"`
	ProgramID  string `json:"to,omitempty" gorm:"type:varchar(255);comment:区块链交易中的to地址,可以是合约地址"`
	Data       string `json:"data"`

	Send    string `json:"send,omitempty" gorm:"type:varchar(255);comment:资产发送地址"`
	Receive string `json:"receive,omitempty" gorm:"type:varchar(255);comment: 资产接收地址"`

	TokenID     uint64 `json:"coinID,omitempty" gorm:"type:varchar(255);comment:代币id"`
	TokenSymbol string `json:"coinSymbol,omitempty" gorm:"type:varchar(255);comment:代币符号"`

	Block uint64 `json:"block,omitempty" gorm:"comment: 区块高度"`

	Status string `json:"status,omitempty" gorm:";enum('new',failed','success')"`

	PendingStatus string `json:"pendingStatus" gorm:"type:enum('new','pending','solid')"'`

	IsSystemSendTrans bool `json:"isSystemTrans,omitempty" gorm:"comment: 是否是我们自己发送的交易"`
	IsSendOutTrans    bool `json:"isSendOutTrans,omitempty" gorm:"comment: 是否是发送到外部的交易"`

	///hash为 4c3c637901c3cd26c751411505ffc028d9f0666b541053b135138e1bf097f17f的 "timestamp": 638637268693249270 ，所以此字段不可信
	TransactionCreateTime uint64 `json:"transactionCreateTime,omitempty" gorm:"comment: 交易生成时间"`

	///一般是交易创建后60s
	TransactionExpireTime uint64 `json:"transactionExpireTime,omitempty" gorm:"comment: 交易过期时间"`

	TransactionBlockTime uint64 `json:"transactionBlockTime,omitempty" gorm:"comment: 交易所以区块时间"`

	DealStatus string `json:"dealStatus" gorm:"type:enum('new','pendingDealed','solidDealed')"`

	BusinessType string `json:"BusinessType" gorm:"type:enum('','pay','withdraw','refund','collect','other');default:'';comment:业务类型"`
}
