package builder

import (
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
)

type BuiltBlock struct {
	BlockId              uint64    `db:"block_id"`
	BlockNumber          uint64    `db:"block_number"`
	Profit               string    `db:"profit"`
	Slot                 uint64    `db:"slot"`
	Hash                 string    `db:"hash"`
	GasLimit             uint64    `db:"gas_limit"`
	GasUsed              uint64    `db:"gas_used"`
	BaseFee              uint64    `db:"base_fee"`
	ParentHash           string    `db:"parent_hash"`
	ProposerPubkey       string    `db:"proposer_pubkey"`
	ProposerFeeRecipient string    `db:"proposer_fee_recipient"`
	BuilderPubkey        string    `db:"builder_pubkey"`
	Timestamp            uint64    `db:"timestamp"`
	TimestampDatetime    time.Time `db:"timestamp_datetime"`
}

type BuiltBlockBundle struct {
	BlockId     uint64  `db:"block_id"`
	BundleId    *uint64 `db:"bundle_id"`
	BlockNumber uint64  `db:"block_number"`
	BundleHash  string  `db:"bundle_hash"`
}

type DbBundle struct {
	DbId       uint64 `db:"id"`
	BundleHash string `db:"bundle_hash"`

	ParamSignedTxs         string    `db:"param_signed_txs"`
	ParamBlockNumber       uint64    `db:"param_block_number"`
	ParamTimestamp         uint64    `db:"param_timestamp"`
	ReceivedTimestamp      time.Time `db:"received_timestamp"`
	ParamRevertingTxHashes string    `db:"param_reverting_tx_hashes"`

	CoinbaseDiff      string `db:"coinbase_diff"`
	TotalGasUsed      uint64 `db:"total_gas_used"`
	StateBlockNumber  uint64 `db:"state_block_number"`
	GasFees           string `db:"gas_fees"`
	EthSentToCoinbase string `db:"eth_sent_to_coinbase"`
}

func SimulatedBundleToDbBundle(bundle *types.SimulatedBundle) DbBundle {
	revertingTxHashes := make([]string, len(bundle.OriginalBundle.RevertingTxHashes))
	for i, rTxHash := range bundle.OriginalBundle.RevertingTxHashes {
		revertingTxHashes[i] = rTxHash.String()
	}

	signedTxsStrings := make([]string, len(bundle.OriginalBundle.Txs))
	for i, tx := range bundle.OriginalBundle.Txs {
		signedTxsStrings[i] = tx.Hash().String()
	}

	return DbBundle{
		BundleHash: bundle.OriginalBundle.Hash.String(),

		ParamSignedTxs:         strings.Join(signedTxsStrings, ","),
		ParamBlockNumber:       bundle.OriginalBundle.BlockNumber.Uint64(),
		ParamTimestamp:         bundle.OriginalBundle.MinTimestamp,
		ParamRevertingTxHashes: strings.Join(revertingTxHashes, ","),

		CoinbaseDiff:      new(big.Rat).SetFrac(bundle.TotalEth, big.NewInt(1e18)).FloatString(18),
		TotalGasUsed:      bundle.TotalGasUsed,
		StateBlockNumber:  bundle.OriginalBundle.BlockNumber.Uint64(),
		GasFees:           new(big.Int).Mul(big.NewInt(int64(bundle.TotalGasUsed)), bundle.MevGasPrice).String(),
		EthSentToCoinbase: new(big.Rat).SetFrac(bundle.EthSentToCoinbase, big.NewInt(1e18)).FloatString(18),
	}
}
