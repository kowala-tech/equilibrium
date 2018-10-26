// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package transaction

import (
	"encoding/json"
	"errors"

	"github.com/kowala-tech/equilibrium/accounts"
	"github.com/kowala-tech/equilibrium/common/hexutil"
	"github.com/kowala-tech/equilibrium/crypto"
)

var _ = (*logMarshaling)(nil)

// MarshalJSON marshals as JSON.
func (l Log) MarshalJSON() ([]byte, error) {
	type Log struct {
		ContractAddress accounts.Address `json:"address" gencodec:"required"`
		Topics          []crypto.Hash    `json:"topics" gencodec:"required"`
		Data            hexutil.Bytes    `json:"data" gencodec:"required"`
		BlockNumber     hexutil.Uint64   `json:"blockNumber"`
		TxHash          crypto.Hash      `json:"transactionHash" gencodec:"required"`
		TxIndex         hexutil.Uint     `json:"transactionIndex" gencodec:"required"`
		BlockHash       crypto.Hash      `json:"blockHash"`
		Index           hexutil.Uint     `json:"logIndex" gencodec:"required"`
		Removed         bool             `json:"removed"`
	}
	var enc Log
	enc.ContractAddress = l.ContractAddress
	enc.Topics = l.Topics
	enc.Data = l.Data
	enc.BlockNumber = hexutil.Uint64(l.BlockNumber)
	enc.TxHash = l.TxHash
	enc.TxIndex = hexutil.Uint(l.TxIndex)
	enc.BlockHash = l.BlockHash
	enc.Index = hexutil.Uint(l.Index)
	enc.Removed = l.Removed
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (l *Log) UnmarshalJSON(input []byte) error {
	type Log struct {
		ContractAddress *accounts.Address `json:"address" gencodec:"required"`
		Topics          []crypto.Hash     `json:"topics" gencodec:"required"`
		Data            *hexutil.Bytes    `json:"data" gencodec:"required"`
		BlockNumber     *hexutil.Uint64   `json:"blockNumber"`
		TxHash          *crypto.Hash      `json:"transactionHash" gencodec:"required"`
		TxIndex         *hexutil.Uint     `json:"transactionIndex" gencodec:"required"`
		BlockHash       *crypto.Hash      `json:"blockHash"`
		Index           *hexutil.Uint     `json:"logIndex" gencodec:"required"`
		Removed         *bool             `json:"removed"`
	}
	var dec Log
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.ContractAddress == nil {
		return errors.New("missing required field 'address' for Log")
	}
	l.ContractAddress = *dec.ContractAddress
	if dec.Topics == nil {
		return errors.New("missing required field 'topics' for Log")
	}
	l.Topics = dec.Topics
	if dec.Data == nil {
		return errors.New("missing required field 'data' for Log")
	}
	l.Data = *dec.Data
	if dec.BlockNumber != nil {
		l.BlockNumber = uint64(*dec.BlockNumber)
	}
	if dec.TxHash == nil {
		return errors.New("missing required field 'transactionHash' for Log")
	}
	l.TxHash = *dec.TxHash
	if dec.TxIndex == nil {
		return errors.New("missing required field 'transactionIndex' for Log")
	}
	l.TxIndex = uint(*dec.TxIndex)
	if dec.BlockHash != nil {
		l.BlockHash = *dec.BlockHash
	}
	if dec.Index == nil {
		return errors.New("missing required field 'logIndex' for Log")
	}
	l.Index = uint(*dec.Index)
	if dec.Removed != nil {
		l.Removed = *dec.Removed
	}
	return nil
}
