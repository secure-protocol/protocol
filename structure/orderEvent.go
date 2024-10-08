package structure

type OrderEvent struct {
	OrderId string `json:"order_id"`
	Address string `json:"address"`
	TxHash  string `json:"tx_hash"`
}

type TransEvent struct {
	TxID          string `json:"txID"`
	NetworkKey    string `json:"networkKey"`
	PendingStatus string `json:"status"`
}
