// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package transaction

import (
	"encoding/json"
	"errors"
	"math/big"

	"github.com/kowala-tech/equilibrium/accounts"
	"github.com/kowala-tech/equilibrium/common"
	"github.com/kowala-tech/equilibrium/common/hexutil"
	"github.com/kowala-tech/equilibrium/crypto"
)

var _ = (*receiptMarshaling)(nil)

// MarshalJSON marshals as JSON.
func (r Receipt) MarshalJSON() ([]byte, error) {
	type Receipt struct {
		PostState       hexutil.Bytes    `json:"postState"`
		Status          hexutil.Uint64   `json:"status"`
		Bloom           common.Bloom     `json:"logsBloom"                     gencodec:"required"`
		Logs            []*Log           `json:"logs"                          gencodec:"required"`
		ComputeFee      *hexutil.Big     `json:"computeFee"                    gencodec:"required"`
		StabilityFee    *hexutil.Big     `json:"stabilityFee"                  gencodec:"required"`
		TxHash          crypto.Hash      `json:"transactionHash"  gencodec:"required"`
		ContractAddress accounts.Address `json:"contractAddress"`
	}
	var enc Receipt
	enc.PostState = r.PostState
	enc.Status = hexutil.Uint64(r.Status)
	enc.Bloom = r.Bloom
	enc.Logs = r.Logs
	enc.ComputeFee = (*hexutil.Big)(r.ComputeFee)
	enc.StabilityFee = (*hexutil.Big)(r.StabilityFee)
	enc.TxHash = r.TxHash
	enc.ContractAddress = r.ContractAddress
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (r *Receipt) UnmarshalJSON(input []byte) error {
	type Receipt struct {
		PostState       *hexutil.Bytes    `json:"postState"`
		Status          *hexutil.Uint64   `json:"status"`
		Bloom           *common.Bloom     `json:"logsBloom"                     gencodec:"required"`
		Logs            []*Log            `json:"logs"                          gencodec:"required"`
		ComputeFee      *hexutil.Big      `json:"computeFee"                    gencodec:"required"`
		StabilityFee    *hexutil.Big      `json:"stabilityFee"                  gencodec:"required"`
		TxHash          *crypto.Hash      `json:"transactionHash"  gencodec:"required"`
		ContractAddress *accounts.Address `json:"contractAddress"`
	}
	var dec Receipt
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.PostState != nil {
		r.PostState = *dec.PostState
	}
	if dec.Status != nil {
		r.Status = uint64(*dec.Status)
	}
	if dec.Bloom == nil {
		return errors.New("missing required field 'logsBloom' for Receipt")
	}
	r.Bloom = *dec.Bloom
	if dec.Logs == nil {
		return errors.New("missing required field 'logs' for Receipt")
	}
	r.Logs = dec.Logs
	if dec.ComputeFee == nil {
		return errors.New("missing required field 'computeFee' for Receipt")
	}
	r.ComputeFee = (*big.Int)(dec.ComputeFee)
	if dec.StabilityFee == nil {
		return errors.New("missing required field 'stabilityFee' for Receipt")
	}
	r.StabilityFee = (*big.Int)(dec.StabilityFee)
	if dec.TxHash == nil {
		return errors.New("missing required field 'transactionHash' for Receipt")
	}
	r.TxHash = *dec.TxHash
	if dec.ContractAddress != nil {
		r.ContractAddress = *dec.ContractAddress
	}
	return nil
}
