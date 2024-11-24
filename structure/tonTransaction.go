package structure

import "gorm.io/gorm"

type TonBlockTrans struct {
	gorm.Model

	TxID       string `json:"txID,omitempty" gorm:"unique;not null;type:varchar(255);comment:交易id或者交易hash(base64格式)"`
	ChainName  string `json:"chainName,omitempty" gorm:"type:varchar(255);comment:链"`
	NetworkKey string `json:"networkKey,omitempty" gorm:"type:varchar(255);comment:网络唯一标识符"`
	//From       string `json:"from,omitempty" gorm:"type:varchar(255);comment:区块链上的发送地址"`
	//To         string `json:"to,omitempty" gorm:"type:varchar(255);comment:区块链交易中的to地址,可以是合约地址"`
	TraceID string `json:"traceID,omitempty" gorm:"type:varchar(255);comment:交易追踪ID"`

	JETTONMasterAddress string `json:"JETTONMasterAddress,omitempty" gorm:"type:varchar(255);comment:jetton代币地址"`
	Account             string `json:"account,omitempty" gorm:"type:varchar(255);comment:交易发送的账户地址,非用户友好地址"`
	FriendlyAccount     string `json:"friendlyAccount,omitempty" gorm:"type:varchar(255);comment:交易发送的账户地址,用户友好地址"`

	Send            string `json:"send,omitempty" gorm:"type:varchar(255);comment:资产发送地址 hex raw 地址"`
	FriendlySend    string `json:"friendlySend,omitempty" gorm:"type:varchar(255);comment:资产发送地址友好地址"`
	Receive         string `json:"receive,omitempty" gorm:"type:varchar(255);comment: 资产接收地址  hex raw 地址"`
	FriendlyReceive string `json:"friendlyReceive,omitempty" gorm:"type:varchar(255);comment:资产发送地址友好地址"`
	NativeValue     string `json:"nativeValue,omitempty" gorm:"type:varchar(255) ;comment: 主代币转账消耗金额，最小精度单位"`
	TokenAmount     string `json:"tokenAmount,omitempty" gorm:"type:varchar(255);comment: 代币转账金额，最小精度单位"`
	NativeDecimal   uint64 `json:"nativeDecimal,omitempty" gorm:"comment: 主代币精度"`
	TokenDecimal    uint64 `json:"tokenDecimal,omitempty" gorm:"comment: 合约代币精度"`

	//JettonWalletAddress    string `json:"contractAddress,omitempty" gorm:"type: varchar(255);comment: 如果是合约交易 合约地址 不为空"`
	CoinID     uint64 `json:"coinID,omitempty" gorm:"type:varchar(255);comment:代币id"`
	CoinSymbol string `json:"coinSymbol,omitempty" gorm:"type:varchar(255);comment:代币符号"`

	TransType uint64 `json:"transType,omitempty" gorm:"comment: 交易类型 1 tonCoin普通交易 ， 2 jetton交易"`
	InOpCode  string `json:"inOpCodee,omitempty" gorm:"comment: ton合约inMsg的opcode"`
	//OutOpCode   string `json:"outOpCode,omitempty" gorm:"comment: ton合约outMsg的opcode"`
	TotalGasFee string `json:"totalGasFee,omitempty" gorm:"type:varchar(255);comment: 总GasFee"`

	McSeqno uint64 `json:"mcSeqno,omitempty" gorm:"comment: 主链区块高度"`

	WorkChain  int64  `json:"workChain"`
	Shard      string `json:"shard"`
	ShardSeqno uint64 `json:"shardSeqno"  gorm:"comment: 分片区块高度"`

	Status        string `json:"status,omitempty" gorm:";enum('new',failed','success')"`
	PendingStatus string `json:"pendingStatus" gorm:"type:enum('new','pending','solid')"`

	IsSystemSendTrans bool `json:"isSystemTrans,omitempty" gorm:"comment: 是否是我们自己发送的交易"`
	IsSendOutTrans    bool `json:"isSendOutTrans,omitempty" gorm:"comment: 是否是发送到外部的交易"`

	TransactionCreateTime uint64 `json:"transactionCreateTime,omitempty" gorm:"comment: 交易生成时间"`

	TransactionBlockTime uint64 `json:"transactionBlockTime,omitempty" gorm:"comment: 交易所以区块时间"`

	DealStatus string `json:"dealStatus" gorm:"type:enum('new','pendingDealed','solidDealed')"`

	Memo string `json:"memo,omitempty" gorm:"comment: textComment"`

	BusinessType string `json:"BusinessType" gorm:"type:enum('','pay','withdraw','refund','collect','other');default:'';comment:业务类型"`
}
