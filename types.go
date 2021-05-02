package events

import (
	"github.com/centrifuge/go-substrate-rpc-client/v3/types"
)

type BlockRewardInfo struct {
	Seed			types.U256	`json:"seed"`
    OnlineTarget	types.U256	`json:"onlineTarget"`
    ComputeTarget	types.U256	`json:"computeTarget"`
}
