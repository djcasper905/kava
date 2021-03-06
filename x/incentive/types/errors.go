package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// DONTCOVER

// Incentive module errors
var (
	ErrClaimNotFound                 = sdkerrors.Register(ModuleName, 2, "no claim with input id found for owner and collateral type")
	ErrClaimPeriodNotFound           = sdkerrors.Register(ModuleName, 3, "no claim period found for id and collateral type")
	ErrInvalidAccountType            = sdkerrors.Register(ModuleName, 4, "account type not supported")
	ErrNoClaimsFound                 = sdkerrors.Register(ModuleName, 5, "no claims with collateral type found for address")
	ErrInsufficientModAccountBalance = sdkerrors.Register(ModuleName, 6, "module account has insufficient balance to pay claim")
	ErrAccountNotFound               = sdkerrors.Register(ModuleName, 7, "account not found")
	ErrInvalidMultiplier             = sdkerrors.Register(ModuleName, 8, "invalid rewards multiplier")
	ErrZeroClaim                     = sdkerrors.Register(ModuleName, 9, "cannot claim - claim amount rounds to zero")
)
