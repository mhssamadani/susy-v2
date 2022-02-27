package common

import (
	"time"

	"github.com/SuSy-One/susy-v2/node/pkg/vaa"

	"github.com/ethereum/go-ethereum/common"
)

type MessagePublication struct {
	TxHash    common.Hash // TODO: rename to identifier? on Solana, this isn't actually the tx hash
	Timestamp time.Time

	Nonce            uint32
	Sequence         uint64
	ConsistencyLevel uint8
	EmitterChain     vaa.ChainID
	EmitterAddress   vaa.Address
	Payload          []byte
}

type ErgoMessagePublication struct {
	TxHash    [32]byte
	Timestamp time.Time

	Nonce            uint32
	Sequence         uint64
	ConsistencyLevel uint8
	EmitterChain     vaa.ChainID
	EmitterAddress   vaa.ErgoAddress
	Payload          []byte
}
