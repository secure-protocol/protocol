package structure

import "gorm.io/gorm"

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
