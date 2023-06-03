package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "oracle/testutil/keeper"
	"oracle/x/networking/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.NetworkingKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
