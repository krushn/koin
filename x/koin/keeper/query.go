package keeper

import (
	"koin/x/koin/types"
)

var _ types.QueryServer = Keeper{}
