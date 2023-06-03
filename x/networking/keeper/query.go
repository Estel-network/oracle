package keeper

import (
	"oracle/x/networking/types"
)

var _ types.QueryServer = Keeper{}
