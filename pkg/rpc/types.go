package rpc

import (
	"encoding/json"
	"time"
)

// Response is a generic response structure for RPC calls
type Response[T any] struct {
	JsonRpc string `json:"jsonrpc"`
	Id      int    `json:"id"`
	Result  T      `json:"result,omitempty"`
	Error   *Error `json:"error,omitempty"`
}

// Error represents an RPC error
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data,omitempty"`
}

func (e *Error) Error() string {
	return e.Message
}

// ResultBlock represents the result of a block query
type ResultBlock struct {
	BlockID BlockID `json:"block_id"`
	Block   Block   `json:"block"`
}

// BlockID represents a block identifier
type BlockID struct {
	Hash  string `json:"hash"`
	Parts struct {
		Total int    `json:"total"`
		Hash  string `json:"hash"`
	} `json:"parts"`
}

// Block represents a block in the blockchain
type Block struct {
	Header     Header     `json:"header"`
	Data       BlockData  `json:"data"`
	Evidence   Evidence   `json:"evidence"`
	LastCommit LastCommit `json:"last_commit"`
}

// Header represents the header of a block
type Header struct {
	Version struct {
		Block string `json:"block"`
	} `json:"version"`
	ChainID         string    `json:"chain_id"`
	Height          int64     `json:"height,string"`
	Time            time.Time `json:"time"`
	LastBlockID     BlockID   `json:"last_block_id"`
	LastCommitHash  string    `json:"last_commit_hash"`
	DataHash        string    `json:"data_hash"`
	ValidatorsHash  string    `json:"validators_hash"`
	NextValidatorsHash string `json:"next_validators_hash"`
	ConsensusHash   string    `json:"consensus_hash"`
	AppHash         string    `json:"app_hash"`
	LastResultsHash string    `json:"last_results_hash"`
	EvidenceHash    string    `json:"evidence_hash"`
	ProposerAddress string    `json:"proposer_address"`
}

// BlockData represents the data in a block
type BlockData struct {
	Txs []string `json:"txs"`
}

// Evidence represents the evidence in a block
type Evidence struct {
	Evidence []json.RawMessage `json:"evidence"`
}

// LastCommit represents the last commit in a block
type LastCommit struct {
	Height  int64  `json:"height,string"`
	Round   int    `json:"round"`
	BlockID BlockID `json:"block_id"`
	Signatures []Signature `json:"signatures"`
}

// Signature represents a signature in a commit
type Signature struct {
	BlockIDFlag      int       `json:"block_id_flag"`
	ValidatorAddress string    `json:"validator_address"`
	Timestamp        time.Time `json:"timestamp"`
	Signature        string    `json:"signature"`
}

// ResultStatus represents the result of a status query
type ResultStatus struct {
	NodeInfo      NodeInfo      `json:"node_info"`
	SyncInfo      SyncInfo      `json:"sync_info"`
	ValidatorInfo ValidatorInfo `json:"validator_info"`
}

// NodeInfo represents information about the node
type NodeInfo struct {
	ProtocolVersion struct {
		P2P   string `json:"p2p"`
		Block string `json:"block"`
		App   string `json:"app"`
	} `json:"protocol_version"`
	ID         string `json:"id"`
	ListenAddr string `json:"listen_addr"`
	Network    string `json:"network"`
	Version    string `json:"version"`
	Channels   string `json:"channels"`
	Moniker    string `json:"moniker"`
	Other      struct {
		TxIndex    string `json:"tx_index"`
		RPCAddress string `json:"rpc_address"`
	} `json:"other"`
}

// SyncInfo represents the sync status of the node
type SyncInfo struct {
	LatestBlockHash   string    `json:"latest_block_hash"`
	LatestAppHash     string    `json:"latest_app_hash"`
	LatestBlockHeight int64     `json:"latest_block_height,string"`
	LatestBlockTime   time.Time `json:"latest_block_time"`
	CatchingUp        bool      `json:"catching_up"`
}

// ValidatorInfo represents information about the validator
type ValidatorInfo struct {
	Address     string `json:"address"`
	PubKey      struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"pub_key"`
	VotingPower string `json:"voting_power"`
}
