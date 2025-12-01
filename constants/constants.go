package constants

const (
	RedisKeyListeningAddressMap      = "ListeningAddressMap_"
	RedisKeyBlockInfo                = "BlockInfo_"
	RedisKeyBlockInfoBlockNum        = "latestBlockNumber"
	RedisKeyBlockInfoSolidBlockNum   = "latestSolidBlockNumber"
	RedisKeyBlockInfoPendingBlockNum = "latestPendingBlockNumber"
)

const (
	OrderEventExchange = "OrderEvent"
	OrderEventQueue    = "OrderEventQueue"
)

const (
	TransEventExchange = "TransEvent"
	TransEventQueue    = "TransEventQueue"
)
