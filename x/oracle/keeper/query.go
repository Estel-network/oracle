package keeper

import (
	"oracle/x/oracle/types"
)

var _ types.QueryServer = Keeper{}
