package keeper

import (
	"secretai/x/secretai/types"
)

var _ types.QueryServer = Keeper{}
