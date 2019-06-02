package types

import (
	"time"

	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
)

type Subscription struct {
	ID                 sdkTypes.ID          `json:"id"`
	TxHash             string               `json:"tx_hash"`
	ClientAddress      csdkTypes.AccAddress `json:"client_address"`
	ClientPubKey       crypto.PubKey        `json:"client_pub_key"`
	RemainingBandwidth sdkTypes.Bandwidth   `json:"remaining_bandwidth"`
	Status             string               `json:"status"`
	CreatedAt          time.Time            `json:"created_at"`
}
