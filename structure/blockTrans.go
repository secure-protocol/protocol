package structure

import "gorm.io/gorm"

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
