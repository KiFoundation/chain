package types

import (
	"fmt"
	"strings"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

// Default parameter constants
const (
	DefaultDenom            = "ibc/05554A9BFDD28894D7F18F4C707AA0930D778751A437A9FE1F4684A3E1199728" // USDC ibc denom from Osmosis to Lum Network mainnet
	DefaultMinDepositAmount = 1000000
)

// Default denoms variable
var DefaultDenoms = []string{DefaultDenom}

// Parameter store keys.
var (
	KeyDepositDenom     = []byte("DepositDenom")
	KeyMinDepositAmount = []byte("MinDepositAmount")
)

// ParamKeyTable for dfract module.
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// DefaultParams return the default dfract module params
func DefaultParams() Params {
	return Params{
		DepositDenoms:    DefaultDenoms,
		MinDepositAmount: DefaultMinDepositAmount,
	}
}

func (p *Params) Validate() error {
	if err := validateDepositDenom(p.DepositDenoms); err != nil {
		return err
	}

	if err := validateMinDepositAmount(p.MinDepositAmount); err != nil {
		return err
	}
	return nil
}

func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyDepositDenom, &p.DepositDenoms, validateDepositDenom),
		paramtypes.NewParamSetPair(KeyMinDepositAmount, &p.MinDepositAmount, validateMinDepositAmount),
	}
}

// Function that ensures that the deposited denom is an array of string
func validateDepositDenom(i interface{}) error {
	v, ok := i.([]string)

	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	for _, denom := range v {
		if strings.TrimSpace(denom) == "" {
			return ErrInvalidDepositDenom
		}
	}

	return nil
}

// Function that ensures that the deposited amount is not inferior or equal to 0
func validateMinDepositAmount(i interface{}) error {
	v, ok := i.(uint32)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	if v <= 0 {
		return ErrInvalidMinDepositAmount
	}
	return nil
}
